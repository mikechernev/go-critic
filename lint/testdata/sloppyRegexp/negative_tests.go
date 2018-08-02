package checker_test

import "regexp"

func g() {
	re := regexp.MustCompile("(TO)(DO)")
	_ = re
}
