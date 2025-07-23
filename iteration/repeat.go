package iteration

import "strings"

func Repeat(char string, repeat int) string {
	var res strings.Builder
	for range repeat {
		res.WriteString(char)
	}
	return res.String()
}
