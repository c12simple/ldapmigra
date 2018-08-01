package pydioldap

import (
	"encoding/json"
	"fmt"

	"github.com/c12simple/ldapmigra/files"
)

// Config structure of ldap config in Pydio
type Config struct {
	config map[string]string
}

// GetConfig return config structure from bootstrap.json file of Pydio
func (c *Config) GetConfig(filePath string) error {
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
			//fmt.Printf("key %v", val)
			ldapConfig, err := extractLdapConfigFromMaster(val)
			if err != nil {
				fmt.Printf("hey error %v", err)
			}
			if ldapConfig == nil && err == nil {
				ldapConfig, err = extractLdapConfigFromSlave(val)
			}
			c.config = make(map[string]string)
			for 

			fmt.Printf("Ldap config Dump: \n%v", c.config)
		}
	}

	return fmt.Errorf("Could not get ldap config from %s", filePath)
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
