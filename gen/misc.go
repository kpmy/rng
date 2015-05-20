package gen

import (
	"strings"
)

func capitalize(s string) (ret string) {
	if s != "" {
		nam := []rune("")
		nam = []rune(s)
		nam[0] = []rune(strings.ToUpper(string([]rune{nam[0]})))[0]
		ret = string(nam)
	}
	return
}
