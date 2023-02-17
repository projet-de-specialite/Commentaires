package main

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
