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
	Texte            string    `json:"Texte"`
	Date_Commentaire time.Time `json:"Date_Commentaire"`
	Id_post          int       `json:"Id_Post"`
}

func NewComment(c *Commentaire) {
	if c == nil {
		log.Fatal(c)
	}

	c.Date_Commentaire = time.Now()
	c.Id_post = 0

	requete := "INSERT INTO commentaires (Texte, Date_Commentaire, Id_Post )VALUES('" + c.Texte + "','" + c.Date_Commentaire.Format("2006-01-02 15:04:05") + "'," + strconv.Itoa(c.Id_post) + ")"
	_, err := config.Get_Db().Exec(requete)

	if err != nil {
		log.Fatal(err)
	}
}

func FindCommentById(id int) *Commentaire {
	var commentaire Commentaire

	row := config.Get_Db().QueryRow("SELECT * FROM commentaires WHERE Id_Commentaire = $1;", id)
	err := row.Scan(&commentaire.Id_Commentaire, &commentaire.Texte, &commentaire.Date_Commentaire, &commentaire.Id_post)

	if err != nil {
		fmt.Println(err)
	}

	return &commentaire
}

func ToString(this Commentaire) string {
	Chaine := ""

	if this.Id_Commentaire != 0 {
		Chaine = "id : " + strconv.Itoa(this.Id_Commentaire) + "\nTexte : " + this.Texte + "\nDate : " + this.Date_Commentaire.Format("2006-01-02 15:04:05")
	}

	return Chaine
}
