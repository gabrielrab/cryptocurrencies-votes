package seeder

import (
	"bufio"
	"cryptocurrencies-votes/database"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main(){
	database.DatabaseConnection()
	defer database.CloseConnection()


	db := database.GetDatabase()

	_, b, _, _ := runtime.Caller(0)
	seedFile := filepath.Dir(b) + "/seed.sql"


	transaction := db.Begin()
	file, err := os.Open(seedFile)

	if err != nil {
		log.Fatal(err)
	}

	content := bufio.NewScanner(file)
	for content.Scan() {
		err:= transaction.Exec(content.Text())
		if err.Error != nil {
			transaction.Rollback()
			break
		}

		if err := content.Err(); err != nil {
			if err != io.EOF {
				transaction.Rollback()
				break
			}
		}
	}

	transaction.Commit()

}