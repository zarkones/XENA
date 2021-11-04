package main

import _ "embed"

//go:embed components/chain.sol
var chain []byte

//go:embed components/atila.js
var atila []byte

//go:embed components/apep
var apep []byte

func main() {
	place("outTest/test.sol", chain)
	place("outTest/atila.js", atila)
	place("outTest/apep", apep)
}
