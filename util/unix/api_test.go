package unix_test

import (
	api "github.com/synesissoftware/libpath.Go/util/unix"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_Basename(t *testing.T) {
	{
		require.Equal(t, ".", api.Basename("."))
	}

	{
		require.Equal(t, "abc.ext", api.Basename("abc.ext"))
		require.Equal(t, "abc", api.Basename("abc"))
		require.Equal(t, ".ext", api.Basename(".ext"))
	}

	{
		require.Equal(t, "abc.ext", api.Basename("/abc.ext"))
		require.Equal(t, "abc", api.Basename("/abc"))
		require.Equal(t, ".ext", api.Basename("/.ext"))
	}

	{
		require.Equal(t, "abc.ext", api.Basename("dir/abc.ext"))
		require.Equal(t, "abc", api.Basename("dir/abc"))
		require.Equal(t, ".ext", api.Basename("dir/.ext"))
	}

	{
		require.Equal(t, "abc.ext", api.Basename("/dir/abc.ext"))
	}
}
