// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 25th February 2025
 * Updated: 21st August 2025
 */

package windows

import (
	"strings"
	"unicode"
)

const (
	PathElementSeparator    = '\\'
	PathElementSeparatorAlt = '/'
	PathSeparator           = ';'
)

var (
	_PathElementSeparators = []byte{
		PathElementSeparator,
		PathElementSeparatorAlt,
	}
)

func highestNonNegative(ix1, ix2 int) int {
	if ix1 < 0 {
		return ix2
	} else if ix2 < 0 {
		return ix1
	} else if ix1 < ix2 {
		return ix2
	} else {
		return ix1
	}
}

func Basename(path string) string {
	// TODO: reimplement this in terms of ANGoLS' LastIndexAnyByteAfter(path, _PathElementSeparators, -1)
	// TODO: we need to detect (partial) UNC so as not to match a UNC share

	ix_sep := strings.LastIndexByte(path, PathElementSeparator)
	ix_sep_alt := strings.LastIndexByte(path, PathElementSeparatorAlt)
	ix := highestNonNegative(ix_sep, ix_sep_alt)

	if ix < 0 {
		return path
	} else {
		return path[ix+1:]
	}
}

func ByteIsValidDriveLetter(c byte) bool {
	return unicode.IsLetter(rune(c))
}

func ByteIsPathElementSeparator(c byte) bool {
	switch c {
	case PathElementSeparator, PathElementSeparatorAlt:

		return true
	default:

		return false
	}
}

// Evaluates whether a byte represents an invalid character in a path.
func ByteIsInvalidInPath(c byte) bool {
	switch c {
	case '"', '*', '<', '>', '?', '|':

		return true
	default:

		return false
	}
}

func CharIsPathElementSeparator(c rune) bool {
	switch c {
	case PathElementSeparator, PathElementSeparatorAlt:

		return true
	default:

		return false
	}
}
