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
	base "main/config"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Texte : ")
	scanner.Scan()
	Texte := scanner.Text()

	base.DatabaseInit()
	base.AjoutCommentaire(Texte)

}
