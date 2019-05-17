package main

import (
	"fmt"
	"math/rand"
)

func main() {

	gerarCpf()

	var sexoEscoliho int
	fmt.Println("Homem ou mulher ? \n homem = 1\n mulher = 2")
	fmt.Scan(&sexoEscoliho)

	homem, mulher := genereteName()

	if sexoEscoliho == 1 {
		fmt.Println(homem)

	} else if sexoEscoliho == 2 {
		fmt.Println(mulher)
	}

}

func genereteName() (homem, mulher string) {

	firstNameMan := [...]string{
		"Eduardo",
		"Cayo",
		"Gustavo",
		"Igor",
		"Matheus",
		"Mateus",
	}
	firstNamewoman := [...]string{
		"Lígia",
		"Ellen",
		"karol",
		"Joaquina",
		"Fernanda",
		"Ana paula",
	}
	lastName := [...]string{
		"Silva",
		"Santos",
		"Oliveira",
		"Souza",
		"Lima",
		"Pereira",
		"Ferreira",
		"Costa",
		"Rodrigues",
	}
	completeNameMan := (firstNameMan[rand.Intn(len(firstNameMan))] + " " + lastName[rand.Intn(len(lastName))])
	completeNamewoman := (firstNamewoman[rand.Intn(len(firstNamewoman))] + " " + lastName[rand.Intn(len(lastName))])

	return completeNameMan, completeNamewoman
}

func gerarCpf() {

	var (
		cpf            []int
		primeiraSoma   int
		primeiroDigito int
		segundaSoma    int
		segundoDigito  int
	)

	//criando os 9 primeiros digitos do cpf
	for cont1 := 0; cont1 < 9; cont1++ {
		aleatorio := rand.Intn(10)
		cpf = append(cpf, aleatorio)
	}

	// Calculo para validar os dois ultimos digitos do CPF
	// Calculo retirado de https://www.geradorcpf.com/algoritmo_do_cpf.htm
	primeiraSoma = (cpf[0] * 10) + (cpf[1] * 9) + (cpf[2] * 8) + (cpf[3] * 7) + (cpf[4] * 6) + (cpf[5] * 5) + (cpf[6] * 4) + (cpf[7] * 3) + (cpf[8] * 2)

	if (primeiraSoma % 11) < 2 {
		primeiroDigito = 0
	} else {
		primeiroDigito = 11 - (primeiraSoma % 11)
	}
	cpf = append(cpf, primeiroDigito)

	segundaSoma = (cpf[0] * 11) + (cpf[1] * 10) + (cpf[2] * 9) + (cpf[3] * 8) + (cpf[4] * 7) + (cpf[5] * 6) + (cpf[6] * 5) + (cpf[7] * 4) + (cpf[8] * 3) + (cpf[9] * 2)

	if (segundaSoma % 11) < 2 {
		segundoDigito = 0
	} else {
		segundoDigito = 11 - (segundaSoma % 11)
	}

	cpf = append(cpf, segundoDigito)

	fmt.Println(primeiraSoma % 11)
	fmt.Println(segundoDigito)
	fmt.Println(cpf)

}
