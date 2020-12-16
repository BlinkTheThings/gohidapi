package hidapi

/*
#include "../../../hidapi/hidapi/hidapi.h"
#cgo LDFLAGS: -L "../../build/" -lhidapi -lsetupapi
*/
import "C"

// VersionStr returns the version string of the hidapi library.
func VersionStr() string {
	return C.GoString(C.hid_version_str())
}
