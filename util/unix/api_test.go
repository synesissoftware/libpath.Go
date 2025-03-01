package unix_test

import (
	unix "github.com/synesissoftware/libpath.Go/util/unix"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_Basename(t *testing.T) {
	{
		require.Equal(t, ".", unix.Basename("."))
	}

	{
		require.Equal(t, "abc.ext", unix.Basename("abc.ext"))
		require.Equal(t, "abc", unix.Basename("abc"))
		require.Equal(t, ".ext", unix.Basename(".ext"))
	}

	{
		require.Equal(t, "abc.ext", unix.Basename("/abc.ext"))
		require.Equal(t, "abc", unix.Basename("/abc"))
		require.Equal(t, ".ext", unix.Basename("/.ext"))
	}

	{
		require.Equal(t, "abc.ext", unix.Basename("dir/abc.ext"))
		require.Equal(t, "abc", unix.Basename("dir/abc"))
		require.Equal(t, ".ext", unix.Basename("dir/.ext"))
	}

	{
		require.Equal(t, "abc.ext", unix.Basename("/dir/abc.ext"))
	}
}
