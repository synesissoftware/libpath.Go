package libpath_test

import (
	"github.com/synesissoftware/libpath.Go"

	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 0
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 4
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, libpath.VersionMajor)
	require.Equal(t, Expected_VersionMinor, libpath.VersionMinor)
	require.Equal(t, Expected_VersionPatch, libpath.VersionPatch)
	require.Equal(t, Expected_VersionAB, libpath.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0000_0000_0004), libpath.Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.0.0", libpath.VersionString())
}
