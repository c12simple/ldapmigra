package main

import(
	"fmt"
)

func main(){
	filePath = ""
	bootstrapJSON, err := GetContent("abc")

	if err != nil {
		fmt.Println("Could not read %s", filePath)
	}
}