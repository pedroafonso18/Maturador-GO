package main

import (
	"fmt"

	"github.com/pedroafonso18/Maturador-GO/internal/config"
	"github.com/pedroafonso18/Maturador-GO/internal/services"
)

func main() {
	for 1 > 0 {
		config.Load()
		if services.ReturnTime() {
			services.Maturador()
		} else {
			fmt.Println("Fora do horário de funcionamento, espere e tente novamente.")
		}
	}
}
