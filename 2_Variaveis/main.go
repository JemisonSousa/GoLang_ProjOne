package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variavel 2"
	fmt.Println("Variável 1: ", variavel1)
	fmt.Println("Variável 2: ", variavel2)

	var (
		variavel3 = "aaaaaaaa"
		variavel4 = "bbbbbbbb"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "varriavel5", "varriavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println("Contante: ", constante1)

	//invertendo variável
	variavel1, variavel2 = variavel5, variavel6
	fmt.Println("Variável 1 agora é: ", variavel1)
	fmt.Println("Variável 2 agora é: ", variavel2)

}
