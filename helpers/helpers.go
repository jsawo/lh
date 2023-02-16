package helpers

func Truncate(str string, length int) (truncated string) {
	if length <= 0 {
		return
	}
	for i, char := range str {
		if i >= length {
			break
		}
		truncated += string(char)
	}
	return
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
