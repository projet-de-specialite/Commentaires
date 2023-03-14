package config

import (
	"database/sql"
	"log"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "user=commentaire dbname=commentaire")

	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS commentaire(id_commentaire serial, Texte varchar(500), pouce_rouge int, pouce_vert int, id_post int)")

	if err != nil {
		log.Fatal(err)
	}
}
