//go:build unix

package libpath_test

import (
	"github.com/synesissoftware/libpath.Go/internal/test_utils"
	"github.com/synesissoftware/libpath.Go/parse/common"
	os_parse "github.com/synesissoftware/libpath.Go/parse/unix"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_CheckPathDescriptorElements_UNIX_EMPTY_PATH(t *testing.T) {
	{
		pd, _ := os_parse.ParsePathStringFlags("", "", common.ParseFlags_None)

		require.Equal(t, "", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}
