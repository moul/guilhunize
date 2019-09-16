package guilhunize

import "strings"

var mapping = [][2]string{
	{"refactors", "refractos"},
	{"refactos", "refractos"},
	{"refactoré", "refractoré"},
	{"refactor", "refracto"},
	{"refacto", "refracto"},
	{"scaleway", "scalaway"},
	{"interface", "i-love-interface"},
	{"déchiffrement", "désenchiffrement"},
	{" chiffrement", " cryptage"},
}

func Guilhunize(input string) string {
	// FIXME: refactor to use words instead of regexps
	input = strings.ToLower(input)
	for _, pair := range mapping {
		input = strings.Replace(input, pair[0], pair[1], -1)
	}
	return input
}

// FIXME: implement "Unguilhunize"
