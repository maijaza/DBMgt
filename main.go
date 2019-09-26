package main

import (
	"fmt"

	"github.com/maijaza/DBMgt/dbmanager"
)

func main() {
	fmt.Println("start dbMgt package.")
	//dbmanager.Test()
	db := new(dbmanager.DBConfig)
	*db = dbmanager.GetConfig()
	fmt.Println("FileName " + db.Database)

}
