package main

import (
	_ "embed"
)

//go:embed components/atila.js
var atila []byte

//go:embed components/pyramid.js
var pyramid []byte

//go:embed components/ra.js
var ra []byte

//go:embed components/apep
var apep []byte

//go:embed components/varvara_python.py
var varvaraPython []byte

//go:embed components/varvara_dotnet.exe
var varvaraDotnet []byte

//go:embed components/varvara_cpp
var varvaraCpp []byte

//go:embed components/components_chain_sol_Botchain.abi
var chainAbi []byte

//go:embed components/components_chain_sol_Botchain.bin
var chainBin []byte

func main() {
	print(version())

	// place("outTest/chain.abi", chainAbi)
	// place("outTest/chain.bin", chainBin)
	// place("outTest/atila.js", atila)
	// place("outTest/pyramid.js", atila)
	// place("outTest/ra.js", atila)
	// place("outTest/apep", apep)
	// place("outTest/varvara.py", varvaraPython)
	// place("outTest/varvaraDotnet.exe", varvaraDotnet)
	// place("outTest/varvaraCpp", varvaraCpp)
}
