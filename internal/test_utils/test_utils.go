package test_utils

import (
	types "github.com/synesissoftware/libpath.Go/parse/common"

	"github.com/stretchr/testify/require"

	"strings"
	"testing"
)

func CheckPathDescriptorElements(t *testing.T, pd types.PathDescriptor) {
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

/* Helper functions and types */

type ClassificationResult struct {
	Path              string
	ParseFlags        types.ParseFlags
	Classification    types.Classification
	Root              string
	PathRootStripped  string
	FirstBadCharIndex int
}

func ClassificationResultBadFromPath(path string, classification types.Classification, firstBadCharIndex int) ClassificationResult {
	return ClassificationResult{
		Path:              path,
		ParseFlags:        0,
		Classification:    classification,
		Root:              "",
		PathRootStripped:  "",
		FirstBadCharIndex: firstBadCharIndex,
	}
}

func ClassificationResultBadFromPathAndFlags(path string, parseFlags types.ParseFlags, classification types.Classification, firstBadCharIndex int) ClassificationResult {
	return ClassificationResult{
		Path:              path,
		ParseFlags:        parseFlags,
		Classification:    classification,
		Root:              "",
		PathRootStripped:  "",
		FirstBadCharIndex: firstBadCharIndex,
	}
}

func ClassificationResultGoodFromPath(path string, classification types.Classification, root, pathRootStripped string) ClassificationResult {
	return ClassificationResult{
		Path:              path,
		ParseFlags:        types.ParseFlags_None,
		Classification:    classification,
		Root:              root,
		PathRootStripped:  pathRootStripped,
		FirstBadCharIndex: -1,
	}
}

func ClassificationResultGoodFromPathAndFlags(path string, parseFlags types.ParseFlags, classification types.Classification, root, pathRootStripped string) ClassificationResult {
	return ClassificationResult{
		Path:              path,
		ParseFlags:        parseFlags,
		Classification:    classification,
		Root:              root,
		PathRootStripped:  pathRootStripped,
		FirstBadCharIndex: -1,
	}
}
