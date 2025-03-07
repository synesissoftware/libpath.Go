package util

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func checkSplitBasename(t *testing.T, expected_stem string, expected_extension string, basename_to_split string) {
	t.Helper()

	actual_stem, actual_extension := SplitBasename(basename_to_split)

	require.Equal(t, expected_stem, actual_stem)
	require.Equal(t, expected_extension, actual_extension)
}

func Test_SplitBasename(t *testing.T) {
	{
		checkSplitBasename(t, "abc", ".ext", "abc.ext")
	}

	{
		checkSplitBasename(t, "abc", "", "abc")
	}

	{
		checkSplitBasename(t, "", ".ext", ".ext")
	}

	{
		checkSplitBasename(t, ".", "", ".")
	}

	{
		checkSplitBasename(t, ".ext.", "", ".ext.")
	}
}
