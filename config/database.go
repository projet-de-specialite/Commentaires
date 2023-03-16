package config

/*    BDD
Name : commentaires
Type : Postgres
Port : 5432:5432
User : comentaires
mdp  : comentaires
*/

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "user=commentaires password=commentaires dbname=commentaires sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	createTablesComment()
}

func createTablesComment() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS commentaires(id_commentaire int NOT NULL UNIQUE PRIMARY KEY GENERATED ALWAYS AS IDENTITY, Texte varchar(500), Date_Commentaire timestamp, id_post int)")

	if err != nil {
		log.Fatal(err)
	}
}

func Get_Db() *sql.DB {
	return db
}
