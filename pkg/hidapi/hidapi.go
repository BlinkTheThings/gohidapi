package hidapi

/*
#include "../../../hidapi/hidapi/hidapi.h"
#cgo LDFLAGS: -lhidapi
#cgo windows LDFLAGS: -L "../../build/windows" -lsetupapi
#cgo linux LDFLAGS: -L "../../build/linux" -ludev
*/
import "C"

// VersionStr returns the version string of the hidapi library.
func VersionStr() string {
	return C.GoString(C.hid_version_str())
}

// Version returns the major, minor, and patch version numbers of the hidapi library.
func Version() (major int, minor int, patch int) {
	v := C.hid_version()
	major, minor, patch = int(v.major), int(v.minor), int(v.patch)
	return
}
