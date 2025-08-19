// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 25th February 2025
 * Updated: 19th August 2025
 */

package windows

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	PathElementSeparator    = '\\'
	PathElementSeparatorAlt = '/'
	PathSeparator           = ';'
)

// Path classification.
type Classification int64

const (
	InvalidSlashRuns    Classification = -3
	InvalidChars                       = -2
	Invalid                            = -1
	Unknown                            = 0
	Empty                              = 1
	Relative                           = 2
	SlashRooted                        = 3
	DriveLetterRelative                = 4
	DriveLetterRooted                  = 5
	UncIncomplete                      = 6
	UncRooted                          = 7
	HomeRooted                         = 8
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
	case DriveLetterRelative:
		return "DriveLetterRelative"
	case DriveLetterRooted:
		return "DriveLetterRooted"
	case UncIncomplete:
		return "UncIncomplete"
	case UncRooted:
		return "UncRooted"
	case HomeRooted:
		return "HomeRooted"
	default:
		return fmt.Sprintf("invalid %T (value=%d)", classification, classification)
	}
}

// Path parse flags.
type ParseFlags uint64

const (
	ParseFlags_None                         ParseFlags = 0
	ParseFlags_IgnoreSlashRuns              ParseFlags = 0x00000001
	ParseFlags_IgnoreInvalidChars           ParseFlags = 0x00000002
	ParseFlags_RecogniseTildeHome           ParseFlags = 0x00000004
	ParseFlags_AssumeDirectory              ParseFlags = 0x00000008
	ParseFlags_IgnoreInvalidCharsInLongPath ParseFlags = 0x00000080
)

var (
	_PathElementSeparators = []rune{
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
	// TODO: reimplement this in terms of ANGoLS' LastIndexAnyByte
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

	if ByteIsPathElementSeparator(path[0]) {
		return SlashRooted, path[:1], path[1:], -1
	}

	if len(path) >= 2 {
		if ByteIsValidDriveLetter(path[0]) {
			if ByteIsInvalidInPath(path[1]) {
				return InvalidChars, "", "", 1
			}
			if ':' == path[1] {
				if len(path) == 2 {
					return DriveLetterRelative, path, "", -1
				} else {
					if ByteIsInvalidInPath(path[2]) {
						return InvalidChars, "", "", 2
					}
					if ByteIsPathElementSeparator(path[2]) {
						return DriveLetterRooted, path[:3], path[3:], -1
					} else {
						return DriveLetterRelative, path[:2], path[2:], -1
					}
				}
			}
		}
	}

	if path[0] == '~' {
		if ParseFlags_RecogniseTildeHome == (parseFlags & ParseFlags_RecogniseTildeHome) {
			if len(path) == 1 {
				return HomeRooted, path, "", -1
			}
			if ByteIsInvalidInPath(path[1]) {
				return InvalidChars, "", "", 1
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
