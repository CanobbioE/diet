package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/boltdb/bolt"
)

const (
	version = "v0.0"
)

var (
	showHelp    bool
	mealToAdd   string
	newFoodItem string
	debug       bool
	db          *bolt.DB
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
	// Creates db if it doesn't exist
	db, err := bolt.Open("diet.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Creates the two buckets, "foodBucket" and "dietBucket" if they doesn't exists.
	db.Update(func(tx *bolt.Tx) error {
		foodBucket, err := tx.CreateBucketIfNotExists([]byte("Food"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		if foodBucket.Stats().KeyN == 0 {
			// TODO  populate bucket
		}

		dietBucket, err := tx.CreateBucketIfNotExists([]byte("Diet"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		// populates dietbucket if it's empty
		if dietBucket.Stats().KeyN == 0 {
			dietBucket.Put([]byte("Group 1"), []byte("0"))
			dietBucket.Put([]byte("Group 2"), []byte("0"))
			dietBucket.Put([]byte("Group 3"), []byte("0"))
			dietBucket.Put([]byte("Group 4"), []byte("0"))
			dietBucket.Put([]byte("Group 5"), []byte("0"))
			dietBucket.Put([]byte("Group 6"), []byte("0"))
			dietBucket.Put([]byte("Group 7"), []byte("0"))
		}

		return nil
	})

	// Reads flags
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.StringVar(&mealToAdd, "a", mealToAdd, "comma separeted string of ingridients eaten")
	flag.StringVar(&mealToAdd, "add", mealToAdd, "comma separeted string of ingridients eaten")
	flag.StringVar(&newFoodItem, "n", newFoodItem, "name of the new food item and the groups it is part of (eg \"foodName 1234567\")")
	flag.StringVar(&newFoodItem, "new", newFoodItem, "name of the new food item and the groups it is part of (eg \"foodName 1234567\")")
	flag.BoolVar(&showHelp, "h", false, "dispaly help")
	flag.BoolVar(&debug, "d", false, "enables debug mode")
	flag.BoolVar(&debug, "debug", false, "enables debug mode")
}

func main() {
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
