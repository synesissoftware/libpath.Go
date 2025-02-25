package windows

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func Test_Basename(t *testing.T) {
	{
		require.Equal(t, ".", Basename("."))
	}

	{
		require.Equal(t, "abc.ext", Basename("abc.ext"))
		require.Equal(t, "abc", Basename("abc"))
		require.Equal(t, ".ext", Basename(".ext"))
	}

	{
		require.Equal(t, "abc.ext", Basename("\\abc.ext"))
		require.Equal(t, "abc", Basename("\\abc"))
		require.Equal(t, ".ext", Basename("\\.ext"))

		require.Equal(t, "abc.ext", Basename("/abc.ext"))
		require.Equal(t, "abc", Basename("/abc"))
		require.Equal(t, ".ext", Basename("/.ext"))
	}

	{
		require.Equal(t, "abc.ext", Basename("dir\\abc.ext"))
		require.Equal(t, "abc", Basename("dir\\abc"))
		require.Equal(t, ".ext", Basename("dir\\.ext"))

		require.Equal(t, "abc.ext", Basename("dir/abc.ext"))
		require.Equal(t, "abc", Basename("dir/abc"))
		require.Equal(t, ".ext", Basename("dir/.ext"))
	}

	{
		require.Equal(t, "abc.ext", Basename("\\dir\\abc.ext"))
		require.Equal(t, "abc.ext", Basename("\\dir/abc.ext"))
		require.Equal(t, "abc.ext", Basename("/dir\\abc.ext"))
		require.Equal(t, "abc.ext", Basename("/dir/abc.ext"))
	}
}
