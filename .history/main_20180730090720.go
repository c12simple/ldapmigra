package main

import (
	"encoding/json"
	"fmt"

	"github.com/c12simple/ldapmigra/files"
)

func main() {
	filePath := "/home/tran/Documents/bootstrap.json.txt"
	bootstrapJSON, err := files.GetContent(filePath)

}

func unMarshalJSON(jsonstr []byte) (map[string]interface{}, error) {
	var content interface{}
	e := json.Unmarshal(jsonstr, &content)
	if e != nil {
		fmt.Printf("Error during unmarshal content %v", e)
		return nil, e
	}
	return content.(map[string]interface{}), nil
}

func extractLdapConfigFromMaster(config interface{}) (interface{}, error) {
	mapConfig := config.(map[string]interface{})
	for key, val := range mapConfig {
		if key == "MASTER_INSTANCE_CONFIG" {
			return val, nil
		}
	}
	return nil, nil
}

func extractLdapConfigFromSlave(config interface{}) (interface{}, error) {
	mapConfig := config.(map[string]interface{})
	for key, val := range mapConfig {
		if key == "SLAVE_INSTANCE_CONFIG" {
			return val, nil
		}
	}
	return nil, nil
}
