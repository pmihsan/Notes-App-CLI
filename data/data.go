package data

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error 

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS notes (
			"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"word"  TEXT,
			"definition" TEXT,
			"category" TEXT
		);`

		statement, err := db.Prepare(createTableSQL)

		if err != nil {
			log.Fatal(err.Error())
		}

		statement.Exec()

		log.Println("Notes Table Created")
}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO notes (word, definition, category)
		VALUES (?, ?, ?)`

	statement, err := db.Prepare(insertNoteSQL)

	if err != nil {
		log.Fatalln(err)
	}
	_, err = statement.Exec(word, definition, category)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Note Inserted Successfully")
}

func DisplayAllNotes(){
	row, err := db.Query("SELECT * FROM notes ORDER BY word")

	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	for row.Next() {
		var id int
		var word string
		var def string
		var categ string

		row.Scan(&id, &word, &def, &categ)
		log.Println("[", categ, "] ", word, "-", def)
	}
}