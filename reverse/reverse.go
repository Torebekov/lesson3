package reverse

import "strings"

func Reverse(s string) string {
	ans := []rune(s)
	var d strings.Builder
	for i := range ans{
		d.WriteString(string(ans[len(ans)-1-i]))
	}
	return d.String()
}
