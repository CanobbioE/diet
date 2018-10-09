package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/CanobbioE/diet/nutrition"
	"github.com/boltdb/bolt"
)

const (
	version = "v0.0.1"
	DB      = "diet.db"
)

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

// init checks for the database existence/integrity and it reads the flags
func init() {
	// Reads flags
	flag.BoolVar(&showHelp, "h", false, "dispaly help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.StringVar(&newFoodItem, "n", newFoodItem, "name of a new food item followed by the groups it is part of (specifying the percentage)\n\t\tExample:\n\t\tdiet -n\"food g1:10 g2:10 g3:10 g4:10 g5:10 g7:50\"\n\t\tWhere \"food\" is the food's name and \"gn:x\" means the food has x percent of the n-th group nutritional value")
	flag.StringVar(&newFoodItem, "new", newFoodItem, "name of a new food item followed by the groups it is part of (specifying the percentage)\n\t\tExample:\n\t\tdiet -n\"food g1:10 g2:10 g3:10 g4:10 g5:10 g7:50\"\n\t\tWhere \"food\" is the food's name and \"gn:x\" means the food has x percent of the n-th group nutritional value")
	flag.StringVar(&mealToAdd, "a", mealToAdd, "comma separeted string of ingridients eaten")
	flag.StringVar(&mealToAdd, "add", mealToAdd, "comma separeted string of ingridients eaten")

	// Creates db if it doesn't exist
	db, err := bolt.Open(DB, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Creates the two buckets, "Food" and "Diet" if they don't exists.
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Food"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		dietBucket, err := tx.CreateBucketIfNotExists([]byte("Diet"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		// populates diet if it's empty
		if dietBucket.Stats().KeyN == 0 {
			for i := 1; i < 8; i++ {
				var buffer bytes.Buffer
				buffer.WriteString("Group " + strconv.Itoa(i))
				err := dietBucket.Put(buffer.Bytes(), []byte("0"))
				if err != nil {
					return fmt.Errorf("put value: %s", err)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()

	if showHelp == true {
		usage(appName, version)
	} else {
		args := make(map[string]interface{})
		args["mealToAdd"] = mealToAdd
		args["newFoodItem"] = newFoodItem

		if err := nutrition.Main(args); err != nil {
			log.Fatal(err)
		}
	}
}
