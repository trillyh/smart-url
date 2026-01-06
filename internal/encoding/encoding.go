package encoding

import (
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Takes a uint64 ID an returns a Base62 encoded string
func Encode(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}

	var builder strings.Builder
	base := uint64(len(alphabet))

	for number > 0 {
		remainder := number % base
		builder.WriteByte(alphabet[remainder])
		number = number / base
	}

	return
}

func reverse(s string) string {
	runes := []rune(s)
	i, j := 0, len(s)-1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

}
