package downloadanime

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
  "github.com/manifoldco/promptui"
)

const BaseSiteURL string = "https://animefire.net"

func selectAnime(animes []Anime) int {
	animesName := make([]string, 0)
	for i := range animes {
		animesName = append(animesName, animes[i].Name)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "▶ {{ .Name | cyan }}",
		Inactive: "  {{ .Name | white }}",
		Selected: "▶ {{ .Name | cyan | underline }}",
	}

	prompt := promptui.Select{
		Label:     "Select the anime",
		Items:     animes,
		Templates: templates,
	}

	index, _, err := prompt.Run()

	if err != nil {
		log.Fatalf("Failed to select anime: %v\n", err)
		os.Exit(1)
	}

	return index
}

func searchAnime(animeName string) (string, error) {
	currentPageURL := fmt.Sprintf("%s/pesquisar/%s", BaseSiteURL, animeName)

	for {
		response, err := http.Get(currentPageURL)
		if err != nil {
			log.Fatalf("Failed to perform search request: %v\n", err)
			os.Exit(1)
		}

		defer response.Body.Close()

		doc, err := goquery.NewDocumentFromReader(response.Body)

		if err != nil {
			log.Fatalf("Failed to parse response: %v\n", err)
			os.Exit(1)
		}

		animes := make([]Anime, 0)
		doc.Find(".row.ml-1.mr-1 a").Each(func(i int, s *goquery.Selection) {
			anime := Anime{
				Name: strings.TrimSpace(s.Text()),
				Url:  s.AttrOr("href", ""),
			}
			animeName = strings.TrimSpace(s.Text())

			animes = append(animes, anime)
		})

		if len(animes) > 0 {
			index := selectAnime(animes)
			selectedAnime := animes[index]

			return selectedAnime.Url, nil
		}

		nextPage, exists := doc.Find(".pagination .next a").Attr("href")
		if !exists || nextPage == "" {
			log.Fatalln("No anime found with the given name")
			os.Exit(1)
		}

		currentPageURL = BaseSiteURL + nextPage
		if err != nil {
			log.Fatalf("Failed to add anime names to the database: %v", err)
		}
	}
}

func SearchAnimeSitePtBr(animeName string){
  animeURL, err := searchAnime(TreatingAnimeName(animeName))

	if err != nil {
		log.Fatalf("Failed to get anime episodes: %v", err)
		os.Exit(1)
	}

  fmt.Println(animeURL)
}
