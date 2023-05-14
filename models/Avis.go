// https://www.synbioz.com/blog/tech/developper-api-go
package models

import (
	"fmt"
	"log"

	"strconv"

	"github.com/projet-de-specialite/Commentaires/config"
)

type Avis struct {
	Id_Commentaire int  `json:"Id_Commentaire"`
	Id_User        int  `json:"Id_User"`
	Valeur         bool `json:"Valeur"`
}

func NewAvis(a *Avis) {
	if a == nil {
		log.Fatal(a)
	}

	var avis Avis
	Commentaire := a.Id_Commentaire
	User := a.Id_User
	Valeur := a.Valeur

	row := config.Get_Db().QueryRow("SELECT * FROM avis WHERE Id_Commentaire = $1 AND Id_User = $2", Commentaire, User)
	void := row.Scan(&avis.Id_Commentaire, &avis.Id_User, &avis.Valeur)

	if void != nil {
		requete := "INSERT INTO avis (Id_Commentaire, Id_User, Valeur )VALUES('" + strconv.Itoa(Commentaire) + "', '" + strconv.Itoa(User) + "', '" + strconv.FormatBool(Valeur) + "')"
		_, err := config.Get_Db().Exec(requete)

		if err != nil {
			log.Fatal(err)
		}

	} else if avis.Valeur == Valeur {
		stmt, err := config.Get_Db().Prepare("DELETE FROM avis WHERE Id_Commentaire = " + strconv.Itoa(Commentaire) + " AND Id_User = " + strconv.Itoa(User) + ";")

		if err != nil {
			log.Fatal(err)
		}

		stmt.Exec()

	} else if avis.Valeur != Valeur {
		fmt.Println("Avis déjà posté")
	}

}

func CountAvisByComment(id_comment int) [2]int {
	var avis [2]int

	row := config.Get_Db().QueryRow("SELECT COUNT(*) FROM avis WHERE Id_Commentaire = $1 AND Valeur = true", id_comment)
	err := row.Scan(&avis[0])
	if err != nil {
		log.Fatal(err)
	}

	row = config.Get_Db().QueryRow("SELECT COUNT(*) FROM avis WHERE Id_Commentaire = $1 AND Valeur = false", id_comment)
	err = row.Scan(&avis[1])
	if err != nil {
		log.Fatal(err)
	}

	return avis
}

func DeleteAvisFromIdCommentaire(id_comment int) {
	var avis Avis
	row := config.Get_Db().QueryRow("SELECT Id_Commentaire FROM avis WHERE Id_Commentaire = $1", id_comment)
	err := row.Scan(&avis.Id_Commentaire)
	if err == nil {
		stmt, err := config.Get_Db().Prepare("DELETE FROM avis WHERE Id_Commentaire = " + strconv.Itoa(id_comment) + ";")

		if err != nil {
			log.Fatal(err)
		}
		stmt.Exec()
	}
}
