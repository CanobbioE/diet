package nutrition

import (
	"fmt"
	"strings"
	"testing"

	"github.com/boltdb/bolt"
)

func TestNewFood(t *testing.T) {
	input := []map[string]string{
		{"name": "all_groups",
			"G1": "10",
			"G2": "10",
			"G3": "10",
			"G4": "10",
			"G5": "10",
			"G6": "10",
			"G7": "30",
		},
		{"name": "one_group",
			"G1": "100",
		},
	}

	for _, i := range input {
		newFood(i)
		db, err := bolt.Open(DB, 0666, nil)
		if err != nil {
			t.Errorf("open db: %s", err)
		}

		err = db.View(func(tx *bolt.Tx) error {
			for g, v := range i {
				if g == "name" {
					continue
				}
				value := string(tx.Bucket([]byte("Food")).Bucket([]byte(i["name"])).Get([]byte(g)))
				if strings.Compare(value, v) != 0 {
					return fmt.Errorf("Inserted wrong value: %v (expected %v)", string(value), v)
				}
			}
			return nil
		})
		if err != nil {
			db.Close()
			t.Error(err)
		}

		db.Close()
	}
}

func TestParseFoodItem(t *testing.T) {
	input := []struct {
		in   string
		want map[string]string
		err  bool
	}{
		{"123", map[string]string{"name": "123"}, true},
		{"pizza g1:10", map[string]string{"name": "pizza", "G1": "10"}, false},
		{"stuff g8:99", map[string]string{}, true},
	}

	for _, i := range input {
		got, err := parseFoodItem(i.in)
		if err != nil && i.err {
			continue
		}
		if err != nil && !i.err {
			t.Error(err)
		}
		for k, v := range got {
			if v != i.want[k] {
				t.Errorf("got %v: %v (expected %v: %v)", k, v, k, i.want[k])
			}
		}
	}
}
