package main

/*    BDD
Name : commentaires
Type : Postgres
Port : 5432:5432
User : comentaires
mdp  : comentaires
*/

import (
	"bufio"
	"fmt"
	"os"

	"main/config"
	"main/models"
)

func main() {
	config.DatabaseInit()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Texte : ")
	scanner.Scan()
	/*Texte := scanner.Text()
	models.NewCommentaire(&c, Texte)*/

	c := *models.FindCommentaireById(2)
	fmt.Println(models.ToString(c))

}
