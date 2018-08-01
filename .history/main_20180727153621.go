package main

import(
	"fmt"
)

func main(){
	filePath := ""
	bootstrapJSON, err := GetContent("abc")

	if err != nil {
		fmt.Printf("Could not read %s", filePath)
	}

	
}