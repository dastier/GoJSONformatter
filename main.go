package main

import (
	"fmt"
)



func main() {
	rawJSON := getInputJSON()
	fmt.Println(stringToPrettyJSON(rawJSON))

}
