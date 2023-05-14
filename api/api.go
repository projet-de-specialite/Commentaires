package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/projet-de-specialite/Commentaires/config"
	"github.com/projet-de-specialite/Commentaires/models"
)

func Handlers() *gin.Engine {
	r := gin.Default()

	v1Commentaires := r.Group("api/v1/comment")
	{
		v1Commentaires.POST("/newComment", PostComment)
		v1Commentaires.POST("/newAvis", PostAvis)
		v1Commentaires.GET("/nombreAvis/:Id_Commentaire", GetNombreAvis)
		v1Commentaires.GET("/printComment/:Id_Post", PrintCommentByPost)
		v1Commentaires.DELETE("/deleteComment/:Id_Commentaire", DeleteComment)
	}

	return r
}

func PostComment(c *gin.Context) { // Créer un commentaire
	var json models.Commentaire
	c.Bind(&json)

	if json.Texte != "" && json.Id_User != 0 && json.Id_Post != "" {
		err := models.NewComment(&json)
		if err == nil {
			c.JSON(201, gin.H{"succes": json})
		} else {
			c.JSON(422, gin.H{"errot": err})
		}

	} else {
		c.JSON(422, gin.H{"error": "field are empty"})
	}
}

func PostAvis(c *gin.Context) { //Post un avis sur un commentaire via un boolean pouce vert = true, pouce rouge = false
	var json models.Avis
	var commentaire models.Commentaire
	c.Bind(&json)

	if json.Id_Commentaire != 0 && json.Id_User != 0 {
		row := config.Get_Db().QueryRow("SELECT Id_Commentaire FROM commentaires WHERE Id_Commentaire = $1", json.Id_Commentaire)
		void := row.Scan(&commentaire.Id_Commentaire)

		if void != nil {
			c.JSON(421, gin.H{"error": "Comment does not exist"})
		} else {
			resultat := models.NewAvis(&json)
			if resultat == "" {
				c.JSON(201, gin.H{"succes": json})
			} else {
				c.JSON(201, gin.H{"Failed": resultat})
			}

		}

	} else {
		c.JSON(421, gin.H{"error": "field are empty"})
	}
}

func GetNombreAvis(c *gin.Context) { // Donne le nombre pouce rouge/vert du commentaire donner par son ID
	ID, err := strconv.Atoi(c.Param("Id_Commentaire"))

	if err != nil {
		fmt.Println(err)
		c.JSON(421, gin.H{"error": "ID incorrect"})
	} else {
		avis := models.CountAvisByComment(ID)
		c.IndentedJSON(http.StatusOK, avis)
	}

}

func DeleteComment(c *gin.Context) { //Supprime un commentaire et tous les avis le concernant
	ID, err := strconv.Atoi(c.Param("Id_Commentaire"))

	if err != nil {
		fmt.Println(err)
		c.JSON(421, gin.H{"error": "ID incorrect"})
	} else {
		err := models.DeleteCommentFromIdComment(ID)
		if err != nil {
			c.JSON(201, gin.H{"Failed": "This Comment does not exist !"})
		} else {
			c.JSON(201, gin.H{"succes": "Row deleted"})
		}

	}
}

func PrintCommentByPost(c *gin.Context) { //Affiche tous les commentaires d'un post
	comments := models.PrintCommentByIdPost(c.Param("Id_Post"))
	c.IndentedJSON(http.StatusOK, comments)

}
