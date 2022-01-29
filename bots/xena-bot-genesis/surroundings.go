package main

import "os"

func version() string {
	content, err := os.ReadFile("/proc/version")
	if err != nil {
		print(err)
	}
	return string(content)
}

func shortVersion() string {
	content, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		print(err)
	}
	return string(content)
}
