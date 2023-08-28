package template

import "github.com/LNKLEO/OMP/src/regex"

func matchP(pattern, text string) bool {
	return regex.MatchString(pattern, text)
}

func replaceP(pattern, text, replaceText string) string {
	return regex.ReplaceAllString(pattern, text, replaceText)
}
