package guilhunize

import (
	"math/rand"
	"strings"
)

var mapping = [][2]string{
	{"refactors", "refractos"},
	{"refactos", "refractos"},
	{"refactoré", "refractoré"},
	{"refactor", "refracto"},
	{"refacto", "refracto"},
	{"scaleway", "scalaway"},
	{"bytecode", "bitcode"},
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

var quotes = []string{
	"hein, quoi ?",
	"putain, fait chier catalina",
	"vous etes qui ?",
	"il etait cavalier, c'etait un mauvais soldat",
}

func Quote() string {
	return quotes[rand.Intn(len(quotes))]
}

// FIXME: implement "Unguilhunize"
