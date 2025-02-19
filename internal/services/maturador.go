package services

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/pedroafonso18/Maturador-GO/internal/api"
	"github.com/pedroafonso18/Maturador-GO/internal/database"
)

func Maturador() {
	connections, err := database.FetchConnections()
	if err != nil {
		fmt.Printf("error getting connections: %s\n", err)
		return
	}

	if len(connections) < 2 {
		fmt.Print("Apenas uma conexão, sem possibilidade de maturação.")
		return
	}

	processConnections(connections)
}

func processConnections(connections []database.Instance) {
	delay := rand.IntN(30 - 8)
	for i := 0; i < len(connections); i++ {
		sender := connections[i]
		receiver := connections[(i+1)%len(connections)]

		if err := exchangeMessages(sender, receiver, delay); err != nil {
			fmt.Printf("error sending message: %s\n", err)
			continue
		}
	}
}

func exchangeMessages(sender, receiver database.Instance, delay int) error {
	templateTipo, templateData := EscolherTemplate()
	if templateTipo != 0 {
		return nil
	}

	if err := sendMessage(sender, receiver, templateData); err != nil {
		return err
	}
	time.Sleep(time.Duration(delay) * time.Second)

	if err := sendMessage(receiver, sender, templateData); err != nil {
		return err
	}
	time.Sleep(time.Duration(delay) * time.Second)

	return nil
}

func sendMessage(from, to database.Instance, message string) error {
	if from.IsEvo {
		return api.SendMessageEvo(to.Numero, from.Name, message)
	}
	return api.SendMessageWuz(to.Numero, message, from.InstanceId)
}
