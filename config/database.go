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
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "host="+os.Getenv("COMMENTAIRES_DATABASE_HOST")+" user="+os.Getenv("COMMENTAIRES_DATABASE_USER")+" dbname="+os.Getenv("COMMENTAIRES_DATABASE_NAME")+" sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	createTablesComment()
	createTablesAvis()
}

func createTablesComment() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS commentaires(id_commentaire int NOT NULL UNIQUE PRIMARY KEY GENERATED ALWAYS AS IDENTITY, Id_User int, Texte varchar(500), Date_Commentaire timestamp, id_post UUID)")

	if err != nil {
		log.Fatal(err)
	}
}

func createTablesAvis() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS avis(Id_Commentaire int REFERENCES commentaires(id_commentaire), Id_User int,Valeur  BOOLEAN, PRIMARY KEY(Id_Commentaire, Id_User))")

	if err != nil {
		log.Fatal(err)
	}
}

func Get_Db() *sql.DB {
	return db
}
