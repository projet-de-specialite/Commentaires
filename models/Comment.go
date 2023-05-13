// https://www.synbioz.com/blog/tech/developper-api-go
package models

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/projet-de-specialite/Commentaires/config"
)

type Commentaire struct {
	Id_Commentaire   int       `json:"Id_Commentaire"`
	Id_User          int       `json:"Id_User"`
	Texte            string    `json:"Texte"`
	Date_Commentaire time.Time `json:"Date_Commentaire"`
	Id_Post          string    `json:"Id_Post"`
}

func NewComment(c *Commentaire) {
	if c == nil {
		log.Fatal(c)
	}

	c.Date_Commentaire = time.Now()

	requete := "INSERT INTO commentaires (Id_User, Texte, Date_Commentaire, Id_Post )VALUES('" + strconv.Itoa(c.Id_User) + "', '" + c.Texte + "','" + c.Date_Commentaire.Format("2006-01-02 15:04:05") + "', '" + c.Id_Post + "')"
	_, err := config.Get_Db().Exec(requete)

	if err != nil {
		log.Fatal(err)
	}
}

func FindCommentById(id int) *Commentaire {
	var commentaire Commentaire

	row := config.Get_Db().QueryRow("SELECT * FROM commentaires WHERE Id_Commentaire = $1;", id)
	err := row.Scan(&commentaire.Id_Commentaire, &commentaire.Texte, &commentaire.Date_Commentaire, &commentaire.Id_Post)

	if err != nil {
		fmt.Println(err)
	}

	return &commentaire
}
