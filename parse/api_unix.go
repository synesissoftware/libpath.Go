// Copyright 2019-2025, Matthew Wilson and Synesis Information Systems. All
// rights reserved. Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

/*
 * Created: 19th August 2025
 * Updated: 21st August 2025
 */

package parse

import (
	types "github.com/synesissoftware/libpath.Go/parse/common"
	os_api "github.com/synesissoftware/libpath.Go/parse/unix"
)

func ParsePathString(path string, reference_directory string) (types.PathDescriptor, error) {
	return os_api.ParsePathStringFlags(path, reference_directory, types.ParseFlags_None)
}

func ParsePathStringFlags(path string, reference_directory string, parseFlags types.ParseFlags) (types.PathDescriptor, error) {
	return os_api.ParsePathStringFlags(path, reference_directory, parseFlags)
}
