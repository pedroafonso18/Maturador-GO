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
	delay2 := rand.IntN(10 - 3)
	for i := 0; i < len(connections); i++ {
		sender := connections[i]
		receiver := connections[(i+1)%len(connections)]

		if err := exchangeMessages(sender, receiver, delay, delay2); err != nil {
			fmt.Printf("error sending message: %s\n", err)
			continue
		}
	}
}

func exchangeMessages(sender, receiver database.Instance, delay int, delay2 int) error {

	templateTipo, templateData := EscolherTemplate()
	if templateTipo == 2 {
		encoded, err := EncodeStickerToBase64(templateData)
		if err != nil {
			return err
		}
		if err := sendSticker(sender, receiver, encoded); err != nil {
			return err
		}
	} else if templateTipo == 0 {
		if err := sendMessage(sender, receiver, templateData); err != nil {
			return err
		}
	}
	time.Sleep(time.Duration(delay) * time.Second)

	templateTipo, templateData = EscolherTemplate()
	if templateTipo == 2 {
		encoded, err := EncodeStickerToBase64(templateData)
		if err != nil {
			return err
		}
		if err := sendSticker(receiver, sender, encoded); err != nil {
			return err
		}
	} else if templateTipo == 0 {
		if err := sendMessage(receiver, sender, templateData); err != nil {
			return err
		}
	} else {
		encoded, err := EncodeAudioToBase64(templateData)
		if err != nil {
			return nil
		}
		if err := sendAudio(receiver, sender, encoded); err != nil {
			return err
		}
	}
	time.Sleep(time.Duration(delay2) * time.Minute)
	return nil
}

func sendMessage(from, to database.Instance, message string) error {
	if from.IsEvo {
		return api.SendMessageEvo(to.Numero, from.Name, message)
	}
	return api.SendMessageWuz(to.Numero, message, from.InstanceId)
}

func sendSticker(from, to database.Instance, base string) error {
	if from.IsEvo {
		return api.SendStickerEvo(to.Numero, from.Name, base)
	}
	return api.SendStickerWuz(to.Numero, base, from.InstanceId)
}

func sendAudio(from, to database.Instance, audio string) error {
	if from.IsEvo {
		return api.SendMessageEvo(to.Numero, from.Name, "Sem audio para EVO")
	}
	return api.SendAudioWuz(to.Numero, audio, from.InstanceId)
}
