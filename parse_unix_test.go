//go:build unix

package libpath_test

import (
	. "github.com/synesissoftware/libpath.Go/internal/test_utils"

	"github.com/synesissoftware/libpath.Go/parse"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_UNIX_EMPTY_PATH(t *testing.T) {
	{
		pd, _ := parse.ParsePathString("", "")

		require.Equal(t, "", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}
}
