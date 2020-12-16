package hidapi

/*
#include "../../../hidapi/hidapi/hidapi.h"
#cgo LDFLAGS: -lhidapi
#cgo windows LDFLAGS: -L "../../build/windows" -lsetupapi
#cgo linux LDFLAGS: -L "../../build/linux" -ludev
*/
import "C"

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
