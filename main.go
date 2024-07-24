package main

import (
	"fmt"
	"linha-de-comando/app"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		fmt.Println(erro)
		os.Exit(1)
	}
}
