package main

/*    BDD
Name : commentaires
Type : Postgres
Port : 5432:5432
User : comentaires
mdp  : comentaires
*/

import (
	"log"
	"net/http"
	"os"

	"github.com/projet-de-specialite/Commentaires/api"
	"github.com/projet-de-specialite/Commentaires/config"
)

func main() {
	config.DatabaseInit()

	err := http.ListenAndServe(os.Getenv("COMMENTAIRES_LISTEN"), api.Handlers())

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
