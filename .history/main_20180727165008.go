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
	content, _ := unMarshalJSON(bootstrapJSON)
	fmt.Println("Dump content:")
	fmt.Printf("%v", content)
}

func unMarshalJSON(jsonstr []byte) (map[string]string, error) {
	var content interface{}
	e := json.Unmarshal(jsonstr, &content)
	if e != nil {
		fmt.Printf("Error during unmarshal content %v", e)
		return nil, e
	}
	return content, nil
}
