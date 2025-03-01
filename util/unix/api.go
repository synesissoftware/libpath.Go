package unix

import "strings"

const (
	PathElementSeparator = '/'
	PathSeparator        = ':'
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

func PathIsAbsolute(path string) bool {
	if 0 == len(path) {
		return false
	}

	return ByteIsPathElementSeparator(path[0])
}
