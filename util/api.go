package util

import "strings"

/* NOTE: if base name ends with '.', then no extension is obtained;
 * otherwise has the normal semantics
 */
func SplitBasename(base_name string) (file_stem, file_ext string) {

	lix_dot := strings.LastIndexByte(base_name, '.')

	if lix_dot < 0 || lix_dot == len(base_name)-1 {

		file_stem = base_name
	} else {

		file_stem = base_name[:lix_dot]
		file_ext = base_name[lix_dot:]
	}

	return
}
