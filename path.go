package libpath

import (
	angols_strings "github.com/synesissoftware/ANGoLS/strings"

	"strings"
)

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
	input          string   // The original input string to the parsing.
	FullPath       string   // The fullest esablishable form of the input.
	Location       string   // The fullest esablishable form of the location of the entry, which is everything up-to-and-including the last (if any) path-name separator.
	Root           string   // The path root, if present.
	Directory      string   // The path directory, which excluses the `Root` (if any) and the `Basename` (if any).
	DirectoryParts []string // Array of the `directory` elements, split on the path-name separator.
	Basename       string   // The "file part", if any, which occurs after the last (if any) path-name separator.
	Stem           string   // The stem of the "file part", if any.
	Extension      string   // The extension of the "file part", if any.
}

// The number of dots directories in the descriptor
func (pd PathDescriptor) NumberOfDotsDirectoryParts() int {
	r := 0

	for _, dp := range pd.DirectoryParts {
		switch countDotsDirectoryPart(dp) {
		case 1:
			fallthrough
		case 2:
			r += 1
		default:
		}
	}

	return r
}

func countDotsDirectoryPart(s string) int {

	ix_no := angols_strings.IndexNotAnyAfter(s, "./", -1)

	if -1 != ix_no {
		return 0
	}

	for len(s) > 1 {
		if elementEndsWithPathNameSeparator(s) {
			s = s[:len(s)-1]
		} else {
			break
		}
	}

	return len(s)
}

func elementIsRoot(s string) bool {
	if 0 == len(s) {

		return false
	} else {

		return '/' == s[0]
	}
}

func byteIsPathElementSeparator(c byte) bool {
	return '/' == c
}

func charIsPathElementSeparator(c rune) bool {
	return '/' == c
}

func elementEndsWithPathNameSeparator(s string) bool {
	switch len(s) {
	case 0:

		return false
	default:

		return byteIsPathElementSeparator(s[len(s)-1])
	}
}

func pathIsAbsolute(path string) bool {
	return unixPathIsAbsolute(path)
}

func unixPathIsAbsolute(path string) bool {
	if 0 == len(path) {

		return false
	} else {

		return '/' == path[0]
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
	var basename string

	if elementEndsWithPathNameSeparator(last) {
		basename = ""
	} else {
		basename = last
		dp_count -= 1
	}

	if elementIsRoot(first) {
		root = first
		dp_from += 1
	} else {
		root = ""
	}

	directory_parts = splits[dp_from:dp_count]

	return root, directory_parts, basename, nil
}

/* NOTE: if base name ends with '.', then no extension is obtained;
 * otherwise has the normal semantics
 */
func splitBasename(base_name string) (stem, extension string) {

	lix_dot := strings.LastIndexByte(base_name, '.')

	if lix_dot < 0 || lix_dot == len(base_name)-1 {

		stem = base_name
	} else {

		stem = base_name[:lix_dot]
		extension = base_name[lix_dot:]
	}

	return
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

	if !pathIsAbsolute(path) && 0 != len(ref_dir) {
		full_path = simplePathJoin(ref_dir, path)
	}

	root, directory_parts, basename, _ := simplePathSplit(full_path)

	directory := simplePathJoin(directory_parts...)
	location := simplePathJoin(root, directory)
	stem, extension := splitBasename(basename)

	pd.FullPath = full_path
	pd.Location = location
	pd.Root = root
	pd.Directory = directory
	pd.DirectoryParts = directory_parts
	pd.Basename = basename
	pd.Stem = stem
	pd.Extension = extension

	return pd, nil
}

func ParsePathString(path string, reference_directory string /*, ... interface{}*/) (PathDescriptor, error) {

	return createPathDescriptor(path, reference_directory)
}
