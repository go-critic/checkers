package checker_test

func ifStmtOK() (int, error) {
	x, err := returnsIntAndError()
	if err != nil {
		return 0, err
	}
	if err := returnsError(); err != nil {
		return 0, err
	}
	if err2 := err; err2 != nil {
		return x, err2
	}
	return x, nil
}
