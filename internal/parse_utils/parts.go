package parse_utils

// T.B.C.
//
// Precondition:
// - none of the parts is empty, and none begins with '/' (or '\\' on Windows);
func CountDotsDirectoryParts(directoryParts []string, isDotsPart func(dp string) bool) int {
	r := 0

	for _, dp := range directoryParts {
		if isDotsPart(dp) {
			r += 1
		}
	}

	return r
}

func ElementIsRooted(s string, f func(firstByte byte) bool) bool {
	switch len(s) {
	case 0:

		return false
	default:

		return f(s[0])
	}
}

func ElementEndsWithByte(s string, f func(firstByte byte) bool) bool {
	switch len(s) {
	case 0:

		return false
	default:

		return f(s[len(s)-1])
	}
}
