// +build windows,npm

package tracer

import (
	"fmt"
	"github.com/DataDog/datadog-agent/pkg/network/driver"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/windows"
	"testing"
	"unsafe"
)

func dnsSupported(t *testing.T) bool {
	return true
}

func httpSupported(t *testing.T) bool {
	return false
}

func TestDriverRequiresPath(t *testing.T) {
	p, err := windows.UTF16PtrFromString(driver.DeviceName)
	assert.Nil(t, err)
	h, err := windows.CreateFile(p,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		0,
		windows.Handle(0))
	if err == nil {
		defer windows.CloseHandle(h)
	}
	assert.NotNil(t, err)
}

func TestDriverCanOpenExpectedPaths(t *testing.T) {
	for _, pathext := range driver.HandleTypeToPathName {
		fullpath := driver.DeviceName + `\` + pathext
		p, err := windows.UTF16PtrFromString(fullpath)
		assert.Nil(t, err)
		h, err := windows.CreateFile(p,
			windows.GENERIC_READ|windows.GENERIC_WRITE,
			windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
			nil,
			windows.OPEN_EXISTING,
			0,
			windows.Handle(0))
		if err == nil {
			defer windows.CloseHandle(h)
		}
		assert.Nil(t, err)
	}
}

func createHandleForHandleType(t driver.HandleType) (windows.Handle, error) {
	pathext, ok := driver.HandleTypeToPathName[t]
	if !ok {
		return 0, fmt.Errorf("Unknown Handle type %v", t)
	}
	fullpath := driver.DeviceName + `\` + pathext
	p, err := windows.UTF16PtrFromString(fullpath)
	if err != nil {
		return 0, err
	}
	h, err := windows.CreateFile(p,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		0,
		windows.Handle(0))
	if err != nil {
		return 0, err
	}
	return h, nil
}
func TestDriverWillAcceptFilters(t *testing.T) {
	simplefilter := driver.FilterDefinition{
		FilterVersion: driver.Signature,
		Size:          driver.FilterDefinitionSize,
		Direction:     driver.DirectionInbound,
		FilterLayer:   driver.LayerTransport,
		Af:            windows.AF_INET,
		Protocol:      windows.IPPROTO_TCP,
	}

	t.Run("Test flow handle will accept flow filter", func(t *testing.T) {
		var id int64
		h, err := createHandleForHandleType(driver.FlowHandle)
		if err == nil {
			defer windows.CloseHandle(h)
		}
		assert.Nil(t, err)

		err = windows.DeviceIoControl(h,
			driver.SetFlowFilterIOCTL,
			(*byte)(unsafe.Pointer(&simplefilter)),
			uint32(unsafe.Sizeof(simplefilter)),
			(*byte)(unsafe.Pointer(&id)),
			uint32(unsafe.Sizeof(id)), nil, nil)
		assert.Nil(t, err)
	})
	t.Run("Test flow handle will accept transport filter", func(t *testing.T) {
		var id int64
		h, err := createHandleForHandleType(driver.DataHandle)
		if err == nil {
			defer windows.CloseHandle(h)
		}
		assert.Nil(t, err)

		err = windows.DeviceIoControl(h,
			driver.SetDataFilterIOCTL,
			(*byte)(unsafe.Pointer(&simplefilter)),
			uint32(unsafe.Sizeof(simplefilter)),
			(*byte)(unsafe.Pointer(&id)),
			uint32(unsafe.Sizeof(id)), nil, nil)
		assert.Nil(t, err)
	})
}

func TestDriverWillNotAcceptMismatchedFilters(t *testing.T) {
	simplefilter := driver.FilterDefinition{
		FilterVersion: driver.Signature,
		Size:          driver.FilterDefinitionSize,
		Direction:     driver.DirectionInbound,
		FilterLayer:   driver.LayerTransport,
		Af:            windows.AF_INET,
		Protocol:      windows.IPPROTO_TCP,
	}

	t.Run("Test flow handle will not accept data filter on flow handle", func(t *testing.T) {
		var id int64
		h, err := createHandleForHandleType(driver.FlowHandle)
		if err == nil {
			defer windows.CloseHandle(h)
		}
		assert.Nil(t, err)

		err = windows.DeviceIoControl(h,
			driver.SetDataFilterIOCTL,
			(*byte)(unsafe.Pointer(&simplefilter)),
			uint32(unsafe.Sizeof(simplefilter)),
			(*byte)(unsafe.Pointer(&id)),
			uint32(unsafe.Sizeof(id)), nil, nil)
		assert.NotNil(t, err)
	})
	t.Run("Test flow handle will not accept flow filter on transport handle", func(t *testing.T) {
		var id int64
		h, err := createHandleForHandleType(driver.DataHandle)
		if err == nil {
			defer windows.CloseHandle(h)
		}
		assert.Nil(t, err)

		err = windows.DeviceIoControl(h,
			driver.SetFlowFilterIOCTL,
			(*byte)(unsafe.Pointer(&simplefilter)),
			uint32(unsafe.Sizeof(simplefilter)),
			(*byte)(unsafe.Pointer(&id)),
			uint32(unsafe.Sizeof(id)), nil, nil)
		assert.NotNil(t, err)
	})
}
