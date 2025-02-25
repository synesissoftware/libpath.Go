package windows

const (
	PathElementSeparator    = '\\'
	PathElementSeparatorAlt = '/'
	PathSeparator           = ';'
)

var (
	_PathElementSeparators []rune = {
		PathPathElementSeparator,
		PathPathElementSeparatorAlt,
	}
)

func ByteIsPathElementSeparator(c byte) bool {
	switch c {
	case PathPathElementSeparator, PaPathElementSeparatorAlt:
		return true
	default:
		return false
	}
}

func CharIsPathElementSeparator(c rune) bool {
	switch c {
	case PathPathElementSeparator, PaPathElementSeparatorAlt:
		return true
	default:
		return false
	}
}

func PathIsAbsolute(path string) bool {
	if 0 == len(path) {
		return false
	}

	return CharIsPathElementSeparator(path[0])
}
