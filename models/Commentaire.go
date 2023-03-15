package models

import (
	"log"
	"strconv"
	"time"

	"main/config"
)

type Commentaire struct {
	id_commentaire   int
	Texte            string
	Date_Commentaire time.Time
	id_post          int
}

func NewCommentaire(c *Commentaire, Chaine string) {
	if c == nil {
		log.Fatal(c)
	}

	c.Date_Commentaire = time.Now()
	c.Texte = Chaine
	c.id_post = 0

	requete := "INSERT INTO commentaires (Texte, Date_Commentaire, Id_Post )VALUES('" + c.Texte + "','" + c.Date_Commentaire.Format("2006-01-02 15:04:05") + "'," + strconv.Itoa(c.id_post) + ")"
	_, err := config.Get_Db().Exec(requete)

	if err != nil {
		log.Fatal(err)
	}
}

func FindCommentaireById(id int) *Commentaire {
	var commentaire Commentaire

	row := config.Get_Db().QueryRow("SELECT * FROM commentaires WHERE id_commentaire = $1;", id)
	err := row.Scan(&commentaire.id_commentaire, &commentaire.Texte, &commentaire.Date_Commentaire, &commentaire.id_post)

	if err != nil {
		log.Fatal(err)
	}

	return &commentaire
}

func ToString(this Commentaire) string {

	return "id : " + strconv.Itoa(this.id_commentaire) + "\nTexte : " + this.Texte + "\nDate : " + this.Date_Commentaire.Format("2006-01-02 15:04:05")
}
