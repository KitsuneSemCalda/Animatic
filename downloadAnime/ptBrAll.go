package downloadanime

import (
	"Animatic/utils"
	"fmt"
)

const BaseAnimeUrlPtBr string = "https://animefire.net"

func searchAnime(animeName string){
  currentPageUrl := fmt.Sprintf("%s/pesquisar/%s", BaseAnimeUrlPtBr, animeName)
  fmt.Println(currentPageUrl)
}

func SelectAnime(animeName string){
  animeName = utils.TreatingAnimeName(animeName)
  searchAnime(animeName)
}
