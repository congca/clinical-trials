package main

import "fmt"
import "github.com/fjukstad/clinical"

func main() {
	studies := clinical.Search("breast cancer",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"")

	for _, study := range studies {
		fmt.Println(study)
	}
}
