package main

import (
	"Animatic/config"
	"fmt"
	"os"
)

func getFolder() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Erro ao obter o diretório do usuário:", err)
		return ""
	}

	return fmt.Sprintf("%s/.local/Animatic", userHomeDir)
}

func main() {
  loadedConfig, _ := config.LoadConfig(getFolder())
}
