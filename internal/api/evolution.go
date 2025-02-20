package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pedroafonso18/Maturador-GO/internal/config"
)

type MessageRequest struct {
	Number string `json:"number"`
	Text   string `json:"text"`
}

type StickerRequest struct {
	Number  string `json:"number"`
	Sticker string `json:"sticker"`
}

func SendMessageEvo(num string, nome string, template string) error {
	fmt.Println("SendMessageEvo: starting execution")
	reqBody := MessageRequest{
		Number: num,
		Text:   template,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Printf("SendMessageEvo: error marshaling request body: %v\n", err)
		return fmt.Errorf("error marshaling request body: %v", err)
	}
	fmt.Println("SendMessageEvo: request body marshaled successfully")

	url := fmt.Sprintf("%s/message/sendText/%s", config.EVOURL, nome)
	fmt.Printf("SendMessageEvo: sending request to URL: %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("SendMessageEvo: error creating request: %v\n", err)
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("apikey", config.EVOTOKEN)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("SendMessageEvo: error sending request: %v\n", err)
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("SendMessageEvo: received response with status: %s\n", resp.Status)
	return nil
}

func SendStickerEvo(num string, nome string, sticker string) error {
	fmt.Println("SendStickerEvo: starting execution")
	reqBody := StickerRequest{
		Number:  num,
		Sticker: sticker,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Printf("SendStickerEvo: error marshaling request body: %v\n", err)
		return fmt.Errorf("error marshaling request body: %v", err)
	}
	fmt.Println("SendStickerEvo: request body marshaled successfully")

	url := fmt.Sprintf("%s/message/sendSticker/%s", config.EVOURL, nome)
	fmt.Printf("SendStickerEvo: sending request to URL: %s\n", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("SendStickerEvo: error creating request: %v\n", err)
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("apikey", config.EVOTOKEN)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("SendStickerEvo: error sending request: %v\n", err)
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("SendStickerEvo: received response with status: %s\n", resp.Status)
	return nil
}
