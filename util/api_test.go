package util_test

import (
	"github.com/synesissoftware/libpath.Go/util"

	"github.com/stretchr/testify/require"

	"testing"
)

func checkSplitBasename(t *testing.T, expected_stem, expected_extension, basename_to_split string) {
	t.Helper()

	actual_stem, actual_extension := util.SplitBasename(basename_to_split)

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
