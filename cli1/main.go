package main

import (
	"encoding/json"
	"fmt"
	"learn-go-modules/lib"
)

func main() {
	person := lib.Person{"remew", 22}
	json, err := json.MarshalIndent(person, "", lib.JsonIndent)
	if err != nil {
		fmt.Println("JSON Error:", err)
		return
	}
	fmt.Println(string(json))
}
