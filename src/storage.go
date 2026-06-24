package main

import (
	"encoding/json"
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

func SavePet(filePath string, pet *Pet) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("could not create directory: %v", err)
	}

	data, err := json.MarshalIndent(pet, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal pet: %v", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write save file: %v", err)
	}

	return nil
}

func LoadPet(filePath string) (*Pet, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			defaultPet := &Pet{
				Name:         "",
				Healthy:      100,
				Hunger:       0,
				Life:         true,
				Current_Time: "",
			}
			err = SavePet(filePath, defaultPet)
			if err != nil {
				return nil, fmt.Errorf("could not create default save file: %v", err)
			}
			return defaultPet, nil
		}
		return nil, fmt.Errorf("could not read save file: %v", err)
	}

	var pet Pet
	err = json.Unmarshal(data, &pet)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal save file: %v", err)
	}

	return &pet, nil
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
