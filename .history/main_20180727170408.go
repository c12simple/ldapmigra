package main

import (
	"encoding/json"
	"fmt"

	"github.com/c12simple/ldapmigra/files"
)

func main() {
	filePath := "/home/tran/Documents/bootstrap.json.txt"
	bootstrapJSON, err := files.GetContent(filePath)

	if err != nil {
		fmt.Printf("Could not read %s %v", filePath, err)
	}
	fmt.Println("Content:")
	fmt.Printf("%s", bootstrapJSON)
	content, err := unMarshalJSON(bootstrapJSON)
	if err != nil {
		fmt.Printf("Error during unmarshal %v", err)
	}

	for 
}

func unMarshalJSON(jsonstr []byte) (map[string]interface{}, error) {
	var content interface{}
	e := json.Unmarshal(jsonstr, &content)
	if e != nil {
		fmt.Printf("Error during unmarshal content %v", e)
		return nil, e
	}
	return content.(map[string]interface{}), nil
}
