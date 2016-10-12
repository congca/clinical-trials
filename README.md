# clinical-trials
Search for clinical trials from Health Canada.

# example 
Search for all breast cancer studies: 

```
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
```
