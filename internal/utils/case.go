package utils

import "unicode"

func MixCase(value string) string {
	runes := []rune(value)
	for i, currentRune := range runes {
		if i%2 == 0 {
			runes[i] = unicode.ToUpper(currentRune)
			continue
		}

		runes[i] = unicode.ToLower(currentRune)
	}

	return string(runes)
}
