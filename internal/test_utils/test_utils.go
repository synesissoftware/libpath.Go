package test_utils

import (
	"github.com/synesissoftware/libpath.Go/parse"

	"github.com/stretchr/testify/require"

	"strings"
	"testing"
)

func CheckPathDescriptorElements(t *testing.T, pd parse.PathDescriptor) {
	t.Helper()

	// Stem + Extension => EntryName
	{
		require.Equal(t, pd.EntryName, pd.Stem+pd.Extension)
	}

	// Location + EntryName => FullPath
	{
		require.Equal(t, pd.FullPath, pd.Location+pd.EntryName)
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
