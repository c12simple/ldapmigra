package files

import "io/ioutil"
// GetContent get text string in bootstrap.json file
func getContent(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}