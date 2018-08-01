package main

import(
	"fmt"
)

func main(){
	string filePath = ""
	bootstrapJSON, err := GetContent("abc")

	if err != nil {
		fmt.Printf("Could not read %s", filePath)
	}


}