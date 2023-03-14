package config

import (
	"database/sql"
	"log"

	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DatabaseInit() {
	var err error

	db, err = sql.Open("postgres", "user=commentaires password=commentaires dbname=commentaires sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS commentaires(id_commentaire serial, Texte varchar(500), Date_Commentaire timestamp, pouce_rouge int, pouce_vert int, id_post int)")

	if err != nil {
		log.Fatal(err)
	}
}

func AjoutCommentaire(Chaine string) {
	requete := "INSERT INTO commentaires (Texte, Date_Commentaire, pouce_rouge, pouce_vert, id_post )VALUES('" + Chaine + "','" + time.Now().Format("2006-01-02 15:04:05") + " ', 0, 0, null)"
	_, err := db.Exec(requete)

	if err != nil {
		log.Fatal(err)
	}
}
