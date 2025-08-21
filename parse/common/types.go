// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 25th February 2025
 * Updated: 21st August 2025
 */

package common

import "fmt"

// Type describing any file-system path in terms of its specific attributes,
// including its root, directory, and entry elements.
//
// When parsing a path string, then the fullest establisable form is
// obtained according to the following:
//
//	1 if the path is absolute, then the path is interpreted entirely without
//	  reference to any other information;
//	2 if the path is relative and a reference directory is provided, then
//	  the path is interpreted as if relative to the reference;
//	3 if the path is relative and a reference directory is not provided,
//	  then the path is interpreted as is;
//
// Subject to the above, the input is converted to its fullest establishable
// form for the purposes of parsing. Hence, for example the `Location` field
// is usually the full path of the input path's directory, except in case 3
// above.
type PathDescriptor struct {
	Classification        Classification
	Input                 string   // The original input string to the parsing.
	FullPath              string   // The fullest esablishable form of the input.
	Location              string   // The fullest esablishable form of the location of the entry, which is everything up-to-and-including the last (if any) path-name separator.
	Root                  string   // The path root, if present.
	Directory             string   // The path directory, which excluses the `Root` (if any) and the `EntryName` (if any).
	DirectoryParts        []string // Array of the `directory` elements, split on the path-name separator.
	NumDotsDirectoryParts int      // The number of dots directories in the descriptor
	EntryName             string   // The "file part", if any, which occurs after the last (if any) path-name separator. In many respects this is the same as Unix "base name", except that parsing a path with a trailing path-name-separator will _always_ results in an empty `EntryName`
	Stem                  string   // The stem of the "file part", if any.
	Extension             string   // The extension of the "file part", if any.
}

func (pd *PathDescriptor) IsAbsolute() bool {
	return len(pd.Root) != 0
}

func (pd *PathDescriptor) IsCanonical() bool {
	return pd.NumDotsDirectoryParts == 0
}

func (pd PathDescriptor) String() string {

	return fmt.Sprintf(""+
		"<%T{"+
		"Classification=%v, "+
		"Input=%q, "+
		"FullPath=%q, "+
		"Location=%q, "+
		"Root=%q, "+
		"Directory=%q, "+
		"DirectoryParts=%q, "+
		"NumDotsDirectoryParts=%d, "+
		"EntryName=%q, "+
		"Stem=%q, "+
		"Extension=%q}"+
		">",
		pd,
		pd.Classification,
		pd.Input,
		pd.FullPath,
		pd.Location,
		pd.Root,
		pd.Directory,
		pd.DirectoryParts,
		pd.NumDotsDirectoryParts,
		pd.EntryName,
		pd.Stem,
		pd.Extension,
	)
}

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
	case Unknown:
		return "Unknown"
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

type ParseFailure struct {
	Classification    Classification
	FirstBadCharIndex int
	Input             string
	Qualifier         string
}

func (pf *ParseFailure) Error() string {

	if len(pf.Qualifier) != 0 {
		return fmt.Sprintf("Parse failure (%s) at/around character index %d: %s", pf.Classification, pf.FirstBadCharIndex, pf.Qualifier)
	} else {
		return fmt.Sprintf("Parse failure (%s) at/around character index %d", pf.Classification, pf.FirstBadCharIndex)
	}
}
