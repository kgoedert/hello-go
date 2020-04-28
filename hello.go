package hello

func Hello() string {
	return "Hello, world."
}

func uncovered() string {
	return "not covered by tests"
}
