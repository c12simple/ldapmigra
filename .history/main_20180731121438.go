package main

import (
	"fmt"
	"github.com/c12simple/ldapmigra/pydioldap"
	
)

func main() {
	filePath := "/home/tran/Documents/bootstrap.json.txt"
	var ldapConfig pydioldap.Config
	err := ldapConfig.GetConfig(filePath)

	if err != nil {
		fmt.Println("Error")
	}	

	ldapConfig.
}
