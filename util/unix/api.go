// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 25th February 2025
 * Updated: 21st August 2025
 */

package unix

import (
	"strings"
)

const (
	PathElementSeparator = '/'
	PathSeparator        = ':'
)

func Basename(path string) string {
	ix := strings.LastIndexByte(path, PathElementSeparator)

	if ix < 0 {
		return path
	} else {
		return path[ix+1:]
	}
}

func ByteIsPathElementSeparator(c byte) bool {
	return PathElementSeparator == c
}

func CharIsPathElementSeparator(c rune) bool {
	return PathElementSeparator == c
}

// Evaluates whether a byte represents an invalid character in a path.
func ByteIsInvalidInPath(c byte) bool {
	switch c {
	case '*', '<', '>', '?', '|':

		return true
	default:

		return false
	}
}
