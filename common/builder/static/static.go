package static

import _ "embed"

//go:embed agents/linux_amd64
var AgentLinuxAmd64 []byte

//go:embed agents/linux_386
var AgentLinux386 []byte

//go:embed agents/linux_arm
var AgentLinuxArm []byte

//go:embed agents/linux_arm64
var AgentLinuxArm64 []byte

//go:embed agents/windows_amd64.exe
var AgentWindowsAmd64 []byte

//go:embed agents/windows_386.exe
var AgentWindows386 []byte

//go:embed agents/windows_arm.exe
var AgentWindowsArm []byte

//go:embed agents/windows_arm64.exe
var AgentWindowsArm64 []byte

//go:embed agents/solaris_amd64
var AgentSolarisAmd64 []byte

//go:embed agents/openbsd_amd64
var AgentOpenBsdAmd64 []byte

//go:embed agents/openbsd_386
var AgentOpenBsd386 []byte

//go:embed agents/openbsd_arm
var AgentOpenBsdArm []byte

//go:embed agents/openbsd_arm64
var AgentOpenBsdArm64 []byte

//go:embed agents/darwin_amd64
var AgentDarwinAmd64 []byte

//go:embed agents/darwin_arm64
var AgentDarwinArm64 []byte

//go:embed agents/modular_linux_amd64
var AgentModularLinuxAmd64 []byte

//go:embed agents/modular_linux_386
var AgentModularLinux386 []byte

//go:embed agents/modular_linux_arm
var AgentModularLinuxArm []byte

//go:embed agents/modular_linux_arm64
var AgentModularLinuxArm64 []byte

//go:embed agents/modular_windows_amd64.exe
var AgentModularWindowsAmd64 []byte

//go:embed agents/modular_windows_386.exe
var AgentModularWindows386 []byte

//go:embed agents/modular_windows_arm.exe
var AgentModularWindowsArm []byte

//go:embed agents/modular_windows_arm64.exe
var AgentModularWindowsArm64 []byte
