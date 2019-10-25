// +build linux

package numa

// #cgo LDFLAGS: -lnuma
//
// #include <numa.h>
import "C"

func IsNUMAAvailable() bool {
	return int(C.numa_available()) == 1
}

func MaxNode() int {
	return int(C.numa_max_node())
}

// Preferred returns the preferred node
// of the current thread.
func Preferred() int {
	return int(C.numa_preferred())
}

// SetPreferred sets the
// preferred node of the current thread.
func SetPreferred(node int) {
	C.numa_set_preferred(C.int(node))
}
