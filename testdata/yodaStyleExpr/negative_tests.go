package checker_test

func g1() {
	var a int
	if a == 10 {
	}

	if 10 > a {
	}
}

func g2() bool {
	f := func() interface{} { return nil }
	return f() != nil
}
