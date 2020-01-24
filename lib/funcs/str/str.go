package str

import (
	"crypto/sha256"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"unicode"
)

func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func GeneratePassword() (string, error) {
	pass, err := password.Generate(8, 10, 0, false, false)
	return pass, err
}

func HashPassword(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func ParseByteToList(in []byte) []string {
	var l []string
	var t string
	i := 0
	for i < len(in) {
		if string(in)[i:i+1] == "," || i+1 == len(in) {
			l = append(l, t)
			t = ""
			i = i + 1
			continue
		}
		if string(in)[i:i+1] == "\n" || string(in)[i:i+1] == "\"" || string(in)[i:i+1] == "[" || string(in)[i:i+1] == "]" {
			i = i + 1
			continue
		}
		t = t + string(in)[i:i+1]
		i = i + 1
	}
	return l
}
