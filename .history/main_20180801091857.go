package main

import (
	"fmt"

	"github.com/c12simple/ldapmigra/cellsldap"
	"github.com/c12simple/ldapmigra/pydioldap"
)

func main() {
	filePath := "test/bootstrap.json"
	var ldapConfig pydioldap.Config
	err := ldapConfig.GetConfig(filePath)

	if err != nil {
		fmt.Printf("Error %v", err)
	}

	// Convert to cellsldap

	var cellsLdap cellsldap.Config
	fmt.Printf("cellsLdap.ConvertFromPydioConfig(ldapConfig.Config)	
}
