// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

/*
 * Created: 25th February 2025
 * Updated: 13th August 2025
 */

package util

import (
	"github.com/synesissoftware/libpath.Go/util/unix"
)

func Basename(path string) string {
	return unix.Basename(path)
}

func ByteIsPathElementSeparator(c byte) bool {
	return unix.ByteIsPathElementSeparator(c)
}

func CharIsPathElementSeparator(c rune) bool {
	return unix.CharIsPathElementSeparator(c)
}

func PathIsAbsolute(path string) bool {
	return unix.PathIsAbsolute(path)
}
