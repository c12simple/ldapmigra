package main

import (
	"github.com/c12simple/ldapmigra/cellsldap"
	"fmt"
	"github.com/c12simple/ldapmigra/pydioldap"
	
)

func main() {
	filePath := "test/bootstrap.json.txt"
	var ldapConfig pydioldap.Config
	err := ldapConfig.GetConfig(filePath)

	if err != nil {
		fmt.Println("Error")
	}	

	// Convert to cellsldap

	var cellsLdap  cellsldap.Config
	cellsLdap.ConvertFromPydioConfig(ldapConfig.Config)
	fmt.Printf("OK %v", cellsLdap)
}
