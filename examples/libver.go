package main

import (
	libpath "github.com/synesissoftware/libpath.Go"
	"github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("libpath v%s\n", ver2go.CalcVersionString(libpath.VersionMajor, libpath.VersionMinor, libpath.VersionPatch, libpath.VersionAB))
	fmt.Printf("ver2go v%s\n", ver2go.CalcVersionString(ver2go.VersionMajor, ver2go.VersionMinor, ver2go.VersionPatch, ver2go.VersionAB))
}
