package main

import "fmt"

func main() {

	var num_8bit int8 = 100 // também chamado de byte (apelido)
	var num_16bit int16 = 10000
	var num_32bit int32 = 1000000000 //tabém chamado de rune (apelido)
	var num_64bit int64 = 1000000000000000000
	var num_AutoBit int = 10000
	var num_sem_sinal uint = 100000 // sem sinal (não aceita negativo)

	fmt.Println("Int 8 bit....................: ", num_8bit)
	fmt.Println("Int 16 bit...................: ", num_16bit)
	fmt.Println("Int 32 bit...................: ", num_32bit)
	fmt.Println("Int 64 bit...................: ", num_64bit)
	fmt.Println("Int auto bit.................: ", num_AutoBit)
	fmt.Println("Int bit só numero positivo...: ", num_sem_sinal)

	var float32 float32 = 12323232442232345467456356234534234534.122444
	var float64 float64 = 222233123444.24121212
	floatAuto := 1223.43 // pega de acordo com a arquitetura do processador x86/x64

	fmt.Println("Float 32 bits................: ", float32)
	fmt.Println("Float 64 bits................: ", float64)
	fmt.Println("Float auto bits..............: ", floatAuto)

	var texto string = "Nome da estring aqui"
	fmt.Println("Texto da String..............: ", texto)

}
