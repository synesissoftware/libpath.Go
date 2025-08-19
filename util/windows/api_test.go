package windows_test

import (
	api "github.com/synesissoftware/libpath.Go/util/windows"

	"github.com/stretchr/testify/require"

	"testing"
)

/* Helper functions and types */

type classification_result struct {
	path                 string
	parseFlags           api.ParseFlags
	classification       api.Classification
	root                 string
	path_root_stripped   string
	first_bad_char_index int
}

func cr_bad_from_path(path string, classification api.Classification, first_bad_char_index int) classification_result {
	return classification_result{
		path:                 path,
		parseFlags:           0,
		classification:       classification,
		root:                 "",
		path_root_stripped:   "",
		first_bad_char_index: first_bad_char_index,
	}

}

func cr_bad_from_path_and_flags(path string, parseFlags api.ParseFlags, classification api.Classification, first_bad_char_index int) classification_result {
	return classification_result{
		path:                 path,
		parseFlags:           parseFlags,
		classification:       classification,
		root:                 "",
		path_root_stripped:   "",
		first_bad_char_index: first_bad_char_index,
	}

}

func cr_good_from_path(path string, classification api.Classification, root, path_root_stripped string) classification_result {
	return classification_result{
		path:                 path,
		parseFlags:           api.ParseFlags_None,
		classification:       classification,
		root:                 root,
		path_root_stripped:   path_root_stripped,
		first_bad_char_index: -1,
	}
}

func cr_good_from_path_and_flags(path string, parseFlags api.ParseFlags, classification api.Classification, root, path_root_stripped string) classification_result {
	return classification_result{
		path:                 path,
		parseFlags:           parseFlags,
		classification:       classification,
		root:                 root,
		path_root_stripped:   path_root_stripped,
		first_bad_char_index: -1,
	}
}

/* Test functions */

func Test_Basename(t *testing.T) {
	{
		require.Equal(t, ".", api.Basename("."))
	}

	{
		require.Equal(t, "abc.ext", api.Basename("abc.ext"))
		require.Equal(t, "abc", api.Basename("abc"))
		require.Equal(t, ".ext", api.Basename(".ext"))
	}

	{
		require.Equal(t, "abc.ext", api.Basename("\\abc.ext"))
		require.Equal(t, "abc", api.Basename("\\abc"))
		require.Equal(t, ".ext", api.Basename("\\.ext"))

		require.Equal(t, "abc.ext", api.Basename("/abc.ext"))
		require.Equal(t, "abc", api.Basename("/abc"))
		require.Equal(t, ".ext", api.Basename("/.ext"))
	}

	{
		require.Equal(t, "abc.ext", api.Basename("dir\\abc.ext"))
		require.Equal(t, "abc", api.Basename("dir\\abc"))
		require.Equal(t, ".ext", api.Basename("dir\\.ext"))

		require.Equal(t, "abc.ext", api.Basename("dir/abc.ext"))
		require.Equal(t, "abc", api.Basename("dir/abc"))
		require.Equal(t, ".ext", api.Basename("dir/.ext"))
	}

	{
		require.Equal(t, "abc.ext", api.Basename("\\dir\\abc.ext"))
		require.Equal(t, "abc.ext", api.Basename("\\dir/abc.ext"))
		require.Equal(t, "abc.ext", api.Basename("/dir\\abc.ext"))
		require.Equal(t, "abc.ext", api.Basename("/dir/abc.ext"))
	}
}

func Test_ClassifyRoot_1(t *testing.T) {

	cr_cases := []classification_result{
		cr_bad_from_path("", api.Empty, 0),

		cr_good_from_path("abc", api.Relative, "", "abc"),
		cr_good_from_path("abc.def", api.Relative, "", "abc.def"),

		cr_good_from_path("/", api.SlashRooted, "/", ""),
		cr_good_from_path(`\`, api.SlashRooted, `\`, ""),

		cr_good_from_path("/abc", api.SlashRooted, "/", "abc"),
		cr_good_from_path(`\abc`, api.SlashRooted, `\`, "abc"),

		cr_good_from_path("/abc.def", api.SlashRooted, "/", "abc.def"),
		cr_good_from_path(`\abc.def`, api.SlashRooted, `\`, "abc.def"),

		cr_good_from_path(`C:/dir/sub-dir/file.ext`, api.DriveLetterRooted, `C:/`, "dir/sub-dir/file.ext"),
		cr_good_from_path(`C:\dir\sub-dir\file.ext`, api.DriveLetterRooted, `C:\`, `dir\sub-dir\file.ext`),

		cr_good_from_path(`C:/`, api.DriveLetterRooted, `C:/`, ""),
		cr_good_from_path(`C:\`, api.DriveLetterRooted, `C:\`, ""),

		cr_good_from_path(`C:dir/sub-dir/file.ext`, api.DriveLetterRelative, "C:", "dir/sub-dir/file.ext"),
		cr_good_from_path(`C:dir\sub-dir\file.ext`, api.DriveLetterRelative, "C:", `dir\sub-dir\file.ext`),

		cr_good_from_path("C:", api.DriveLetterRelative, "C:", ""),
		cr_good_from_path("C:", api.DriveLetterRelative, "C:", ""),

		cr_good_from_path("~", api.Relative, "", "~"),
		cr_good_from_path("~/", api.Relative, "", "~/"),
		cr_good_from_path("~/abc", api.Relative, "", "~/abc"),

		cr_good_from_path_and_flags("~", api.ParseFlags_RecogniseTildeHome, api.HomeRooted, "~", ""),
		cr_good_from_path_and_flags("~/", api.ParseFlags_RecogniseTildeHome, api.HomeRooted, "~/", ""),
		cr_good_from_path_and_flags(`~\`, api.ParseFlags_RecogniseTildeHome, api.HomeRooted, `~\`, ""),
		cr_good_from_path_and_flags("~/abc", api.ParseFlags_RecogniseTildeHome, api.HomeRooted, "~/", "abc"),
		cr_good_from_path_and_flags(`~\abc`, api.ParseFlags_RecogniseTildeHome, api.HomeRooted, `~\`, "abc"),

		cr_bad_from_path("|a", api.InvalidChars, 0),
		cr_bad_from_path("a|", api.InvalidChars, 1),
		cr_good_from_path("abcdef|", api.Relative, "", "abcdef|"),
		cr_bad_from_path("C|:a", api.InvalidChars, 1),
		cr_bad_from_path("C:|a", api.InvalidChars, 2),
		cr_bad_from_path_and_flags("~|", api.ParseFlags_RecogniseTildeHome, api.InvalidChars, 1),
	}

	for _, cr_case := range cr_cases {

		classification, root, path_root_stripped, first_bad_char_index := api.ClassifyRoot(cr_case.path, cr_case.parseFlags)

		require.Equal(t, cr_case.classification, classification, "actual classification '%[1]s' (%[1]d) does not match expected classificaton '%[2]s' (%[2]d) with input path '%s' and parseFlags '0x%08x'", classification, cr_case.classification, cr_case.path, cr_case.parseFlags)
		require.Equal(t, cr_case.root, root, "actual root '%t' does not match expected root '%t' with input path '%s' and parseFlags '0x%08x'", root, cr_case.root, cr_case.path, cr_case.parseFlags)
		require.Equal(t, cr_case.path_root_stripped, path_root_stripped, "actual path_root_stripped '%t' does not match expected path_root_stripped '%t' with input path '%s' and parseFlags '0x%08x'", path_root_stripped, cr_case.path_root_stripped, cr_case.path, cr_case.parseFlags)
		require.Equal(t, cr_case.first_bad_char_index, first_bad_char_index, "actual first_bad_char_index '%t' does not match expected first_bad_char_index '%t' with input path '%s' and parseFlags '0x%08x'", first_bad_char_index, cr_case.first_bad_char_index, cr_case.path, cr_case.parseFlags)
	}
}
