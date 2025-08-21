//go:build windows

package libpath_test

import (
	"github.com/synesissoftware/libpath.Go/internal/test_utils"
	"github.com/synesissoftware/libpath.Go/parse/common"
	os_parse "github.com/synesissoftware/libpath.Go/parse/windows"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_ParsePathString_WINDOWS_EMPTY_PATH(t *testing.T) {
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

func Test_ParsePathString_WINDOWS_VOLUME_ROOTED_COMPLETE_PATHS(t *testing.T) {
	{
		pd, _ := os_parse.ParsePathStringFlags("C:\\dir\\file.ext", "", common.ParseFlags_None)

		require.Equal(t, "C:\\dir\\file.ext", pd.FullPath)
		require.Equal(t, "C:\\dir\\", pd.Location)
		require.Equal(t, "C:\\", pd.Root)
		require.Equal(t, "dir\\", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir\\"}, pd.DirectoryParts)
		require.Equal(t, "file.ext", pd.EntryName)
		require.Equal(t, "file", pd.Stem)
		require.Equal(t, ".ext", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := os_parse.ParsePathStringFlags("C:/dir/file.ext", "", common.ParseFlags_None)

		require.Equal(t, "C:/dir/file.ext", pd.FullPath)
		require.Equal(t, "C:/dir/", pd.Location)
		require.Equal(t, "C:/", pd.Root)
		require.Equal(t, "dir/", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir/"}, pd.DirectoryParts)
		require.Equal(t, "file.ext", pd.EntryName)
		require.Equal(t, "file", pd.Stem)
		require.Equal(t, ".ext", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}

func Test_ParsePathString_WINDOWS_VOLUME_ROOTED_DIRECTORY(t *testing.T) {
	{
		pd, _ := os_parse.ParsePathStringFlags("C:\\dir\\", "", common.ParseFlags_None)

		require.Equal(t, "C:\\dir\\", pd.FullPath)
		require.Equal(t, "C:\\dir\\", pd.Location)
		require.Equal(t, "C:\\", pd.Root)
		require.Equal(t, "dir\\", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir\\"}, pd.DirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := os_parse.ParsePathStringFlags("C:/dir/", "", common.ParseFlags_None)

		require.Equal(t, "C:/dir/", pd.FullPath)
		require.Equal(t, "C:/dir/", pd.Location)
		require.Equal(t, "C:/", pd.Root)
		require.Equal(t, "dir/", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir/"}, pd.DirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}

func Test_ParsePathString_WINDOWS_VOLUME_ROOTED_FILES(t *testing.T) {
	{
		pd, _ := os_parse.ParsePathStringFlags("C:\\file.ext", "", common.ParseFlags_None)

		require.Equal(t, "C:\\file.ext", pd.FullPath)
		require.Equal(t, "C:\\", pd.Location)
		require.Equal(t, "C:\\", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "file.ext", pd.EntryName)
		require.Equal(t, "file", pd.Stem)
		require.Equal(t, ".ext", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := os_parse.ParsePathStringFlags("C:/file.ext", "", common.ParseFlags_None)

		require.Equal(t, "C:/file.ext", pd.FullPath)
		require.Equal(t, "C:/", pd.Location)
		require.Equal(t, "C:/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "file.ext", pd.EntryName)
		require.Equal(t, "file", pd.Stem)
		require.Equal(t, ".ext", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}

func Test_ParsePathString_WINDOWS_VOLUME_ONLY(t *testing.T) {
	{
		pd, _ := os_parse.ParsePathStringFlags("C:\\", "", common.ParseFlags_None)

		require.Equal(t, "C:\\", pd.FullPath)
		require.Equal(t, "C:\\", pd.Location)
		require.Equal(t, "C:\\", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := os_parse.ParsePathStringFlags("C:/", "", common.ParseFlags_None)

		require.Equal(t, "C:/", pd.FullPath)
		require.Equal(t, "C:/", pd.Location)
		require.Equal(t, "C:/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}
