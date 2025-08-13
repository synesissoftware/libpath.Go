// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

/*
 * Created: 25th February 2025
 * Updated: 13th August 2025
 */

package util

import (
	"github.com/synesissoftware/libpath.Go/util/windows"
)

func Basename(path string) string {
	return windows.Basename(path)
}

func ByteIsPathElementSeparator(c byte) bool {
	return windows.ByteIsPathElementSeparator(c)
}

func CharIsPathElementSeparator(c rune) bool {
	return windows.CharIsPathElementSeparator(c)
}

func PathIsAbsolute(path string) bool {
	return windows.PathIsAbsolute(path)
}
