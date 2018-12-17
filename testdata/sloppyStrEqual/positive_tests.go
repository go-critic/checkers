package checker_test

import "strings"

func foo(x, y string) {
	_ = strings.ToLower(x) == y

	_ = strings.ToLower(x) == strings.ToLower(y)

	_ = x == strings.ToLower(y)

	_ = strings.ToLower(x) == "y"

	_ = strings.ToLower(x) == strings.ToLower("y")

	_ = x == strings.ToLower("y")
}

func bar(x, y string) {
	_ = strings.ToUpper(x) == y

	_ = strings.ToUpper(x) == strings.ToUpper(y)

	_ = x == strings.ToUpper(y)

	_ = strings.ToUpper(x) == "y"

	_ = strings.ToUpper(x) == strings.ToUpper("y")

	_ = x == strings.ToUpper("y")
}
