package nutrition

import "fmt"

func Main(args map[string]interface{}) {
	for _, i := range args {
		fmt.Println(i)
	}
}
