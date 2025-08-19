// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 25th February 2025
 * Updated: 19th August 2025
 */

package unix

import (
	"fmt"
	"strings"
)

const (
	PathElementSeparator = '/'
	PathSeparator        = ':'
)

// Path classification.
type Classification int64

const (
	InvalidSlashRuns Classification = -3
	InvalidChars                    = -2
	Invalid                         = -1
	Unknown                         = 0
	Empty                           = 1
	Relative                        = 2
	SlashRooted                     = 3
	_Reserved1                      = 4
	_Reserved2                      = 5
	_Reserved3                      = 6
	_Reserved4                      = 7
	HomeRooted                      = 8
)

func (classification Classification) String() string {
	switch classification {
	case InvalidSlashRuns:
		return "InvalidSlashRuns"
	case InvalidChars:
		return "InvalidChars"
	case Invalid:
		return "Invalid"
	case Empty:
		return "Empty"
	case Relative:
		return "Relative"
	case SlashRooted:
		return "SlashRooted"
	case HomeRooted:
		return "HomeRooted"
	default:
		return fmt.Sprintf("invalid %T (value=%d)", classification, classification)
	}
}

// Path parse flags.
type ParseFlags uint64

const (
	ParseFlags_None               ParseFlags = 0
	ParseFlags_IgnoreSlashRuns    ParseFlags = 0x00000001
	ParseFlags_IgnoreInvalidChars ParseFlags = 0x00000002
	ParseFlags_RecogniseTildeHome ParseFlags = 0x00000004
	ParseFlags_AssumeDirectory    ParseFlags = 0x00000008
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

// Evaluates whether a byte represents an invalid character in a path.
func ByteIsInvalidInPath(c byte) bool {
	switch c {
	case '*', '<', '>', '?', '|':
		return true
	default:
		return false
	}
}

// Examines the path to the degree necessary to be able to classify it.
//
// # Parameters:
// - `path` - the given path to be classified;
// - `parseFlags` - flags that moderate the classification;
//
// # Returns:
// `(classification : Classification, root : string, path_root_stripped : string, first_bad_char_index : int)`
func ClassifyRoot(path string, parseFlags ParseFlags) (Classification, string, string, int) {

	if len(path) == 0 {
		return Empty, "", "", 0
	}

	if path[0] == '/' {
		return SlashRooted, "/", path[1:], -1
	}

	if path[0] == '~' {
		if ParseFlags_RecogniseTildeHome == (parseFlags & ParseFlags_RecogniseTildeHome) {
			if len(path) == 1 {
				return HomeRooted, path, "", -1
			}
			if ByteIsPathElementSeparator(path[1]) {
				return HomeRooted, path[:2], path[2:], -1
			}
		}
	}

	if ByteIsInvalidInPath(path[0]) {
		return InvalidChars, "", "", 0
	}

	return Relative, "", path, -1
}

func PathIsAbsolute(path string) bool {
	if 0 == len(path) {
		return false
	}

	return ByteIsPathElementSeparator(path[0])
}
