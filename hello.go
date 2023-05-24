package main

import "fmt"

func main() {
	var nome string
	var idade int

	fmt.Print("Qual o seu nome? ")
	fmt.Scan(&nome)

	fmt.Print("Qual a sua idade? ")
	fmt.Scan(&idade)

	fmt.Println("Olá", nome, "você tem", idade, "anos!")
}
