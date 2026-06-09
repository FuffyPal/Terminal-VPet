package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

const (
	version    = "0.8.1"
	maxhealthy = 100
	minhealthy = 0
	maxhungry  = 100
	minhungry  = 0
)

type Pet struct {
	Name    string `json:"name"`
	Healthy int    `json:"healthy"`
	Hunger  int    `json:"hunger"`
	Life    bool   `json:"life"`
}

type Food struct {
	name   string
	energy int
}

func main() {
	nameFlag := flag.String("name", "", "Pet's name")
	statusFlag := flag.Bool("status", false, "Pet's status")
	versionFlag := flag.Bool("version", false, "App's version")
	eatFlag := flag.Int("eat", 0, "pet eating x(1,2,3,4)")
	savefileFlag := flag.String("savefile", "", "Custom path for save file")
	flag.Usage = func() {
		fmt.Println("\n╔===================================================================╗")
		fmt.Println("║       🐾  VIRTUAL PET GAME (", version, ") CLI HELP  🐾                 ║")
		fmt.Println("╚===================================================================╝")
		fmt.Println("Usage:")
		fmt.Println("  go run . [options]\n")
		fmt.Println("Options:")
		fmt.Println("  --name <string>   💡 Sets your pet's name (e.g., --name=pamuk)")
		fmt.Println("  --status          📊 Shows current stats of your pet and exits")
		fmt.Println("  --eat <1-4>       🍖 Feeds your pet with a food from the market:")
		fmt.Println("                      [1] Omlet 🍳  [2] Fish 🐟")
		fmt.Println("                      [3] Meat 🥩   [4] Apple 🍎")
		fmt.Println("  --savefile <path> 💾 Custom path for save file")
		fmt.Println("╔===================================================================╗")
		fmt.Println("║Examples:                                                          ║")
		fmt.Println("║  go run . --name=boncuk --status                                  ║")
		fmt.Println("║  go run . --eat=1                                                 ║")
		fmt.Println("╚===================================================================╝\n")
	}
	flag.Parse()

	if *versionFlag {
		fmt.Printf("vpet version: %s \n", version)
		return
	}

	var saveFilePath string
	if *savefileFlag != "" {
		saveFilePath = *savefileFlag
	} else {
		saveDir, err := GetSaveDir()
		if err != nil {
			fmt.Println("Err:", err)
			return
		}
		saveFilePath = filepath.Join(saveDir, "savefile.json")
	}

	myPet, err := LoadPet(saveFilePath)
	if err != nil {
		fmt.Println("Err (LoadPet):", err)
		return
	}

	var foodList = []Food{
		{name: "Omlet 🍳", energy: 20},
		{name: "Fish 🐟", energy: 15},
		{name: "Meat 🥩", energy: 35},
		{name: "Apple 🍎", energy: 5},
	}

	if *nameFlag != "" {
		myPet.Name = *nameFlag
		fmt.Printf("\n✨ %s has successfully hatched! Let the adventure begin...\n", myPet.Name)
	}

	myPet.tick(5)

	if !myPet.Life {
		myPet.status()
		err = SavePet(saveFilePath, myPet)
		return
	}

	if *eatFlag >= 1 && *eatFlag <= len(foodList) {
		choosfood := foodList[*eatFlag-1]

		fmt.Printf("\nYou fed %s with %s!\n", myPet.Name, choosfood.name)

		myPet.eat(choosfood.energy)
	} else if *eatFlag != 0 {
		fmt.Println("Invalid choice! Your pet stayed hungry.")
	}

	if *statusFlag {
		myPet.status()
	}

	err = SavePet(saveFilePath, myPet)
	if err != nil {
		fmt.Println("Err (SavePet):", err)
	}
}
