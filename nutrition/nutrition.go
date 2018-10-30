package nutrition

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/boltdb/bolt"
)

/*
The DB has two main buckets: "Food" and "Diet":

"Diet" contains 7 keys:
	"G1"..."G7"
The values represents the percentage that has been eaten for the specified group.

"Food" contains sub-buckets named after the food item they represent:
	"foodname" which has the same structure as "Diet", represents a food item.
*/

const DB = "../diet.db"

func Main(args map[string]interface{}) error {
	mealToAdd := args["mealToAdd"].(string)
	if mealToAdd != "" {
		// parseMeal
		addMeal("")
	}

	newFoodItem := args["newFoodItem"].(string)
	if newFoodItem != "" {
		parsedItem, err := parseFoodItem(newFoodItem)
		if err != nil {
			log.Panic(err)
		}
		newFood(parsedItem)
	}
	return nil
}

// Creates a new food item based on the map returned by parsing the user input
func newFood(item map[string]string) {
	// open db
	db, err := bolt.Open(DB, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create a sub-bucket for the new food item
	err = db.Update(func(tx *bolt.Tx) error {
		foodBucket := tx.Bucket([]byte("Food"))
		newItemBucket, err := foodBucket.CreateBucketIfNotExists([]byte(item["name"]))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		// populate sub-bucket
		for key, value := range item {
			if key == "name" {
				continue
			}
			err := newItemBucket.Put([]byte(key), []byte(value))
			if err != nil {
				return fmt.Errorf("put value: %s", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

// Checks and parse the user input and returns it in the form of a map
func parseFoodItem(newFoodItem string) (map[string]string, error) {
	ret := make(map[string]string)

	// check input format, should be "name gn:x"
	ok, err := regexp.MatchString("^\\w+(\\s[Gg]\\d:\\d+)+", newFoodItem)
	if err != nil {
		log.Panic(err)
	}
	if !ok {
		return nil, errors.New("wrong pattern for food item")
	}

	// parse food name
	tmp := strings.Split(newFoodItem, " ")
	ret["name"] = strings.ToLower(tmp[0])

	// parse groups percentages
	for i := 1; i < len(tmp); i++ {
		aux := strings.Split(tmp[i], ":")
		if n := aux[0][1] - 48; n > 7 {
			return nil, fmt.Errorf("group %d does not exist", n)
		}
		ret[strings.ToUpper(aux[0])] = aux[1]
	}
	return ret, nil
}

func addMeal(meal string) error {
	// open db
	db, err := bolt.Open(DB, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// search db for foodItem
	err = db.View(func(tx *bolt.Tx) error {
		foodBucket := tx.Bucket([]byte("Food"))
		itemBucket := foodBucket.Bucket([]byte(meal))
		if itemBucket != nil {
			// retrive values
			itemBucket.ForEach(func(k, v []byte) error {
				// foreach do:
				// update diet bucket
				return nil
			})
		} else {
			return fmt.Errorf("")
		}
		return nil
		// show balance
	})
	return nil
}
