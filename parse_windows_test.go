//go:build windows

package libpath_test

import (
	. "github.com/synesissoftware/libpath.Go/internal/test_utils"

	"github.com/synesissoftware/libpath.Go/parse"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_Windows_EMPTY_PATH(t *testing.T) {
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

func Test_Windows_VOLUME_ROOTED_COMPLETE_PATHS(t *testing.T) {
	{
		pd, _ := parse.ParsePathString("C:\\dir\\file.ext", "")

		require.Equal(t, "C:\\dir\\file.ext", pd.FullPath)
		require.Equal(t, "C:\\dir\\", pd.Location)
		require.Equal(t, "C:\\", pd.Root)
		require.Equal(t, "dir\\", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir\\"}, pd.DirectoryParts)
		require.Equal(t, "file.ext", pd.EntryName)
		require.Equal(t, "file", pd.Stem)
		require.Equal(t, ".ext", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("C:/dir/file.ext", "")

		require.Equal(t, "C:/dir/file.ext", pd.FullPath)
		require.Equal(t, "C:/dir/", pd.Location)
		require.Equal(t, "C:/", pd.Root)
		require.Equal(t, "dir/", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir/"}, pd.DirectoryParts)
		require.Equal(t, "file.ext", pd.EntryName)
		require.Equal(t, "file", pd.Stem)
		require.Equal(t, ".ext", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}
}

func Test_Windows_VOLUME_ROOTED_DIRECTORY(t *testing.T) {
	{
		pd, _ := parse.ParsePathString("C:\\dir\\", "")

		require.Equal(t, "C:\\dir\\", pd.FullPath)
		require.Equal(t, "C:\\dir\\", pd.Location)
		require.Equal(t, "C:\\", pd.Root)
		require.Equal(t, "dir\\", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir\\"}, pd.DirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("C:/dir/", "")

		require.Equal(t, "C:/dir/", pd.FullPath)
		require.Equal(t, "C:/dir/", pd.Location)
		require.Equal(t, "C:/", pd.Root)
		require.Equal(t, "dir/", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir/"}, pd.DirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}
}

func Test_Windows_VOLUME_ROOTED_FILES(t *testing.T) {
	{
		pd, _ := parse.ParsePathString("C:\\file.ext", "")

		require.Equal(t, "C:\\file.ext", pd.FullPath)
		require.Equal(t, "C:\\", pd.Location)
		require.Equal(t, "C:\\", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "file.ext", pd.EntryName)
		require.Equal(t, "file", pd.Stem)
		require.Equal(t, ".ext", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("C:/file.ext", "")

		require.Equal(t, "C:/file.ext", pd.FullPath)
		require.Equal(t, "C:/", pd.Location)
		require.Equal(t, "C:/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "file.ext", pd.EntryName)
		require.Equal(t, "file", pd.Stem)
		require.Equal(t, ".ext", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}
}

func Test_Windows_VOLUME_ONLY(t *testing.T) {
	{
		pd, _ := parse.ParsePathString("C:\\", "")

		require.Equal(t, "C:\\", pd.FullPath)
		require.Equal(t, "C:\\", pd.Location)
		require.Equal(t, "C:\\", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("C:/", "")

		require.Equal(t, "C:/", pd.FullPath)
		require.Equal(t, "C:/", pd.Location)
		require.Equal(t, "C:/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		CheckPathDescriptorElements(t, pd)
	}
}
