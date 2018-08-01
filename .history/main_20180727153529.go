package main

import(
	"fmt"
)

func main(){
	bootstrapJSON, err := GetContent("abc")

	if err != nill {
		fmt.Println("Could not read %s")
	}
}