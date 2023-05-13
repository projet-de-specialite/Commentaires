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
		//v1Commentaires.GET("", GetUsers)
		//v1Commentaires.GET(":id", GetUser)
		//v1Commentaires.PUT(":id", EditUser)
		//v1Commentaires.DELETE(":id", DeleteUser)
	}

	return r
}

func PostComment(c *gin.Context) {
	var json models.Commentaire
	c.Bind(&json)

	if json.Texte != "" && json.Id_User != 0 && json.Id_Post != "" {
		models.NewComment(&json)
		c.JSON(201, gin.H{"succes": json})
	} else {
		c.JSON(422, gin.H{"error": "field are empty"})
	}
}

func PostAvis(c *gin.Context) {
	var json models.Avis
	var commentaire models.Commentaire
	c.Bind(&json)

	if json.Id_Commentaire != 0 && json.Id_User != 0 {
		row := config.Get_Db().QueryRow("SELECT Id_Commentaire FROM commentaires WHERE Id_Commentaire = $1", json.Id_Commentaire)
		void := row.Scan(&commentaire.Id_Commentaire)

		if void != nil {
			c.JSON(421, gin.H{"error": "Comment does not exist"})
		} else {
			models.NewAvis(&json)
			c.JSON(201, gin.H{"succes": json})
		}

	} else {
		c.JSON(421, gin.H{"error": "field are empty"})
	}
}

func GetNombreAvis(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("Id_Commentaire"))

	if err != nil {
		fmt.Println(err)
		c.JSON(421, gin.H{"error": "ID incorrect"})
	} else {
		avis := models.CountAvisByComment(ID)
		c.IndentedJSON(http.StatusOK, avis)
	}

}
