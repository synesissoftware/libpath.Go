// Copyright 2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Created: 18th August 2025
 * Updated: 21st August 2025
 */

package common

// Path parse flags.
type ParseFlags uint64

const (
	ParseFlags_None                         ParseFlags = 0
	ParseFlags_IgnoreSlashRuns              ParseFlags = 0x00000001
	ParseFlags_IgnoreInvalidChars           ParseFlags = 0x00000002
	ParseFlags_RecogniseTildeHome           ParseFlags = 0x00000004
	ParseFlags_RecognizeTildeHome                      = ParseFlags_RecogniseTildeHome
	ParseFlags_AssumeDirectory              ParseFlags = 0x00000008
	ParseFlags_IgnoreInvalidCharsInLongPath ParseFlags = 0x00000080
)
