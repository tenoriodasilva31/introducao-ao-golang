package main

import "fmt"

const ebulicaoK = 373

func main() {

	var tempK = ebulicaoK
	var tempC = (tempK - 273)

	fmt.Println("Ponto de ebulição da água em Kº é:", tempK)
	fmt.Println("Ponto de ebulição da água em Cº é:", tempC)
}
