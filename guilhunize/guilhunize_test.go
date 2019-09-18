package guilhunize

import (
	"fmt"
	"math/rand"
)

func ExampleGuilhunize() {
	input := "salut, j'aime bien faire des interfaces dans les refactors, surtout lorsqu'il faut faire du chiffrement ou du déchiffrement."
	fmt.Println(Guilhunize(input))
	// Output: salut, j'aime bien faire des i-love-interfaces dans les refractos, surtout lorsqu'il faut faire du cryptage ou du désenchiffrement.
}

func ExampleQuote() {
	rand.Seed(4242)
	fmt.Println(Quote())
	// Output: hein, quoi ?
}
