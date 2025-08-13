// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 25th February 2025
 * Updated: 13th August 2025
 */

package unix

import "strings"

const (
	PathElementSeparator = '/'
	PathSeparator        = ':'
)

func Basename(path string) string {
	ix := strings.LastIndexByte(path, '/')

	if ix < 0 {
		return path
	} else {
		return path[ix+1:]
	}
}

func ByteIsPathElementSeparator(c byte) bool {
	return '/' == c
}

func CharIsPathElementSeparator(c rune) bool {
	return '/' == c
}

func PathIsAbsolute(path string) bool {
	if 0 == len(path) {
		return false
	}

	return ByteIsPathElementSeparator(path[0])
}
