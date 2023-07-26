package main

import (
	"Animatic/config"
  "Animatic/utils"
	"fmt"
	"os"

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


func main() {
  loadedConfig, err := config.LoadConfig(utils.GetFolder())
  
  if (err != nil){
    fmt.Printf("Ocorreu um erro: %v\n", err)
    os.Exit(1)
  }

  prompt := promptui.Prompt{
		Label: fmt.Sprintf("Digite o nome do anime a ser %s", action),
	}

	animeName, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
  
  if loadedConfig.DownloadAll(){
    if loadedConfig.PortugueseSearch() {
      fmt.Println(animeName)
    }
  }

	os.Exit(0)
}
