package radixtree

import "strings"

func makeString(runes []rune) string {
	var b strings.Builder
	for _, r := range runes {
		b.WriteRune(r)
	}
	return b.String()
}

func deleteCharAt(a []rune, i int) []rune {
	return append(a[:i], a[i+1:]...)
}