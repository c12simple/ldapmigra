package main

import(
	"fmt"

	"git"
)

func main(){
	filePath := "/home/tran/Documents/bootstrap.json.txt"
	bootstrapJSON, err := GetContent("abc")

	if err != nil {
		fmt.Printf("Could not read %s", filePath)
	}
	fmt.Println("Content:")
	fmt.Printf("%s", bootstrapJSON)

}