package api

import (
	"github.com/gin-gonic/gin"

	"main/models"
)

func Handlers() *gin.Engine {
	r := gin.Default()

	v1Commentaires := r.Group("api/v1/commentaire")
	{
		v1Commentaires.POST("", PostCommentaire)
		//v1Commentaires.GET("", GetUsers)
		//v1Commentaires.GET(":id", GetUser)
		//v1Commentaires.PUT(":id", EditUser)
		//v1Commentaires.DELETE(":id", DeleteUser)
	}

	return r
}

func PostCommentaire(c *gin.Context) {
	var json models.Commentaire
	c.Bind(&json)

	if json.Texte != "" {
		models.NewCommentaire(&json)
		c.JSON(201, gin.H{"succes": json})
	} else {
		c.JSON(422, gin.H{"error": "field are empty"})
	}
}
