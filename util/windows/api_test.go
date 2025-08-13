package windows_test

import (
	"github.com/synesissoftware/libpath.Go/util/windows"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_Basename(t *testing.T) {
	{
		require.Equal(t, ".", windows.Basename("."))
	}

	{
		require.Equal(t, "abc.ext", windows.Basename("abc.ext"))
		require.Equal(t, "abc", windows.Basename("abc"))
		require.Equal(t, ".ext", windows.Basename(".ext"))
	}

	{
		require.Equal(t, "abc.ext", windows.Basename("\\abc.ext"))
		require.Equal(t, "abc", windows.Basename("\\abc"))
		require.Equal(t, ".ext", windows.Basename("\\.ext"))

		require.Equal(t, "abc.ext", windows.Basename("/abc.ext"))
		require.Equal(t, "abc", windows.Basename("/abc"))
		require.Equal(t, ".ext", windows.Basename("/.ext"))
	}

	{
		require.Equal(t, "abc.ext", windows.Basename("dir\\abc.ext"))
		require.Equal(t, "abc", windows.Basename("dir\\abc"))
		require.Equal(t, ".ext", windows.Basename("dir\\.ext"))

		require.Equal(t, "abc.ext", windows.Basename("dir/abc.ext"))
		require.Equal(t, "abc", windows.Basename("dir/abc"))
		require.Equal(t, ".ext", windows.Basename("dir/.ext"))
	}

	{
		require.Equal(t, "abc.ext", windows.Basename("\\dir\\abc.ext"))
		require.Equal(t, "abc.ext", windows.Basename("\\dir/abc.ext"))
		require.Equal(t, "abc.ext", windows.Basename("/dir\\abc.ext"))
		require.Equal(t, "abc.ext", windows.Basename("/dir/abc.ext"))
	}
}
