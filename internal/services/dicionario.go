package services

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
)

var basePath, err = os.Getwd()

var templates = map[int][]string{
	0: {"Tudo certo amigão? Como tá?",
		"Opa juninho, tudo supimpa?",
		"Aqui tá tudo perfeito!",
		"Bom dia! Espero que seu dia seja incrível. 😊",
		"Seja a mudança que você quer ver no mundo! 🌍",
		"Que tal uma pausa para respirar fundo e relaxar? 😌",
		"Acredite em você! Grandes coisas estão por vir. 🚀",
		"Nunca é tarde para tentar algo novo. Vamos nessa! 💪",
		"Sorria! O mundo precisa de mais alegria. 😃",
		"A vida é feita de momentos, aproveite cada um deles. ⏳",
		"Tenha um excelente dia! Você é incrível. ✨",
		"Lembre-se: O esforço de hoje é o sucesso de amanhã! 🔥",
		"A gratidão transforma o pouco em suficiente. 🙏",
	},
	1: {
		filepath.Join(basePath, "resources", "audio", "sound1.ogg"),
		filepath.Join(basePath, "resources", "audio", "sound2.ogg"),
		filepath.Join(basePath, "resources", "audio", "sound3.ogg"),
		filepath.Join(basePath, "resources", "audio", "sound4.ogg"),
		filepath.Join(basePath, "resources", "audio", "sound5.ogg"),
	},
	2: {
		filepath.Join(basePath, "resources", "sticker", "sticker1.webp"),
		filepath.Join(basePath, "resources", "sticker", "sticker2.webp"),
		filepath.Join(basePath, "resources", "sticker", "sticker3.webp"),
		filepath.Join(basePath, "resources", "sticker", "sticker4.webp"),
	}}

func EscolherTemplate() (int, string) {
	if err != nil {
		fmt.Printf("Erro ao pegar template.")
		return 0, ""
	}
	tipo := rand.IntN(3)
	opcoes, existe := templates[tipo]

	if !existe {
		return -1, "Nenhuma opção de template disponível"
	}

	escolha := opcoes[rand.IntN(len(opcoes))]
	return tipo, escolha
}
