package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func stringToPrettyJSON(s string) string {
	bytes := []byte(s)
	var p2 interface{}
	if err := json.Unmarshal(bytes, &p2); err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	m := p2.(map[string]interface{})
	json, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

func getInputJSON() string {
	var s []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			break
		}
		s = append(s, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return string(strings.Join(s, "\n"))
}

func main() {
	rawJSON := getInputJSON()
	fmt.Println(stringToPrettyJSON(rawJSON))

}
