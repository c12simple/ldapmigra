package main

import(
	"fmt"
)

func main(){
	filePath := "Documents/bootstrap.json.txt "
	bootstrapJSON, err := GetContent("abc")

	if err != nil {
		fmt.Printf("Could not read %s", filePath)
	}
	fmt.Println("Content:")
	fmt.Printf("%s", bootstrapJSON)

}