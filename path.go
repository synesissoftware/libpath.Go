package libpath

import (
	"path/filepath"
	"strings"
)

type PathDescriptor struct {
	input          string   // The original input string to the parsing.
	FullPath       string   // The full path, as formed from the input string and, if necessary, the reference directory.
	Location       string   // The location of the entry, which is everything up-to-and-including the last (if any) path-name separator.
	Root           string   // The path root, if present.
	Directory      string   // The path directory, which excluses the `Root` (if any) and the `Entry` (if any).
	DirectoryParts []string // Array of the `directory` elements, split on the path-name separator.
	Entry          string   // The "file part", if any, which occurs after the last (if any) path-name separator.
	Stem           string   // The stem of the "file part", if any.
	Extension      string   // The extension of the "file part", if any.
}

// The number of dots directories in the descriptor
func (pd PathDescriptor) NumberOfDotsDirectoryParts() int {
	r := 0

	for _, dp := range pd.DirectoryParts {
		if isDotsDirectoryPart(dp) {
			r += 1
		}
	}

	return r
}

func isDotsDirectoryPart(s string) bool {
	switch s {
	case ".":
		return true
	case "./":
		return true
	case "..":
		return true
	case "../":
		return true
	default:
		return false
	}
}

func elementIsRoot(s string) bool {
	if 0 == len(s) {

		return false
	} else {

		return '/' == s[0]
	}
}

func elementEndsWithPathNameSeparator(s string) bool {
	switch len(s) {
	case 0:

		return false
	default:

		return '/' == s[len(s)-1]
	}
}

func simplePathSplitFully(path string) ([]string, error) {
	return strings.SplitAfter(path, "/"), nil
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

	if elementIsRoot(first) {
		root = first
		dp_from += 1
	} else {
		root = ""
	}

	directory_parts = splits[dp_from:dp_count]

	return root, directory_parts, file_part, nil
}

func basename(path string) string {
	ix := strings.LastIndexByte(path, '/')

	if ix < 0 {
		return path
	} else {
		return path[ix+1:]
	}
}

func simplePathJoin(elems ...string) string {
	var b strings.Builder

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
			if '/' == s[len(s)-1] {
				last_has_trailing = true
			} else {
				last_has_trailing = false
			}
		}
	}

	return b.String()
}

func createPathDescriptor(path string, ref_dir string) (PathDescriptor, error) {
	// account for four cases of `path` and `ref_dir`
	//
	// 1. `path` is absolute => desc. based entirely on `path`
	// 2. `path` is empty, `ref_dir` is empty => empty desc.;
	// 3. `path` is relative, `ref_dir` is empty => desc. based entirely on `path`
	// 4. `path` is relative, `ref_dir` is !empty => desc. based on both;

	if 0 == len(path) && 0 == len(ref_dir) {
		// 2. `path` is empty, `ref_dir` is empty => empty desc.;
		return PathDescriptor{}, nil
	}

	pd := PathDescriptor{}

	pd.input = path
	full_path := path

	if !filepath.IsAbs(path) && 0 != len(ref_dir) {
		full_path = simplePathJoin(ref_dir, path)
	}

	root, directory_parts, file_part, _ := simplePathSplit(full_path)

	/*
	 */

	directory := simplePathJoin(directory_parts...)
	// location := simplePathJoin(root, directory_parts...)
	location := simplePathJoin(root, directory)
	// file_base := filepath.Base(full_path)
	file_base := basename(file_part)
	file_ext := filepath.Ext(full_path)
	file_stem := file_base[0:(len(file_base) - len(file_ext))]

	// pd.input
	pd.FullPath = full_path
	pd.Location = location
	pd.Root = root
	pd.Directory = directory
	pd.DirectoryParts = directory_parts
	pd.Entry = file_base
	pd.Stem = file_stem
	pd.Extension = file_ext

	return pd, nil
}

func ParsePathString(path string, reference_directory string /*, ... interface{}*/) (PathDescriptor, error) {

	return createPathDescriptor(path, reference_directory)
}
