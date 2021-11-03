// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.Datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package tagset

import (
	"strings"

	"github.com/twmb/murmur3"
)

// A cachingFactory caches tagsets with no eviction policy.
//
// This type implements Factory.
type cachingFactory struct {
	baseFactory
	caches [numCacheIDs]map[uint64]*Tags
}

func newCachingFactory() *cachingFactory {
	var caches [numCacheIDs]map[uint64]*Tags
	for i := range caches {
		caches[i] = make(map[uint64]*Tags)
	}
	return &cachingFactory{
		caches: caches,
	}
}

// NewTags implements Factory.NewTags
func (f *cachingFactory) NewTags(tags []string) *Tags {
	tagsMap := make(map[uint64]string, len(tags))
	hash := uint64(0)
	for _, t := range tags {
		h := murmur3.StringSum64(t)
		_, seen := tagsMap[h]
		if seen {
			continue
		}
		tagsMap[h] = t
		hash ^= h
	}

	return f.getCachedTags(byTagsetHashCache, hash, func() *Tags {
		// write hashes and rewrite tags based on the map
		hashes := make([]uint64, len(tagsMap))
		tags = tags[:len(tagsMap)]
		i := 0
		for h, t := range tagsMap {
			tags[i] = t
			hashes[i] = h
			i++
		}

		return &Tags{tags, hashes, hash}
	})
}

// NewUniqueTags implements Factory.NewUniqueTags
func (f *cachingFactory) NewUniqueTags(tags ...string) *Tags {
	hashes, hash := calcHashes(tags)
	return f.getCachedTags(byTagsetHashCache, hash, func() *Tags {
		return &Tags{tags, hashes, hash}
	})
}

// NewTagsFromMap implements Factory.NewTagsFromMap
func (f *cachingFactory) NewTagsFromMap(src map[string]struct{}) *Tags {
	tags := make([]string, 0, len(src))
	for tag := range src {
		tags = append(tags, tag)
	}
	hashes, hash := calcHashes(tags)
	return f.getCachedTags(byTagsetHashCache, hash, func() *Tags {
		return &Tags{tags, hashes, hash}
	})
}

// NewTag implements Factory.NewTag
func (f *cachingFactory) NewTag(tag string) *Tags {
	hash := murmur3.StringSum64(tag)
	return f.getCachedTags(byTagsetHashCache, hash, func() *Tags {
		return &Tags{[]string{tag}, []uint64{hash}, hash}
	})
}

// NewBuilder implements Factory.NewBuilder
func (f *cachingFactory) NewBuilder(capacity int) *Builder {
	return f.baseFactory.newBuilder(f, capacity)
}

// NewSliceBuilder implements Factory.NewSliceBuilder
func (f *cachingFactory) NewSliceBuilder(levels, capacity int) *SliceBuilder {
	return f.baseFactory.newSliceBuilder(f, levels, capacity)
}

// ParseDSD implements Factory.ParseDSD
func (f *cachingFactory) ParseDSD(data []byte) (*Tags, error) {
	// TODO: GO FASTER
	return f.getCachedTags(byDSDHashCache, murmur3.Sum64(data), func() *Tags {
		tags := strings.Split(string(data), ",")
		return f.NewTags(tags)
	}), nil
}

// Union implements Factory.Union
func (f *cachingFactory) Union(a, b *Tags) *Tags {
	tags := make(map[string]struct{}, len(a.tags)+len(b.tags))
	for _, t := range a.tags {
		tags[t] = struct{}{}
	}
	for _, t := range b.tags {
		tags[t] = struct{}{}
	}
	return f.NewTagsFromMap(tags)
}

// DisjointUnion implements Factory.DisjoingUnion
func (f *cachingFactory) DisjointUnion(a, b *Tags) *Tags {
	hash := a.hash ^ b.hash
	return f.getCachedTags(byTagsetHashCache, hash, func() *Tags {

		tags := make([]string, len(a.tags)+len(b.tags))
		copy(tags[:len(a.tags)], a.tags)
		copy(tags[len(a.tags):], b.tags)

		hashes := make([]uint64, len(a.hashes)+len(b.hashes))
		copy(hashes[:len(a.hashes)], a.hashes)
		copy(hashes[len(a.hashes):], b.hashes)
		return &Tags{tags, hashes, hash}
	})
}

// getCachedTags implements Factory.getCachedTags
func (f *cachingFactory) getCachedTags(cacheID cacheID, hash uint64, miss func() *Tags) *Tags {
	cache := f.caches[cacheID]
	v, ok := cache[hash]
	if !ok {
		v = miss()
		cache[hash] = v
	}
	return v
}
