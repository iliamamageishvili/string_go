package strings

func CountUnicodeChars(str string) int {
	return len([]rune(str))
}
