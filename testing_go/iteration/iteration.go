package iteration

const repeatCount = 5

func ForLoop(character string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
