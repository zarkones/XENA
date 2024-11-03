package main

import (
	"c2/core"
	"c2/core/analyzer"
	"c2/core/env"
	"c2/core/proxy"
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

	if err := env.Validate(); err != nil {
		fmt.Println("environment error:", err)
		os.Exit(1)
	}

	if err := core.InitKeys(); err != nil {
		fmt.Println("Fatal! Failed to initialize C2's cryptographic keys!", err)
		os.Exit(1)
	}

	if err := db.Init(*dbName); err != nil {
		fmt.Println("failed to initialize the database:", err)
		os.Exit(1)
	}

	go func() {
		if err := proxy.Start(env.PROXY_HOST, env.PROXY_PORT, core.CertPath, core.PrivateKey); err != nil {
			fmt.Println("proxy error:", err)
			return
		}
	}()

	go analyzer.Start()

	if err := srv.Start(); err != nil {
		fmt.Println("failed to initialize the C2 API:", err)
		os.Exit(1)
	}
}
