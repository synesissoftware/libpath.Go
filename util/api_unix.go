package util

import (
	"github.com/synesissoftware/libpath.Go/util/unix"
)

func Basename(path string) string {
	return unix.Basename(path)
}

func ByteIsPathElementSeparator(c byte) bool {
	return unix.ByteIsPathElementSeparator(c)
}

func CharIsPathElementSeparator(c rune) bool {
	return unix.CharIsPathElementSeparator(c)
}

func PathIsAbsolute(path string) bool {
	return unix.PathIsAbsolute(path)
}
