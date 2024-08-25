package main

import (
	"c2/core"
	"c2/db"
	"c2/srv"
	"flag"
	"fmt"
	"os"
)

func main() {
	dbName := flag.String("db-name", "xena.db", "Name of the XENA's database file.")

	flag.Parse()

	os.MkdirAll(core.PATH_DOWNLOADS, 0777)
	os.MkdirAll(core.PATH_MODULES, 0777)

	if err := core.InitKeys(); err != nil {
		fmt.Println("Fatal! Failed to initialize C2's cryptographic keys!")
		os.Exit(1)
	}

	if err := db.Init(*dbName); err != nil {
		fmt.Println("failed to initialize the database:", err.Error())
		os.Exit(1)
	}

	if err := srv.Start(); err != nil {
		fmt.Println("failed to initialize the C2 API:", err.Error())
		os.Exit(1)
	}
}
