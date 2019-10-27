// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"runtime"

	"github.com/rakyll/go-numa"
)

func main() {
	if !numa.IsAvailable() {
		log.Fatalln("NUMA is not available on this machine.")
	}

	// Make sure the underlying OS thread
	// doesn't change because NUMA can only
	// lock the OS thread to a specific node.
	runtime.LockOSThread()

	// Runs the current goroutine always in node 0.
	numa.SetPreferred(0)
	// Do work...
}
