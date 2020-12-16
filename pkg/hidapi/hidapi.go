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
