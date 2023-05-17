package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var tentativa [][]int

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	for {
		var tentativas []int
		var chute int
		numero_aleatorio := random()

		for {
			fmt.Print("Digite um número entre 1 e 100: ")
			fmt.Scan(&chute)

			tentativas = append(tentativas, chute)

			acertou := numero_aleatorio == chute
			maior := numero_aleatorio < chute
			menor := numero_aleatorio > chute

			if chute < 0 || chute > 100 {
				fmt.Println("Digite um número entre 1 e 100.")
				continue
			}
			if acertou {
				fmt.Println("Parabéns, você acertou!")
				fmt.Printf("Você utilizou %d tentativas", len(tentativas))
				fmt.Println(".")
				break
			} else {
				if maior {
					fmt.Println("O número é menor que ", chute)
				} else if menor {
					fmt.Println("O número é maior que ", chute)
				}
			}
		}
		tentativa = append(tentativa, tentativas)

		var x string
		fmt.Print("Você deseja jogar novamente? (s/n): ")
		fmt.Scan(&x)

		if x != "s" {
			break
		}
	}
	fmt.Println("Fim de jogo.")
	for i, y := range tentativa {
		fmt.Printf("Você utilizou %d tentativas na jogada %d.\n", len(y), i+1)
	}
}
func random() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}
