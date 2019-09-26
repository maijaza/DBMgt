package main

import (
	"fmt"
	"os"

	"github.com/maijaza/DBMgt/dbmanager"
)

func main() {
	fmt.Println("start dbMgt package.")
	db := new(dbmanager.DBConfig)

	if len(os.Args) > 1 {
		db.FileName = os.Args[1]
	} else {
		db.FileName = "db"
	}
	*db = dbmanager.GetConfig(db.FileName)
	fmt.Println("FileName " + db.Database)

}
