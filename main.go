package main

import (
	"fmt"
	"os"

)

func loadcfg(configPath string){
  _, exists := os.Stat(configPath)

  if os.IsExist(exists){

  }else{
    fmt.Printf("O arquivo de configuração não existe em: %s\n", configPath)
  }
}

func main(){
  loadcfg("~/.local/share/animatics/config.cfg")
}
