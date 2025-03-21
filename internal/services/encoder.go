package services

import (
	"encoding/base64"
	"os"
)

func EncodeAudioToBase64(filePath string) (string, error) {
	audioBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	base64str := base64.StdEncoding.EncodeToString(audioBytes)

	return base64str, nil
}

func EncodeStickerToBase64(filepath string) (string, error) {
	stickerBytes, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	base64str := base64.StdEncoding.EncodeToString(stickerBytes)

	return base64str, nil
}
