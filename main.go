package main

import (
	"fmt"
)



func main() {
	rawJSON := getInputJSON()
	// the line just to test git actions!!
	fmt.Println(stringToPrettyJSON(rawJSON))

}
