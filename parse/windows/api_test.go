package windows

import (
	test_utils "github.com/synesissoftware/libpath.Go/internal/test_utils"
	"github.com/synesissoftware/libpath.Go/parse/common"
	api "github.com/synesissoftware/libpath.Go/parse/windows"

	"github.com/stretchr/testify/require"

	"testing"
)

func Test_ClassifyRoot_1(t *testing.T) {

	cr_cases := []test_utils.ClassificationResult{
		test_utils.ClassificationResultBadFromPath("", common.Empty, 0),

		test_utils.ClassificationResultGoodFromPath("abc", common.Relative, "", "abc"),
		test_utils.ClassificationResultGoodFromPath("abc.def", common.Relative, "", "abc.def"),

		test_utils.ClassificationResultGoodFromPath("/", common.SlashRooted, "/", ""),
		test_utils.ClassificationResultGoodFromPath(`\`, common.SlashRooted, `\`, ""),

		test_utils.ClassificationResultGoodFromPath("/abc", common.SlashRooted, "/", "abc"),
		test_utils.ClassificationResultGoodFromPath(`\abc`, common.SlashRooted, `\`, "abc"),

		test_utils.ClassificationResultGoodFromPath("/abc.def", common.SlashRooted, "/", "abc.def"),
		test_utils.ClassificationResultGoodFromPath(`\abc.def`, common.SlashRooted, `\`, "abc.def"),

		test_utils.ClassificationResultGoodFromPath(`C:/dir/sub-dir/file.ext`, common.DriveLetterRooted, `C:/`, "dir/sub-dir/file.ext"),
		test_utils.ClassificationResultGoodFromPath(`C:\dir\sub-dir\file.ext`, common.DriveLetterRooted, `C:\`, `dir\sub-dir\file.ext`),

		test_utils.ClassificationResultGoodFromPath(`C:/`, common.DriveLetterRooted, `C:/`, ""),
		test_utils.ClassificationResultGoodFromPath(`C:\`, common.DriveLetterRooted, `C:\`, ""),

		test_utils.ClassificationResultGoodFromPath(`C:dir/sub-dir/file.ext`, common.DriveLetterRelative, "C:", "dir/sub-dir/file.ext"),
		test_utils.ClassificationResultGoodFromPath(`C:dir\sub-dir\file.ext`, common.DriveLetterRelative, "C:", `dir\sub-dir\file.ext`),

		test_utils.ClassificationResultGoodFromPath("C:", common.DriveLetterRelative, "C:", ""),
		test_utils.ClassificationResultGoodFromPath("C:", common.DriveLetterRelative, "C:", ""),

		test_utils.ClassificationResultGoodFromPath("~", common.Relative, "", "~"),
		test_utils.ClassificationResultGoodFromPath("~/", common.Relative, "", "~/"),
		test_utils.ClassificationResultGoodFromPath("~/abc", common.Relative, "", "~/abc"),

		test_utils.ClassificationResultGoodFromPathAndFlags("~", common.ParseFlags_RecogniseTildeHome, common.HomeRooted, "~", ""),
		test_utils.ClassificationResultGoodFromPathAndFlags("~/", common.ParseFlags_RecogniseTildeHome, common.HomeRooted, "~/", ""),
		test_utils.ClassificationResultGoodFromPathAndFlags(`~\`, common.ParseFlags_RecogniseTildeHome, common.HomeRooted, `~\`, ""),
		test_utils.ClassificationResultGoodFromPathAndFlags("~/abc", common.ParseFlags_RecogniseTildeHome, common.HomeRooted, "~/", "abc"),
		test_utils.ClassificationResultGoodFromPathAndFlags(`~\abc`, common.ParseFlags_RecogniseTildeHome, common.HomeRooted, `~\`, "abc"),

		test_utils.ClassificationResultBadFromPath("|a", common.InvalidChars, 0),
		test_utils.ClassificationResultBadFromPath("a|", common.InvalidChars, 1),
		test_utils.ClassificationResultGoodFromPath("abcdef|", common.Relative, "", "abcdef|"),
		test_utils.ClassificationResultBadFromPath("C|:a", common.InvalidChars, 1),
		test_utils.ClassificationResultBadFromPath("C:|a", common.InvalidChars, 2),
		test_utils.ClassificationResultBadFromPathAndFlags("~|", common.ParseFlags_RecogniseTildeHome, common.InvalidChars, 1),
	}

	for _, cr_case := range cr_cases {

		classification, root, pathRootStripped, firstBadCharIndex := api.ClassifyRoot(cr_case.Path, cr_case.ParseFlags)

		require.Equal(t, cr_case.Classification, classification, "actual classification '%[1]s' (%[1]d) does not match expected classificaton '%[2]s' (%[2]d) with input path '%s' and parseFlags '0x%08x'", classification, cr_case.Classification, cr_case.Path, cr_case.ParseFlags)
		require.Equal(t, cr_case.Root, root, "actual root '%t' does not match expected root '%t' with input path '%s' and parseFlags '0x%08x'", root, cr_case.Root, cr_case.Path, cr_case.ParseFlags)
		require.Equal(t, cr_case.PathRootStripped, pathRootStripped, "actual pathRootStripped '%t' does not match expected pathRootStripped '%t' with input path '%s' and parseFlags '0x%08x'", pathRootStripped, cr_case.PathRootStripped, cr_case.Path, cr_case.ParseFlags)
		require.Equal(t, cr_case.FirstBadCharIndex, firstBadCharIndex, "actual firstBadCharIndex '%t' does not match expected firstBadCharIndex '%t' with input path '%s' and parseFlags '0x%08x'", firstBadCharIndex, cr_case.FirstBadCharIndex, cr_case.Path, cr_case.ParseFlags)
	}
}
