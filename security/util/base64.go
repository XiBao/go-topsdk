package util

import (
	"encoding/base64"
	"sort"
)

var sorted_base64_runes = []rune("+/0123456789=ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func searchRunes(a []rune, x rune) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func IsBase64(s string) bool {
	if s == "" {
		return false
	}
	for _, c := range []rune(s) {
		pos := searchRunes(sorted_base64_runes, c)
		if pos < 0 || pos >= len(sorted_base64_runes) || sorted_base64_runes[pos] != c {
			return false
		}
	}
	return true
}

func IsLetterOrDigit(s string) bool {
	runes := []rune(s)
	if len(runes) != 1 {
		return false
	}
	code := int(runes[0])
	return 0 <= code && code <= 127
}
