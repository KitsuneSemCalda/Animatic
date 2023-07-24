package main

import (
	"Animatic/config"
	"fmt"
	"os"
  "strings"

  "github.com/manifoldco/promptui"
)

type model struct {
	selectedOption optionMsg
	animeName      string
}

type optionMsg string

var action string

const (
	optionDownload optionMsg = "Baixar Anime"
	optionWatch    optionMsg = "Assistir Anime"
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
  config.LoadConfig(getFolder())
  p := promptui.Select{
		Label: "Escolha uma opção:",
		Items: []string{string(optionDownload), string(optionWatch)},
	}

	_, option, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	m := model{
		selectedOption: optionMsg(option),
	}
  
  if (m.selectedOption == optionDownload){
	  action = "baixado"
  }else{
    action = "assistido"
  }

  prompt := promptui.Prompt{
		Label: fmt.Sprintf("Digite o nome do anime a ser %s", action),
	}
	animeName, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	m.animeName = strings.TrimSpace(animeName)
  fmt.Println(m.animeName)
	os.Exit(0)
}
