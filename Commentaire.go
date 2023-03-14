package main

/*    BDD
Name : commentaires
Type : Postgres
Port : 5432:5432
User : Comentaires
mdp  : Comentaires
*/

import (
	"bufio"
	"fmt"
	"os"
)

func ZoneSaisi(saisi string) string {
	return saisi
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Texte : ")
	scanner.Scan()
	Texte := scanner.Text()
	fmt.Println(ZoneSaisi(Texte))
}
