package main

import (
	"Animatic/config"
	downloadanime "Animatic/downloadAnime"
	"Animatic/utils"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type model struct {
	animeName      string
}


func main() {
  loadedConfig, err := config.LoadConfig(utils.GetFolder())
  
  if (err != nil){
    fmt.Printf("Ocorreu um erro ao carregar as configurações: %v\n", err)
    os.Exit(1)
  }

  db, err := utils.InitializeDB()

  if (err != nil){
    fmt.Printf("Ocorreu um erro em abrir o banco de dados: %v\n", err)
  }

  defer db.Close()

  prompt := promptui.Prompt{
		Label: fmt.Sprintf("Digite o nome do anime a ser %s", "baixado"),
	}

	animeName, err := prompt.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
  
  if loadedConfig.DownloadAll(){
    if loadedConfig.PortugueseSearch() {
      downloadanime.SelectAnime(db, animeName)
    }
  }

	os.Exit(0)
}
