package main

import "os"

func place(name string, data []byte) {
	var file, _ = os.Create(name)
	defer file.Close()
	file.Write(data)
}
