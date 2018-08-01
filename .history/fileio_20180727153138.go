package main

import (
	"io/ioutil"
)

// GetContent get text string in bootstrap.
func GetContent(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

}

