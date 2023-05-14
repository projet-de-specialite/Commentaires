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

func NewComment(c *Commentaire) error { //localhost/api/comment/newComment
	if c == nil {
		log.Fatal(c)
	}

	c.Date_Commentaire = time.Now()

	requete := "INSERT INTO commentaires (Id_User, Texte, Date_Commentaire, Id_Post )VALUES('" + strconv.Itoa(c.Id_User) + "', '" + c.Texte + "','" + c.Date_Commentaire.Format("2006-01-02 15:04:05") + "', '" + c.Id_Post + "')"
	_, err := config.Get_Db().Exec(requete)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func FindCommentById(id int) *Commentaire { //non utilisé
	var commentaire Commentaire

	row := config.Get_Db().QueryRow("SELECT * FROM commentaires WHERE Id_Commentaire = $1;", id)
	err := row.Scan(&commentaire.Id_Commentaire, &commentaire.Id_User, &commentaire.Texte, &commentaire.Date_Commentaire, &commentaire.Id_Post)

	if err != nil {
		fmt.Println(err)
	}

	return &commentaire
}

func DeleteCommentFromIdComment(id_comment int) error { //localhost/api/comment/deleteComment/:Id_Commentaire
	var comment Commentaire

	row := config.Get_Db().QueryRow("SELECT Id_Commentaire FROM commentaires WHERE Id_Commentaire = $1", id_comment)
	err := row.Scan(&comment.Id_Commentaire)

	if err != nil {
		fmt.Println(err)
	} else {
		stmt, err := config.Get_Db().Prepare("DELETE FROM avis WHERE Id_Commentaire = " + strconv.Itoa(id_comment) + ";")

		if err != nil {
			fmt.Println(err)

		} else {
			stmt.Exec()

			stmt, err = config.Get_Db().Prepare("DELETE FROM commentaires WHERE Id_Commentaire = " + strconv.Itoa(id_comment) + ";")

			if err != nil {
				fmt.Println(err)
			} else {
				stmt.Exec()
			}
		}

	}
	return err
}

func PrintCommentByIdPost(id_post string) *[]Commentaire { //localhost/api/comment/printComment/:Id_Post
	var comments []Commentaire

	rows, err := config.Get_Db().Query("SELECT * FROM Commentaires WHERE Id_Post=$1", id_post)

	if err != nil {
		fmt.Println("This Post Does not exist !")
	} else {
		for rows.Next() {
			var comment Commentaire

			err = rows.Scan(&comment.Id_Commentaire, &comment.Id_User, &comment.Texte, &comment.Date_Commentaire, &comment.Id_Post)

			if err != nil {
				fmt.Println(err)
			} else {
				comments = append(comments, comment)
			}

		}
	}

	return &comments
}
