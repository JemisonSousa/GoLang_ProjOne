package main

import (
	"fmt"
	auxiliar "modulo_inicial/Auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("cdcd_#_gmail.com")
	fmt.Println(erro)
}
