package checker_test

import "regexp"

func f() {
	/// regexp can be rewritten with strings.HasPrefix
	re := regexp.MustCompile("^TODO")
	_ = re
}
