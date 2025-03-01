//go:build unix

package libpath_test

import (
	"strings"

	"github.com/stretchr/testify/require"
	"github.com/synesissoftware/libpath.Go/parse"

	"testing"
)

func checkPathDescriptorElements(t *testing.T, pd parse.PathDescriptor) {
	t.Helper()

	// Stem + Extension => Entry
	{
		require.Equal(t, pd.Entry, pd.Stem+pd.Extension)
	}

	// Location + Entry => FullPath
	{
		require.Equal(t, pd.FullPath, pd.Location+pd.Entry)
	}

	// Root + Directory + Stem + Extension => FullPath
	{
		require.Equal(t, pd.FullPath, pd.Root+pd.Directory+pd.Stem+pd.Extension)
	}

	// Root + DirectoryParts => Location
	{
		require.Equal(t, pd.Location, pd.Root+strings.Join(pd.DirectoryParts, ""))
	}
}

func Test_UNIX_EMPTY_PATH(t *testing.T) {
	{
		pd, _ := parse.ParsePathString("", "")

		require.Equal(t, "", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}
}
