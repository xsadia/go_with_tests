package iteration

import s "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += s.ToUpper(character)
	}

	return repeated
}

func main() {
	Repeat("a", 5)
}
