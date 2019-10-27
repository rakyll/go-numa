// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux

// Package numa contains libnuma bindings for Go.
package numa

// #cgo LDFLAGS: -lnuma
//
// #include <numa.h>
import "C"

// IsAvailable returns true if NUMA is
// available.
func IsAvailable() bool {
	return int(C.numa_available()) != -1
}

// MaxNode returns the highest node number
// available on this machine.
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
