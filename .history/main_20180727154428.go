package main

import(
	"fmt"
	
)

func main(){
	filePath := "/home/tran/Documents/bootstrap.json.txt"
	bootstrapJSON, err := Get GetContent("abc")

	if err != nil {
		fmt.Printf("Could not read %s", filePath)
	}
	fmt.Println("Content:")
	fmt.Printf("%s", bootstrapJSON)

}