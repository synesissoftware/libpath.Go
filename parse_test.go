package libpath_test

import (
	"github.com/synesissoftware/libpath.Go/internal/test_utils"
	"github.com/synesissoftware/libpath.Go/parse"
	"github.com/synesissoftware/libpath.Go/parse/common"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_ParsePathString_EMPTY(t *testing.T) {

	{
		pd, _ := parse.ParsePathString("", "")

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

func Test_ParsePathString_DOT(t *testing.T) {

	{
		{
			pd, _ := parse.ParsePathString(".", "")

			require.Equal(t, ".", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, ".", pd.EntryName)
			require.Equal(t, ".", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("..", "")

			require.Equal(t, "..", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "..", pd.EntryName)
			require.Equal(t, "..", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("...", "")

			require.Equal(t, "...", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "...", pd.EntryName)
			require.Equal(t, "...", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString(".....", "")

			require.Equal(t, ".....", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, ".....", pd.EntryName)
			require.Equal(t, ".....", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("..abc.def..", "")

			require.Equal(t, "..abc.def..", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "..abc.def..", pd.EntryName)
			require.Equal(t, "..abc.def..", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

	}

}

func Test_ParsePathString_STEM_ONLY(t *testing.T) {

	{
		pd, _ := parse.ParsePathString("abc", "")

		require.Equal(t, "abc", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "abc", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "/")

		require.Equal(t, "/abc", pd.FullPath)
		require.Equal(t, "/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "abc", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", ".")

		require.Equal(t, "./abc", pd.FullPath)
		require.Equal(t, "./", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 1, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"./"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "./")

		require.Equal(t, "./abc", pd.FullPath)
		require.Equal(t, "./", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 1, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"./"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "..")

		require.Equal(t, "../abc", pd.FullPath)
		require.Equal(t, "../", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "../", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 1, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"../"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "../")

		require.Equal(t, "../abc", pd.FullPath)
		require.Equal(t, "../", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "../", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 1, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"../"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "/dir-1/dir-2")

		require.Equal(t, "/dir-1/dir-2/abc", pd.FullPath)
		require.Equal(t, "/dir-1/dir-2/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "dir-1/dir-2")

		require.Equal(t, "dir-1/dir-2/abc", pd.FullPath)
		require.Equal(t, "dir-1/dir-2/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}

func Test_ParsePathString_BASENAME_ONLY(t *testing.T) {

	{
		pd, _ := parse.ParsePathString("abc.ex", "")

		require.Equal(t, "abc.ex", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "abc.ex", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc.ex", "/")

		require.Equal(t, "/abc.ex", pd.FullPath)
		require.Equal(t, "/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, "abc.ex", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc.ex", "./")

		require.Equal(t, "./abc.ex", pd.FullPath)
		require.Equal(t, "./", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 1, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"./"}, pd.DirectoryParts)
		require.Equal(t, "abc.ex", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc.ex", "/dir-1/dir-2")

		require.Equal(t, "/dir-1/dir-2/abc.ex", pd.FullPath)
		require.Equal(t, "/dir-1/dir-2/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, "abc.ex", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc.ex", "dir-1/dir-2")

		require.Equal(t, "dir-1/dir-2/abc.ex", pd.FullPath)
		require.Equal(t, "dir-1/dir-2/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, "abc.ex", pd.EntryName)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}

func Test_ParsePathString_EXTENSION_ONLY(t *testing.T) {

	{
		pd, _ := parse.ParsePathString(".ex", "")

		require.Equal(t, ".ex", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, ".ex", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString(".ex", "/")

		require.Equal(t, "/.ex", pd.FullPath)
		require.Equal(t, "/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, ".ex", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString(".ex", "./")

		require.Equal(t, "./.ex", pd.FullPath)
		require.Equal(t, "./", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 1, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"./"}, pd.DirectoryParts)
		require.Equal(t, ".ex", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString(".ex", "/dir-1/dir-2")

		require.Equal(t, "/dir-1/dir-2/.ex", pd.FullPath)
		require.Equal(t, "/dir-1/dir-2/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, ".ex", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString(".ex", "dir-1/dir-2")

		require.Equal(t, "dir-1/dir-2/.ex", pd.FullPath)
		require.Equal(t, "dir-1/dir-2/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, ".ex", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}

func Test_ParsePathString_DIRECTORY_ONLY(t *testing.T) {

	{
		pd, _ := parse.ParsePathString("abc/", "")

		require.Equal(t, "abc/", pd.FullPath)
		require.Equal(t, "abc/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "abc/", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"abc/"}, pd.DirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc/", "/")

		require.Equal(t, "/abc/", pd.FullPath)
		require.Equal(t, "/abc/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "abc/", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, 0, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"abc/"}, pd.DirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc/", "./")

		require.Equal(t, "./abc/", pd.FullPath)
		require.Equal(t, "./abc/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./abc/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, 1, pd.NumDotsDirectoryParts)
		require.Equal(t, []string{"./", "abc/"}, pd.DirectoryParts)
		require.Equal(t, "", pd.EntryName)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		test_utils.CheckPathDescriptorElements(t, pd)
	}
}

func Test_ParsePathString_AbsolutePath_ensuring_ignoring_ReferenceDirectory(t *testing.T) {

	/* path = "/" */
	{
		{
			pd, _ := parse.ParsePathString("/", "")

			require.Equal(t, "/", pd.FullPath)
			require.Equal(t, "/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/", "abc")

			require.Equal(t, "/", pd.FullPath)
			require.Equal(t, "/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/", "/dir-1/dir-2")

			require.Equal(t, "/", pd.FullPath)
			require.Equal(t, "/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}
	}

	/* path = "/dir-1/dir-2/" */
	{
		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/", "")

			require.Equal(t, "/dir-1/dir-2/", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/", "abc")

			require.Equal(t, "/dir-1/dir-2/", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/", "/dir-1/dir-2")

			require.Equal(t, "/dir-1/dir-2/", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}
	}

	/* path = "/dir-1/dir-2/file.ext" */
	{
		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/file.ext", "")

			require.Equal(t, "/dir-1/dir-2/file.ext", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "file.ext", pd.EntryName)
			require.Equal(t, "file", pd.Stem)
			require.Equal(t, ".ext", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/file.ext", "abc")

			require.Equal(t, "/dir-1/dir-2/file.ext", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "file.ext", pd.EntryName)
			require.Equal(t, "file", pd.Stem)
			require.Equal(t, ".ext", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/file.ext", "/dir-1/dir-2")

			require.Equal(t, "/dir-1/dir-2/file.ext", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "file.ext", pd.EntryName)
			require.Equal(t, "file", pd.Stem)
			require.Equal(t, ".ext", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}
	}
}

func Test_ParsePathString_HOME_ROOTED(t *testing.T) {

	/* path = "~" (ParseFlags_RecogniseTildeHome not specified) */
	{
		{
			pd, _ := parse.ParsePathString("~", "")

			require.Equal(t, "~", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "~", pd.EntryName)
			require.Equal(t, "~", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("~", "abc")

			require.Equal(t, "abc/~", pd.FullPath)
			require.Equal(t, "abc/", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "abc/", pd.Directory)
			require.Equal(t, 1, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "~", pd.EntryName)
			require.Equal(t, "~", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("~", "/dir-1/dir-2")

			require.Equal(t, "/dir-1/dir-2/~", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "~", pd.EntryName)
			require.Equal(t, "~", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}
	}

	/* path = "~" (ParseFlags_RecogniseTildeHome specified) */
	{
		{
			pd, _ := parse.ParsePathStringFlags("~", "", common.ParseFlags_RecogniseTildeHome)

			require.Equal(t, "~", pd.FullPath)
			require.Equal(t, "~", pd.Location)
			require.Equal(t, "~", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathStringFlags("~", "abc", common.ParseFlags_RecogniseTildeHome)

			require.Equal(t, "~", pd.FullPath)
			require.Equal(t, "~", pd.Location)
			require.Equal(t, "~", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathStringFlags("~", "/dir-1/dir-2", common.ParseFlags_RecogniseTildeHome)

			require.Equal(t, "~", pd.FullPath)
			require.Equal(t, "~", pd.Location)
			require.Equal(t, "~", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, 0, pd.NumDotsDirectoryParts)
			require.Equal(t, "", pd.EntryName)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			test_utils.CheckPathDescriptorElements(t, pd)
		}
	}
}
