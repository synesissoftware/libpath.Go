package main

import (
	clasp "github.com/synesissoftware/CLASP.Go"
	libclimate "github.com/synesissoftware/libCLImate.Go"
	libpath_parse "github.com/synesissoftware/libpath.Go/parse"
	libpath_common "github.com/synesissoftware/libpath.Go/parse/common"
	libpath_parse_unix "github.com/synesissoftware/libpath.Go/parse/unix"
	libpath_parse_windows "github.com/synesissoftware/libpath.Go/parse/windows"

	"fmt"
	"os"
	"slices"
	"strings"
)

var (
	OperatingSystem_Ambient = "ambient"
	OperatingSystem_Unix    = "unix"
	OperatingSystem_Windows = "windows"
	OperatingSystems        = []string{
		OperatingSystem_Ambient,
		OperatingSystem_Unix,
		OperatingSystem_Windows,
	}
)

func main() {

	// Specify specifications, parse, and checking standard flags

	parseFlags := libpath_common.ParseFlags_None

	flag_RecogniseTildeHome := clasp.Flag("--recognise-tilde-home").SetAliases("-t", "--recognize-tilde-home").SetHelp("Interprets a leading '~' as representing the home directory")
	option_OperatingSystem := clasp.Option("--operating-system").SetAlias("-o").SetHelp("Specifies the operating system interpretation to be applied to the path").SetValues(OperatingSystems...)

	operatingSystem := OperatingSystem_Ambient

	climate, _ := libclimate.Init(func(cl *libclimate.Climate) (err error) {

		cl.Version = "0.0.0"

		cl.AddFlagFunc(flag_RecogniseTildeHome, func() {
			parseFlags |= libpath_common.ParseFlags_RecogniseTildeHome
		})

		cl.AddOptionFunc(option_OperatingSystem, func(o *clasp.Argument, a *clasp.Specification) {

			operatingSystem = strings.ToLower(o.Value)

			if !slices.Contains(OperatingSystems, operatingSystem) {
				cl.Abort(fmt.Sprintf("invalid value '%s' for option %s; use --help for usage", o.Value, option_OperatingSystem.Name), nil)
			}
		})
		cl.AddAlias("--operating-system="+OperatingSystem_Unix, "-u")
		cl.AddAlias("--operating-system="+OperatingSystem_Windows, "-w")

		cl.InfoLines = []string{
			"libPath.Go Examples",
			"",
			":version:",
			"",
		}

		cl.ValueNames = []string{"path"}
		cl.ValuesConstraint = []int{1}
		cl.ValuesString = "<path>"

		return nil
	}, libclimate.InitFlag_PanicOnFailure)

	r, _ := climate.ParseAndVerify(os.Args, libclimate.ParseFlag_PanicOnFailure)

	path := r.Values[0].Value

	var parsed_path libpath_common.PathDescriptor
	var err error

	switch operatingSystem {
	case OperatingSystem_Ambient:

		parsed_path, err = libpath_parse.ParsePathStringFlags(path, "", parseFlags)
	case OperatingSystem_Unix:

		parsed_path, err = libpath_parse_unix.ParsePathStringFlags(path, "", parseFlags)
	case OperatingSystem_Windows:

		parsed_path, err = libpath_parse_windows.ParsePathStringFlags(path, "", parseFlags)
	}

	if err != nil {
		climate.Abort(fmt.Sprintf("failed to parse path '%s'", path), err)
	}

	fmt.Printf("path '%s' parses as follows:\n\n\t%s\n", path, parsed_path)
}
