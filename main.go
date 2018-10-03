package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

const version = "v0.0"

var (
	showHelp    bool
	mealToAdd   string
	newFoodItem string
)

func usage(appName, version string) {
	fmt.Printf("Usage: %s [OPTIONS]\n", appName)
	fmt.Println("OPTIONS:\n")
	flag.VisitAll(func(f *flag.Flag) {
		if len(f.Name) > 1 {
			fmt.Printf("\t-%s, -%s\t%s\n", f.Name[0:1], f.Name, f.Usage)
		}
	})
	fmt.Printf("\nVersion: %s\n", version)
}

func init() {

	// TODO: create db if it doesn-t exist
	// read flags
	flag.BoolVar(&showHelp, "h", false, "dispaly help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.StringVar(&mealToAdd, "a", mealToAdd, "comma separeted string of ingridients eaten")
	flag.StringVar(&mealToAdd, "add", mealToAdd, "comma separeted string of ingridients eaten")
	flag.StringVar(&newFoodItem, "n", newFoodItem, "name of the new food item and the groups it is part of (eg \"foodName 1234567\")")
	flag.StringVar(&newFoodItem, "new", newFoodItem, "name of the new food item and the groups it is part of (eg \"foodName 1234567\")")
}

func main() {
	// TODO gui
	appName := path.Base(os.Args[0])
	flag.Parse()

	if showHelp == true {
		usage(appName, version)
	}

	if mealToAdd != "" {
		// TODO: use the controller to add the meal to the model
		fmt.Printf("Specified %s as meal to add\n", mealToAdd)
	}
	if newFoodItem != "" {
		// TODO: use the controller to add the item after parsing the flag
		fmt.Printf("Specified %s as a new food item\n", newFoodItem)
	}
}
