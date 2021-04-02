package pb

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *ClientGroupedStats) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Service":
			z.Service, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Resource":
			z.Resource, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Resource")
				return
			}
		case "HTTPStatusCode":
			z.HTTPStatusCode, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "HTTPStatusCode")
				return
			}
		case "Type":
			z.Type, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "DBType":
			z.DBType, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "DBType")
				return
			}
		case "Hits":
			z.Hits, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Hits")
				return
			}
		case "Errors":
			z.Errors, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Errors")
				return
			}
		case "Duration":
			z.Duration, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "OkSummary":
			z.OkSummary, err = dc.ReadBytes(z.OkSummary)
			if err != nil {
				err = msgp.WrapError(err, "OkSummary")
				return
			}
		case "ErrorSummary":
			z.ErrorSummary, err = dc.ReadBytes(z.ErrorSummary)
			if err != nil {
				err = msgp.WrapError(err, "ErrorSummary")
				return
			}
		case "Synthetics":
			z.Synthetics, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "Synthetics")
				return
			}
		case "TopLevelHits":
			z.TopLevelHits, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "TopLevelHits")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientGroupedStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 13
	// write "Service"
	err = en.Append(0x8d, 0xa7, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Service)
	if err != nil {
		err = msgp.WrapError(err, "Service")
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "Resource"
	err = en.Append(0xa8, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Resource)
	if err != nil {
		err = msgp.WrapError(err, "Resource")
		return
	}
	// write "HTTPStatusCode"
	err = en.Append(0xae, 0x48, 0x54, 0x54, 0x50, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.HTTPStatusCode)
	if err != nil {
		err = msgp.WrapError(err, "HTTPStatusCode")
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Type)
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "DBType"
	err = en.Append(0xa6, 0x44, 0x42, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.DBType)
	if err != nil {
		err = msgp.WrapError(err, "DBType")
		return
	}
	// write "Hits"
	err = en.Append(0xa4, 0x48, 0x69, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Hits)
	if err != nil {
		err = msgp.WrapError(err, "Hits")
		return
	}
	// write "Errors"
	err = en.Append(0xa6, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Errors)
	if err != nil {
		err = msgp.WrapError(err, "Errors")
		return
	}
	// write "Duration"
	err = en.Append(0xa8, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Duration)
	if err != nil {
		err = msgp.WrapError(err, "Duration")
		return
	}
	// write "OkSummary"
	err = en.Append(0xa9, 0x4f, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.OkSummary)
	if err != nil {
		err = msgp.WrapError(err, "OkSummary")
		return
	}
	// write "ErrorSummary"
	err = en.Append(0xac, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.ErrorSummary)
	if err != nil {
		err = msgp.WrapError(err, "ErrorSummary")
		return
	}
	// write "Synthetics"
	err = en.Append(0xaa, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Synthetics)
	if err != nil {
		err = msgp.WrapError(err, "Synthetics")
		return
	}
	// write "TopLevelHits"
	err = en.Append(0xac, 0x54, 0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x69, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.TopLevelHits)
	if err != nil {
		err = msgp.WrapError(err, "TopLevelHits")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientGroupedStats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 13
	// string "Service"
	o = append(o, 0x8d, 0xa7, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65)
	o = msgp.AppendString(o, z.Service)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Resource"
	o = append(o, 0xa8, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65)
	o = msgp.AppendString(o, z.Resource)
	// string "HTTPStatusCode"
	o = append(o, 0xae, 0x48, 0x54, 0x54, 0x50, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65)
	o = msgp.AppendUint32(o, z.HTTPStatusCode)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "DBType"
	o = append(o, 0xa6, 0x44, 0x42, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.DBType)
	// string "Hits"
	o = append(o, 0xa4, 0x48, 0x69, 0x74, 0x73)
	o = msgp.AppendUint64(o, z.Hits)
	// string "Errors"
	o = append(o, 0xa6, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73)
	o = msgp.AppendUint64(o, z.Errors)
	// string "Duration"
	o = append(o, 0xa8, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint64(o, z.Duration)
	// string "OkSummary"
	o = append(o, 0xa9, 0x4f, 0x6b, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79)
	o = msgp.AppendBytes(o, z.OkSummary)
	// string "ErrorSummary"
	o = append(o, 0xac, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79)
	o = msgp.AppendBytes(o, z.ErrorSummary)
	// string "Synthetics"
	o = append(o, 0xaa, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x73)
	o = msgp.AppendBool(o, z.Synthetics)
	// string "TopLevelHits"
	o = append(o, 0xac, 0x54, 0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x69, 0x74, 0x73)
	o = msgp.AppendUint64(o, z.TopLevelHits)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientGroupedStats) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Service":
			z.Service, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Service")
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Resource":
			z.Resource, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Resource")
				return
			}
		case "HTTPStatusCode":
			z.HTTPStatusCode, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "HTTPStatusCode")
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Type")
				return
			}
		case "DBType":
			z.DBType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DBType")
				return
			}
		case "Hits":
			z.Hits, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Hits")
				return
			}
		case "Errors":
			z.Errors, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Errors")
				return
			}
		case "Duration":
			z.Duration, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "OkSummary":
			z.OkSummary, bts, err = msgp.ReadBytesBytes(bts, z.OkSummary)
			if err != nil {
				err = msgp.WrapError(err, "OkSummary")
				return
			}
		case "ErrorSummary":
			z.ErrorSummary, bts, err = msgp.ReadBytesBytes(bts, z.ErrorSummary)
			if err != nil {
				err = msgp.WrapError(err, "ErrorSummary")
				return
			}
		case "Synthetics":
			z.Synthetics, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Synthetics")
				return
			}
		case "TopLevelHits":
			z.TopLevelHits, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TopLevelHits")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientGroupedStats) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Service) + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.StringPrefixSize + len(z.Resource) + 15 + msgp.Uint32Size + 5 + msgp.StringPrefixSize + len(z.Type) + 7 + msgp.StringPrefixSize + len(z.DBType) + 5 + msgp.Uint64Size + 7 + msgp.Uint64Size + 9 + msgp.Uint64Size + 10 + msgp.BytesPrefixSize + len(z.OkSummary) + 13 + msgp.BytesPrefixSize + len(z.ErrorSummary) + 11 + msgp.BoolSize + 13 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientStatsBucket) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Start":
			z.Start, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "Duration":
			z.Duration, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]ClientGroupedStats, zb0002)
			}
			for za0001 := range z.Stats {
				err = z.Stats[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientStatsBucket) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Start"
	err = en.Append(0x83, 0xa5, 0x53, 0x74, 0x61, 0x72, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Start)
	if err != nil {
		err = msgp.WrapError(err, "Start")
		return
	}
	// write "Duration"
	err = en.Append(0xa8, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Duration)
	if err != nil {
		err = msgp.WrapError(err, "Duration")
		return
	}
	// write "Stats"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001 := range z.Stats {
		err = z.Stats[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientStatsBucket) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Start"
	o = append(o, 0x83, 0xa5, 0x53, 0x74, 0x61, 0x72, 0x74)
	o = msgp.AppendUint64(o, z.Start)
	// string "Duration"
	o = append(o, 0xa8, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendUint64(o, z.Duration)
	// string "Stats"
	o = append(o, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Stats)))
	for za0001 := range z.Stats {
		o, err = z.Stats[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientStatsBucket) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Start":
			z.Start, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Start")
				return
			}
		case "Duration":
			z.Duration, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Duration")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]ClientGroupedStats, zb0002)
			}
			for za0001 := range z.Stats {
				bts, err = z.Stats[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientStatsBucket) Msgsize() (s int) {
	s = 1 + 6 + msgp.Uint64Size + 9 + msgp.Uint64Size + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Stats {
		s += z.Stats[za0001].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ClientStatsPayload) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hostname":
			z.Hostname, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Hostname")
				return
			}
		case "Env":
			z.Env, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Env")
				return
			}
		case "Version":
			z.Version, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]ClientStatsBucket, zb0002)
			}
			for za0001 := range z.Stats {
				err = z.Stats[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		case "Lang":
			z.Lang, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Lang")
				return
			}
		case "TracerVersion":
			z.TracerVersion, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "TracerVersion")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ClientStatsPayload) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "Hostname"
	err = en.Append(0x86, 0xa8, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Hostname)
	if err != nil {
		err = msgp.WrapError(err, "Hostname")
		return
	}
	// write "Env"
	err = en.Append(0xa3, 0x45, 0x6e, 0x76)
	if err != nil {
		return
	}
	err = en.WriteString(z.Env)
	if err != nil {
		err = msgp.WrapError(err, "Env")
		return
	}
	// write "Version"
	err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Version)
	if err != nil {
		err = msgp.WrapError(err, "Version")
		return
	}
	// write "Stats"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001 := range z.Stats {
		err = z.Stats[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	// write "Lang"
	err = en.Append(0xa4, 0x4c, 0x61, 0x6e, 0x67)
	if err != nil {
		return
	}
	err = en.WriteString(z.Lang)
	if err != nil {
		err = msgp.WrapError(err, "Lang")
		return
	}
	// write "TracerVersion"
	err = en.Append(0xad, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.TracerVersion)
	if err != nil {
		err = msgp.WrapError(err, "TracerVersion")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ClientStatsPayload) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "Hostname"
	o = append(o, 0x86, 0xa8, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Hostname)
	// string "Env"
	o = append(o, 0xa3, 0x45, 0x6e, 0x76)
	o = msgp.AppendString(o, z.Env)
	// string "Version"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Version)
	// string "Stats"
	o = append(o, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Stats)))
	for za0001 := range z.Stats {
		o, err = z.Stats[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	// string "Lang"
	o = append(o, 0xa4, 0x4c, 0x61, 0x6e, 0x67)
	o = msgp.AppendString(o, z.Lang)
	// string "TracerVersion"
	o = append(o, 0xad, 0x54, 0x72, 0x61, 0x63, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.TracerVersion)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ClientStatsPayload) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Hostname":
			z.Hostname, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Hostname")
				return
			}
		case "Env":
			z.Env, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Env")
				return
			}
		case "Version":
			z.Version, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Version")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]ClientStatsBucket, zb0002)
			}
			for za0001 := range z.Stats {
				bts, err = z.Stats[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		case "Lang":
			z.Lang, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Lang")
				return
			}
		case "TracerVersion":
			z.TracerVersion, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TracerVersion")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ClientStatsPayload) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Hostname) + 4 + msgp.StringPrefixSize + len(z.Env) + 8 + msgp.StringPrefixSize + len(z.Version) + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Stats {
		s += z.Stats[za0001].Msgsize()
	}
	s += 5 + msgp.StringPrefixSize + len(z.Lang) + 14 + msgp.StringPrefixSize + len(z.TracerVersion)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *StatsPayload) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "AgentHostname":
			z.AgentHostname, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "AgentHostname")
				return
			}
		case "AgentEnv":
			z.AgentEnv, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "AgentEnv")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]ClientStatsPayload, zb0002)
			}
			for za0001 := range z.Stats {
				err = z.Stats[za0001].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *StatsPayload) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "AgentHostname"
	err = en.Append(0x83, 0xad, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.AgentHostname)
	if err != nil {
		err = msgp.WrapError(err, "AgentHostname")
		return
	}
	// write "AgentEnv"
	err = en.Append(0xa8, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76)
	if err != nil {
		return
	}
	err = en.WriteString(z.AgentEnv)
	if err != nil {
		err = msgp.WrapError(err, "AgentEnv")
		return
	}
	// write "Stats"
	err = en.Append(0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Stats)))
	if err != nil {
		err = msgp.WrapError(err, "Stats")
		return
	}
	for za0001 := range z.Stats {
		err = z.Stats[za0001].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StatsPayload) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "AgentHostname"
	o = append(o, 0x83, 0xad, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.AgentHostname)
	// string "AgentEnv"
	o = append(o, 0xa8, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76)
	o = msgp.AppendString(o, z.AgentEnv)
	// string "Stats"
	o = append(o, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Stats)))
	for za0001 := range z.Stats {
		o, err = z.Stats[za0001].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Stats", za0001)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StatsPayload) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "AgentHostname":
			z.AgentHostname, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AgentHostname")
				return
			}
		case "AgentEnv":
			z.AgentEnv, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AgentEnv")
				return
			}
		case "Stats":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Stats")
				return
			}
			if cap(z.Stats) >= int(zb0002) {
				z.Stats = (z.Stats)[:zb0002]
			} else {
				z.Stats = make([]ClientStatsPayload, zb0002)
			}
			for za0001 := range z.Stats {
				bts, err = z.Stats[za0001].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Stats", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StatsPayload) Msgsize() (s int) {
	s = 1 + 14 + msgp.StringPrefixSize + len(z.AgentHostname) + 9 + msgp.StringPrefixSize + len(z.AgentEnv) + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Stats {
		s += z.Stats[za0001].Msgsize()
	}
	return
}