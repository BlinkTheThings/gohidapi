package hidapi

/*
#include "../../../hidapi/hidapi/hidapi.h"
#cgo LDFLAGS: -lhidapi
#cgo windows LDFLAGS: -L "../../build/windows" -lsetupapi
#cgo linux LDFLAGS: -L "../../build/linux" -ludev
*/
import "C"

type DeviceInfo struct {
	VendorID  uint16
	ProductID uint16
	Path      string
	Release   uint16
	Interface int
	Usage     uint16
	UsagePage uint16
}

// VersionStr returns the version string of the HIDAPI library.
func VersionStr() string {
	return C.GoString(C.hid_version_str())
}

// Version returns the major, minor, and patch version numbers of the HIDAPI library.
func Version() (major int, minor int, patch int) {
	v := C.hid_version()
	major, minor, patch = int(v.major), int(v.minor), int(v.patch)
	return
}

// Init initializes the HIDAPI library.
func Init() int {
	return int(C.hid_init())
}

// Exit finalizes the HIDAPI library.
func Exit() int {
	return int(C.hid_exit())
}

func Enumerate(vid int, pid int) (devices []DeviceInfo) {

	devices = make([]DeviceInfo, 0)

	deviceList := C.hid_enumerate(C.ushort(vid), C.ushort(pid))

	curDev := deviceList

	for curDev != nil {
		devices = append(devices, DeviceInfo{
			VendorID:  uint16(curDev.vendor_id),
			ProductID: uint16(curDev.product_id),
			Path:      C.GoString(curDev.path),
			Release:   uint16(curDev.release_number),
			Interface: int(curDev.interface_number),
			Usage:     uint16(curDev.usage),
			UsagePage: uint16(curDev.usage_page),
		})

		curDev = curDev.next
	}

	C.hid_free_enumeration(deviceList)

	return
}
