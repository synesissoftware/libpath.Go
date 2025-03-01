//go:build windows

package util

import (
	"github.com/synesissoftware/libpath.Go/util/windows"
)

func Basename(path string) string {
	return windows.Basename(path)
}

func ByteIsPathElementSeparator(c byte) bool {
	return windows.ByteIsPathElementSeparator(c)
}

func CharIsPathElementSeparator(c rune) bool {
	return windows.CharIsPathElementSeparator(c)
}

func PathIsAbsolute(path string) bool {
	return windows.PathIsAbsolute(path)
}
