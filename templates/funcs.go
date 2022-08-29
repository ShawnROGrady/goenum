package templates

func isAsciiLower(r rune) bool {
	return r >= 97 && r <= 122
}

func isAsciiUpper(r rune) bool {
	return r >= 65 && r <= 90
}

func unexported(name string) string {
	rs := []rune(name)
	if !isAsciiUpper(rs[0]) {
		return name
	}

	rs[0] = rs[0] + 32
	return string(rs)
}
