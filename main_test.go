package main

import "os"

func Example() {
	os.Args = []string{"", "salut, j'aime bien faire des interfaces dans les refactors, surtout lorsqu'il faut faire du chiffrement ou du déchiffrement."}
	main()
	// Output: salut, j'aime bien faire des i-love-interfaces dans les refractos, surtout lorsqu'il faut faire du cryptage ou du désenchiffrement.
}
