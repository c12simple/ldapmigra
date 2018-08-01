package main

import (
	"encoding/json"
	"fmt"

	"github.com/c12simple/ldapmigra/files"
)

func main() {
	filePath := "/home/tran/Documents/bootstrap.json.txt"
	bootstrapJSON, err := files.GetContent(filePath)

	if err != nil {
		fmt.Printf("Could not read %s %v", filePath, err)
	}

	content, err := unMarshalJSON(bootstrapJSON)
	if err != nil {
		fmt.Printf("Error during unmarshal %v", err)
	}

	for key, val := range content {
		if key == "core.auth" {
			//coreAuth := make(map[string]interface{})
			ldapConfig, err = ext(val.([]byte))
			if err != nil {

			}

		}
	}
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

func extractLdapConfigFromMaster(config interface{}) (map[string]string, error){
	mapConfig := config.(map[string]interface{})
	for key, val := range mapConfig {
		if key == "MASTER_INSTANCE_CONFIG" {
			return val.(map[string]string), nil
		}
	}
	return nil, nil
}

func extractLdapConfigFromSlave(config interface{}) (map[string]string, error){
	mapConfig := config.(map[string]interface{})
	for key, val := range mapConfig {
		if key == "SLAVE_INSTANCE_CONFIG" {
			return val.(map[string]string), nil
		}
	}
	return nil, nil
}