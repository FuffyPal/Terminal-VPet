package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetSaveDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get not home folder: %v", err)
	}

	var targetDir string

	switch runtime.GOOS {
	case "windows":
		targetDir = filepath.Join(homeDir, "AppData", "Roaming", "pal-pet")
	case "darwin":
		targetDir = filepath.Join(homeDir, "Library", "Application Support", "pal-pet")
	default:
		targetDir = filepath.Join(homeDir, ".local", "share", "pal-pet")
	}

	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		return "", fmt.Errorf("save file not create: %v", err)
	}

	return targetDir, nil
}

// func hi() {
// 	saveDir, err := GetSaveDir()
// 	if err != nil {
// 		fmt.Println("Hata:", err)
// 		return
// 	}
// 	saveFilePath := filepath.Join(saveDir, "savefile.dat") // savefile path

// 	fmt.Println("OS:", runtime.GOOS)
// 	fmt.Println("save path:", saveFilePath)

// }
