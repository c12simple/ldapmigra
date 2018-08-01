package main

import(
	"fmt"
	"io/ioutil"
)

func main(){
	filePath := "/home/tran/Documents/bootstrap.json.txt"
	bootstrapJSON, err := getContent("abc")

	if err != nil {
		fmt.Printf("Could not read %s ", filePath)
	}
	fmt.Println("Content:")
	fmt.Printf("%s", bootstrapJSON)

}

// GetContent get text string in bootstrap.json file
func getContent(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}