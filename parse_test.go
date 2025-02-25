package libpath

import (
	"github.com/stretchr/testify/require"
	"github.com/synesissoftware/libpath.Go/parse"

	"strings"
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

func Test_empty(t *testing.T) {

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

func Test_dot(t *testing.T) {

	{
		{
			pd, _ := parse.ParsePathString(".", "")

			require.Equal(t, ".", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, ".", pd.Entry)
			require.Equal(t, ".", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("..", "")

			require.Equal(t, "..", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, "..", pd.Entry)
			require.Equal(t, "..", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("...", "")

			require.Equal(t, "...", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, "...", pd.Entry)
			require.Equal(t, "...", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString(".....", "")

			require.Equal(t, ".....", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, ".....", pd.Entry)
			require.Equal(t, ".....", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("..abc.def..", "")

			require.Equal(t, "..abc.def..", pd.FullPath)
			require.Equal(t, "", pd.Location)
			require.Equal(t, "", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, "..abc.def..", pd.Entry)
			require.Equal(t, "..abc.def..", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

	}

}

func Test_Parse_Stem_only(t *testing.T) {

	{
		pd, _ := parse.ParsePathString("abc", "")

		require.Equal(t, "abc", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "abc", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "/")

		require.Equal(t, "/abc", pd.FullPath)
		require.Equal(t, "/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "abc", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", ".")

		require.Equal(t, "./abc", pd.FullPath)
		require.Equal(t, "./", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"./"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 1, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "./")

		require.Equal(t, "./abc", pd.FullPath)
		require.Equal(t, "./", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"./"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 1, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "..")

		require.Equal(t, "../abc", pd.FullPath)
		require.Equal(t, "../", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "../", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"../"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 1, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "../")

		require.Equal(t, "../abc", pd.FullPath)
		require.Equal(t, "../", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "../", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"../"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 1, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "/dir-1/dir-2")

		require.Equal(t, "/dir-1/dir-2/abc", pd.FullPath)
		require.Equal(t, "/dir-1/dir-2/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc", "dir-1/dir-2")

		require.Equal(t, "dir-1/dir-2/abc", pd.FullPath)
		require.Equal(t, "dir-1/dir-2/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, "abc", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}
}

func Test_Parse_Basename_only(t *testing.T) {

	{
		pd, _ := parse.ParsePathString("abc.ex", "")

		require.Equal(t, "abc.ex", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "abc.ex", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc.ex", "/")

		require.Equal(t, "/abc.ex", pd.FullPath)
		require.Equal(t, "/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, "abc.ex", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc.ex", "./")

		require.Equal(t, "./abc.ex", pd.FullPath)
		require.Equal(t, "./", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"./"}, pd.DirectoryParts)
		require.Equal(t, "abc.ex", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 1, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc.ex", "/dir-1/dir-2")

		require.Equal(t, "/dir-1/dir-2/abc.ex", pd.FullPath)
		require.Equal(t, "/dir-1/dir-2/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, "abc.ex", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc.ex", "dir-1/dir-2")

		require.Equal(t, "dir-1/dir-2/abc.ex", pd.FullPath)
		require.Equal(t, "dir-1/dir-2/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, "abc.ex", pd.Entry)
		require.Equal(t, "abc", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}
}

func Test_Parse_Extension_only(t *testing.T) {

	{
		pd, _ := parse.ParsePathString(".ex", "")

		require.Equal(t, ".ex", pd.FullPath)
		require.Equal(t, "", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, ".ex", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString(".ex", "/")

		require.Equal(t, "/.ex", pd.FullPath)
		require.Equal(t, "/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "", pd.Directory)
		require.Equal(t, 0, len(pd.DirectoryParts))
		require.Equal(t, ".ex", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString(".ex", "./")

		require.Equal(t, "./.ex", pd.FullPath)
		require.Equal(t, "./", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"./"}, pd.DirectoryParts)
		require.Equal(t, ".ex", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 1, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString(".ex", "/dir-1/dir-2")

		require.Equal(t, "/dir-1/dir-2/.ex", pd.FullPath)
		require.Equal(t, "/dir-1/dir-2/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, ".ex", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString(".ex", "dir-1/dir-2")

		require.Equal(t, "dir-1/dir-2/.ex", pd.FullPath)
		require.Equal(t, "dir-1/dir-2/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "dir-1/dir-2/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
		require.Equal(t, ".ex", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, ".ex", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}
}

func Test_Parse_Directory_only(t *testing.T) {

	{
		pd, _ := parse.ParsePathString("abc/", "")

		require.Equal(t, "abc/", pd.FullPath)
		require.Equal(t, "abc/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "abc/", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"abc/"}, pd.DirectoryParts)
		require.Equal(t, "", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc/", "/")

		require.Equal(t, "/abc/", pd.FullPath)
		require.Equal(t, "/abc/", pd.Location)
		require.Equal(t, "/", pd.Root)
		require.Equal(t, "abc/", pd.Directory)
		require.Equal(t, 1, len(pd.DirectoryParts))
		require.Equal(t, []string{"abc/"}, pd.DirectoryParts)
		require.Equal(t, "", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}

	{
		pd, _ := parse.ParsePathString("abc/", "./")

		require.Equal(t, "./abc/", pd.FullPath)
		require.Equal(t, "./abc/", pd.Location)
		require.Equal(t, "", pd.Root)
		require.Equal(t, "./abc/", pd.Directory)
		require.Equal(t, 2, len(pd.DirectoryParts))
		require.Equal(t, []string{"./", "abc/"}, pd.DirectoryParts)
		require.Equal(t, "", pd.Entry)
		require.Equal(t, "", pd.Stem)
		require.Equal(t, "", pd.Extension)

		require.Equal(t, 1, pd.NumberOfDotsDirectoryParts())

		checkPathDescriptorElements(t, pd)
	}
}

func Test_Parse_AbsolutePath_ensuring_ignoring_ReferenceDirectory(t *testing.T) {

	/* path = "/" */
	{
		{
			pd, _ := parse.ParsePathString("/", "")

			require.Equal(t, "/", pd.FullPath)
			require.Equal(t, "/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, "", pd.Entry)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/", "abc")

			require.Equal(t, "/", pd.FullPath)
			require.Equal(t, "/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, "", pd.Entry)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/", "/dir-1/dir-2")

			require.Equal(t, "/", pd.FullPath)
			require.Equal(t, "/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "", pd.Directory)
			require.Equal(t, 0, len(pd.DirectoryParts))
			require.Equal(t, "", pd.Entry)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
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
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "", pd.Entry)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/", "abc")

			require.Equal(t, "/dir-1/dir-2/", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "", pd.Entry)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/", "/dir-1/dir-2")

			require.Equal(t, "/dir-1/dir-2/", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "", pd.Entry)
			require.Equal(t, "", pd.Stem)
			require.Equal(t, "", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
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
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "file.ext", pd.Entry)
			require.Equal(t, "file", pd.Stem)
			require.Equal(t, ".ext", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/file.ext", "abc")

			require.Equal(t, "/dir-1/dir-2/file.ext", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "file.ext", pd.Entry)
			require.Equal(t, "file", pd.Stem)
			require.Equal(t, ".ext", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}

		{
			pd, _ := parse.ParsePathString("/dir-1/dir-2/file.ext", "/dir-1/dir-2")

			require.Equal(t, "/dir-1/dir-2/file.ext", pd.FullPath)
			require.Equal(t, "/dir-1/dir-2/", pd.Location)
			require.Equal(t, "/", pd.Root)
			require.Equal(t, "dir-1/dir-2/", pd.Directory)
			require.Equal(t, 2, len(pd.DirectoryParts))
			require.Equal(t, []string{"dir-1/", "dir-2/"}, pd.DirectoryParts)
			require.Equal(t, "file.ext", pd.Entry)
			require.Equal(t, "file", pd.Stem)
			require.Equal(t, ".ext", pd.Extension)

			require.Equal(t, 0, pd.NumberOfDotsDirectoryParts())

			checkPathDescriptorElements(t, pd)
		}
	}
}
