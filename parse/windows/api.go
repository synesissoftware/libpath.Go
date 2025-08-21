package windows

import (
	parse_utils "github.com/synesissoftware/libpath.Go/internal/parse_utils"
	types "github.com/synesissoftware/libpath.Go/parse/common"
	util_common "github.com/synesissoftware/libpath.Go/util"
	os_util "github.com/synesissoftware/libpath.Go/util/windows"

	angols_strings "github.com/synesissoftware/ANGoLS/strings"

	"fmt"
	std_strings "strings"
)

var (
	_PathElementSeparators = []byte{
		os_util.PathElementSeparator,
		os_util.PathElementSeparatorAlt,
	}
)

func isDotsDirectoryPart(s string) bool {
	ix_no := angols_strings.IndexNotAnyAfter(s, `./\`, -1)

	return -1 == ix_no
}

func countDotsDirectoryParts(directoryParts []string) int {
	return parse_utils.CountDotsDirectoryParts(directoryParts, isDotsDirectoryPart)
}

func elementEndsWithPathNameSeparator(s string) bool {
	return parse_utils.ElementEndsWithByte(s, os_util.ByteIsPathElementSeparator)
}

func simplePathSplitFully(path string) ([]string, error) {
	return angols_strings.SplitAfterAnyBytes(path, _PathElementSeparators), nil
}

func simplePathSplit(path string) (string, []string, string, error) {
	splits, err := simplePathSplitFully(path)
	if err != nil {
		return "", nil, "", nil
	}

	dp_from := 0
	dp_count := len(splits)

	first := splits[0]
	last := splits[len(splits)-1]

	var root string
	var directory_parts []string
	var file_part string

	if elementEndsWithPathNameSeparator(last) {
		file_part = ""
	} else {
		file_part = last
		dp_count -= 1
	}

	if parse_utils.ElementIsRooted(first, os_util.ByteIsPathElementSeparator) {
		root = first
		dp_from += 1
	} else {
		root = ""
	}

	directory_parts = splits[dp_from:dp_count]

	return root, directory_parts, file_part, nil
}

func simplePathJoin(elems ...string) string {
	var b std_strings.Builder

	b.Grow(256) // just a guess for now

	last_has_trailing := false
	for _, s := range elems {
		if 0 == len(s) {
			// skip
		} else {
			if !last_has_trailing && 0 != b.Len() {
				b.WriteByte('/')
			}
			b.WriteString(s)
			if os_util.ByteIsPathElementSeparator(s[len(s)-1]) {
				last_has_trailing = true
			} else {
				last_has_trailing = false
			}
		}
	}

	return b.String()
}

func splitPathRootStripped(pathRootStripped string) (directory string, directoryParts []string, entryBasename string, err error) {

	splits, err := simplePathSplitFully(pathRootStripped)
	if err != nil {
		return "", nil, "", nil
	}

	if len(splits) > 0 {

		last := splits[len(splits)-1]
		dp_count := len(splits)

		if elementEndsWithPathNameSeparator(last) {
			directory = pathRootStripped
			entryBasename = ""
		} else {
			directory = pathRootStripped[:len(pathRootStripped)-len(last)]
			entryBasename = last
			dp_count -= 1
		}

		directoryParts = splits[:dp_count]

		err = nil
	}

	return
}

func createPathDescriptor(path string, ref_dir string, parseFlags types.ParseFlags) (types.PathDescriptor, error) {

	// account for four cases of `path` and `ref_dir`
	//
	//  |   path   |  refDir  |      output      |
	//  | -------- | -------- | ---------------- |
	//  | empty    |     -    | (empty)          |
	//  | absolute |     -    | path             |
	//  | relative | !empty   | refDir + path    |
	//
	// 1. `path` is absolute => desc. based entirely on `path`
	// 2. `path` is empty, `ref_dir` is empty => empty desc.;
	// 3. `path` is relative, `ref_dir` is empty => desc. based entirely on `path`
	// 4. `path` is relative, `ref_dir` is !empty => desc. based on both;

	classification, root, pathRootStripped, firstBadCharIndex := ClassifyRoot(path, parseFlags)

	switch classification {
	case types.InvalidSlashRuns:

		return types.PathDescriptor{}, &types.ParseFailure{
			Classification:    classification,
			FirstBadCharIndex: firstBadCharIndex,
			Input:             path,
			Qualifier:         "",
		}
	case types.InvalidChars, types.Invalid:

		return types.PathDescriptor{}, &types.ParseFailure{
			Classification:    classification,
			FirstBadCharIndex: firstBadCharIndex,
			Input:             path,
			Qualifier:         "",
		}
	case types.Unknown:

		panic(fmt.Errorf("classification %s should never happen", &classification))
	case types.Empty:

		return types.PathDescriptor{}, nil
	case types.UncRooted, types.HomeRooted, types.DriveLetterRooted:

		directory, directoryParts, entryBasename, err := splitPathRootStripped(pathRootStripped)

		if err != nil {
			return types.PathDescriptor{}, err
		} else {

			entryStem, entryExtension := util_common.SplitBasename(entryBasename)

			entryBasename := os_util.Basename(entryBasename)

			location := path[:len(path)-len(entryBasename)]

			return types.PathDescriptor{
				Classification:        classification,
				Input:                 path,
				FullPath:              path,
				Location:              location,
				Root:                  root,
				Directory:             directory,
				DirectoryParts:        directoryParts,
				NumDotsDirectoryParts: countDotsDirectoryParts(directoryParts),
				EntryName:             entryBasename,
				Stem:                  entryStem,
				Extension:             entryExtension,
			}, nil
		}
	}

	full_path := path

	if 0 != len(ref_dir) {
		full_path = simplePathJoin(ref_dir, path)
	}

	root, directory_parts, file_part, _ := simplePathSplit(full_path)

	directory := simplePathJoin(directory_parts...)
	location := simplePathJoin(root, directory)
	file_base := os_util.Basename(file_part)
	file_stem, file_ext := util_common.SplitBasename(file_base)

	pd := types.PathDescriptor{}

	pd.Classification = classification
	pd.Input = path
	pd.FullPath = full_path
	pd.Location = location
	pd.Root = root
	pd.Directory = directory
	pd.DirectoryParts = directory_parts
	pd.NumDotsDirectoryParts = countDotsDirectoryParts(pd.DirectoryParts)
	pd.EntryName = file_base
	pd.Stem = file_stem
	pd.Extension = file_ext

	return pd, nil
}

// Examines the path to the degree necessary to be able to classify it.
//
// # Parameters:
// - `path` - the given path to be classified;
// - `parseFlags` - flags that moderate the classification;
//
// # Returns:
// `(classification : Classification, root : string, pathRootStripped : string, firstBadCharIndex : int)`
func ClassifyRoot(path string, parseFlags types.ParseFlags) (types.Classification, string, string, int) {

	if len(path) == 0 {
		return types.Empty, "", "", 0
	}

	if os_util.ByteIsPathElementSeparator(path[0]) {
		return types.SlashRooted, path[:1], path[1:], -1
	}

	if len(path) >= 2 {
		if os_util.ByteIsValidDriveLetter(path[0]) {
			if os_util.ByteIsInvalidInPath(path[1]) {
				return types.InvalidChars, "", "", 1
			}
			if ':' == path[1] {
				if len(path) == 2 {
					return types.DriveLetterRelative, path, "", -1
				} else {
					if os_util.ByteIsInvalidInPath(path[2]) {
						return types.InvalidChars, "", "", 2
					}
					if os_util.ByteIsPathElementSeparator(path[2]) {
						return types.DriveLetterRooted, path[:3], path[3:], -1
					} else {
						return types.DriveLetterRelative, path[:2], path[2:], -1
					}
				}
			}
		}
	}

	if path[0] == '~' {
		if types.ParseFlags_RecogniseTildeHome == (parseFlags & types.ParseFlags_RecogniseTildeHome) {
			if len(path) == 1 {
				return types.HomeRooted, path, "", -1
			}
			if os_util.ByteIsInvalidInPath(path[1]) {
				return types.InvalidChars, "", "", 1
			}
			if os_util.ByteIsPathElementSeparator(path[1]) {
				return types.HomeRooted, path[:2], path[2:], -1
			}
		}
	}

	if os_util.ByteIsInvalidInPath(path[0]) {
		return types.InvalidChars, "", "", 0
	}

	return types.Relative, "", path, -1
}

func ParsePathStringFlags(path string, reference_directory string, parseFlags types.ParseFlags) (types.PathDescriptor, error) {
	return createPathDescriptor(path, reference_directory, parseFlags)
}
