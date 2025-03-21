package main

import (
	"c2/core"
	"c2/core/env"
	"c2/core/proxy"
	"c2/db"
	"c2/models"
	pipelinesRepo "c2/repos/pipelines"
	"c2/srv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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

	// Init default pipelines:
	initDefaultPipelines()

	go func() {
		if err := proxy.Start(env.PROXY_HOST, env.PROXY_PORT, core.CertPath, core.PrivateKey); err != nil {
			fmt.Println("proxy error:", err)
			return
		}
	}()

	if err := srv.Start(); err != nil {
		fmt.Println("failed to initialize the C2 API:", err)
		os.Exit(1)
	}
}

func initDefaultPipelines() (err error) {
	fileRecords, err := os.ReadDir(core.PATH_DEFAULT_PIPELINES)
	if err != nil {
		return err
	}
	for _, record := range fileRecords {
		if record.IsDir() {
			continue
		}
		filePath := filepath.Join(core.PATH_DEFAULT_PIPELINES, record.Name())
		rawPipelines, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		var pipelines []models.Pipeline
		if err := json.Unmarshal(rawPipelines, &pipelines); err != nil {
			return err
		}
		for _, pipeline := range pipelines {
			if err := pipelinesRepo.Upsert(&pipeline); err != nil {
				return err
			}
		}
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}
	return nil
}
