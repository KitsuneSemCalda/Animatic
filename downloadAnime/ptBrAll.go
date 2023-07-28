package downloadanime

import (
	"Animatic/utils"
	"database/sql"
  "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	
  "github.com/PuerkitoBio/goquery"
	"github.com/cavaliergopher/grab/v3"
	"github.com/manifoldco/promptui"
)

const BaseAnimeUrlPtBr string = "https://animefire.net"

func DownloadVideo(db *sql.DB, destPath string, url string, animeName string, episode string) error {
	episode = utils.EpisodeFormatter(episode)
	episodeFilename := fmt.Sprintf("S01E%s.mp4", episode)

	err := utils.AddAnimeToTB(db, animeName, episode, destPath)
	if err != nil {
		fmt.Printf("Error input animeInfo: %v\n", err)
	}
  
  destPath = filepath.Join(destPath, episodeFilename)
  fmt.Printf("Download the anime %s Episode %s\n", animeName, episode)

	client := grab.NewClient()

	req, _ := grab.NewRequest(destPath, url)
	resp := client.Do(req)

	// Wait for the download to complete
	<-resp.Done

	if err := resp.Err(); err != nil {
		fmt.Printf("Error downloading episode: %v\n", err)
		return err
	}

		fmt.Printf("Episode %s of anime %s was downloaded to %s\n", episode, animeName, destPath)
	return nil
}

func extractVideoUrl(url string) (string, error){
  response, err := http.Get(url)

  if err != nil {
    return "", err
  }

  doc, _ := goquery.NewDocumentFromReader(response.Body)

	videoElements := doc.Find("video")

	if videoElements.Length() > 0 {
		oldDataVideo, _ := videoElements.Attr("data-video-src")
		return oldDataVideo, nil
	} else {
		videoElements = doc.Find("div")
		if videoElements.Length() > 0 {
			oldDataVideo, _ := videoElements.Attr("data-video-src")
			return oldDataVideo, nil
		}
	}

	return "", nil
}

func extractActualVideoURL(videoSrc string) (string, error) {
	response, err := http.Get(videoSrc)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status: %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var videoResponse VideoResponse
	err = json.Unmarshal(body, &videoResponse)
	if err != nil {
		return "", err
	}

	if len(videoResponse.Data) == 0 {
		return "", fmt.Errorf("no video data found")
	}

	return videoResponse.Data[0].Src, nil
}

// Modify the `downloadAll` function to use the Plex naming convention
func downloadAll(db *sql.DB, destPath string, anime Anime, epList []Episode) {
	animeName := utils.TreatingAnimeName(anime.Name)
	animeFolder := filepath.Join(destPath, animeName)

	for i := range epList {
		episodeNumber := utils.EpisodeFormatter(epList[i].Number)
		episodePath := filepath.Join(animeFolder, fmt.Sprintf("S01E%s.mp4", episodeNumber))

		animeURL := epList[i].URL
		videoURL, err := extractVideoUrl(animeURL)
		if err != nil {
			log.Fatalf("Failed to extract video URL: %v", err)
		}

		videoURL, err = extractActualVideoURL(videoURL)
		if err != nil {
			log.Fatal("Failed to extract the api")
		}

		DownloadVideo(db, episodePath, videoURL, anime.Name, epList[i].Number)

		if err != nil {
			log.Fatal("Failed to download episode")
		}
	}
}

func getAnimeEpisodes(animeURL string) ([]Episode, error) {
	resp, err := http.Get(animeURL)

	if err != nil {
		log.Fatalf("Failed to get anime details: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse anime details: %v\n", err)
		os.Exit(1)
	}

	episodeContainer := doc.Find("a.lEp.epT.divNumEp.smallbox.px-2.mx-1.text-left.d-flex")

	episodes := make([]Episode, 0)

	episodeContainer.Each(func(i int, s *goquery.Selection) {
		episodeNum := s.Text()
		episodeURL, _ := s.Attr("href")

		episode := Episode{
			Number: episodeNum,
			URL:    episodeURL,
		}
		episodes = append(episodes, episode)
	})
	return episodes, nil
}

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

// Add more error handling in the `SelectAnime` function
func SelectAnime(db *sql.DB, animeName string) {
	animeName = utils.TreatingAnimeName(animeName)
	animeURL, animeSelectedName, err := searchAnime(animeName)

	if err != nil {
		log.Fatalf("Failed to locate anime: %v\n", err)
		os.Exit(1)
	}

	destPath := filepath.Join(utils.GetFolder(), "Anime")

	epList, err := getAnimeEpisodes(animeURL)

	if err != nil {
		log.Fatalf("Failed to get anime episodes: %v\n", err)
	}

	downloadAll(db, destPath, Anime{Name: animeSelectedName, URL: animeURL}, epList)
}

func searchAnime(animeName string) (string, string ,error){
  currentPageUrl := fmt.Sprintf("%s/pesquisar/%s", BaseAnimeUrlPtBr, animeName)

  for {
		response, err := http.Get(currentPageUrl)
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
				URL:  s.AttrOr("href", ""),
			}
			animeName = strings.TrimSpace(s.Text())

			animes = append(animes, anime)
		})

		if len(animes) > 0 {
			index := selectAnime(animes)
			selectedAnime := animes[index]

			return selectedAnime.URL, selectedAnime.Name, nil
		}

		nextPage, exists := doc.Find(".pagination .next a").Attr("href")
		if !exists || nextPage == "" {
			log.Fatalln("No anime found with the given name")
			os.Exit(1)
		}

		currentPageUrl = BaseAnimeUrlPtBr + nextPage
		if err != nil {
			log.Fatalf("Failed to add anime names to the database: %v", err)
		}
	}
}

