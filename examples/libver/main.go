package main

import (
	libpath "github.com/synesissoftware/libpath.Go"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("libpath v%s\n", libpath.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
