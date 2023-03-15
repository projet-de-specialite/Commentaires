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

	"main/api"
	"main/config"
)

func main() {
	config.DatabaseInit()

	err := http.ListenAndServe(":80", api.Handlers())

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
