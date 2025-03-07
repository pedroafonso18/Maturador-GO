package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pedroafonso18/Maturador-GO/internal/config"
	"github.com/pedroafonso18/Maturador-GO/internal/services"
)

func main() {
	fmt.Println("Iniciando Maturador GO...")

	// Try to load config, but don't fail if .env is missing
	err := config.Load()
	if err != nil {
		fmt.Printf("Aviso: %v\n", err)
		fmt.Println("Tentando continuar com variáveis de ambiente...")
	}

	// Check if resources directory exists
	if _, err := os.Stat("/app/resources"); os.IsNotExist(err) {
		fmt.Println("AVISO: Diretório de recursos não encontrado em /app/resources")
	}

	for {
		if services.ReturnTime() {
			fmt.Println("Iniciando ciclo de maturação...")
			services.Maturador()
			fmt.Println("Maturação concluída, aguardando próximo ciclo...")
		} else {
			fmt.Println("Fora do horário de funcionamento, aguardando próxima verificação...")
		}

		// Add reasonable sleep to prevent CPU overload
		time.Sleep(5 * time.Minute)
	}
}
