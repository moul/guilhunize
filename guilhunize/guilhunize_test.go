package guilhunize

import (
	"fmt"
)

func ExampleGuilhunize() {
	input := "salut, j'aime bien faire des interfaces dans les refactors, surtout lorsqu'il faut faire du chiffrement ou du déchiffrement."
	fmt.Println(Guilhunize(input))
	// Output: salut, j'aime bien faire des i-love-interfaces dans les refractos, surtout lorsqu'il faut faire du cryptage ou du désenchiffrement.
}

func ExampleQuote() {
	fmt.Println(Quote())
}
