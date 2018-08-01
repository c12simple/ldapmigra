package main

import(
	"fmt"
	"io/ioutil"
)

func main(){
	filePath := "/home/tran/Documents/bootstrap.json.txt"
	bootstrapJSON, err := getContent(filePath)

	if err != nil {
		fmt.Printf("Could not read %s %v", filePath, err)
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

func unMarshalJSON(jsonstr []byte) (map[string]string, err){
	content := make(map[string]string)
	Unmar
}