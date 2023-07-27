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
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/cavaliergopher/grab/v3"
	"github.com/manifoldco/promptui"
)

var mw sync.WaitGroup

const BaseAnimeUrlPtBr string = "https://animefire.net"

func DownloadVideo(db *sql.DB, destPath string, url string, animeName string, episode string) {
	err := utils.AddAnimeToTB(db, animeName, episode, destPath)
	if err != nil {
		fmt.Printf("Error input animeInfo: %v\n", err)
		return
	}

	client := grab.NewClient()

	req, err := grab.NewRequest(destPath, url)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Create a channel to receive updates on download progress
	progressChan := make(chan *grab.Response)

	// Perform the download in a separate goroutine
	resp := client.Do(req)

	// Create a WaitGroup to wait for the download to complete
	var wg sync.WaitGroup
	wg.Add(1)

	// Start a goroutine to monitor the download progress
	go func() {
		defer wg.Done()

		// Loop to receive progress updates until the download is complete
		for !resp.IsComplete() {
			select {
			case <-progressChan:
				// Update progress if needed
			}
		}
	}()

	// Wait for the download to complete
	wg.Wait()

	// Check for any download errors
	if resp.Err() != nil {
		fmt.Printf("Error downloading: %v\n", resp.Err())
		return
	}

	// The file is downloaded successfully, so you can proceed with further actions
	fmt.Printf("Episode %s of anime %s was downloaded to %s\n", episode, animeName, destPath)
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


func downloadAll(db *sql.DB,destPath string, anime Anime ,epList []Episode){
  for i := range epList{
    episodePath := filepath.Join(destPath, epList[i].Number + ".mp4")
    animeUrl := epList[i].URL
    videoUrl, err := extractVideoUrl(animeUrl)
    if err != nil{
      log.Fatalf("Failed to extract video URL: %v", err)
    }
    
    videoUrl, err = extractActualVideoURL(videoUrl)
    
    if err != nil {
      log.Fatal("Failed to extract the api")
    }

    go DownloadVideo(db, episodePath, videoUrl, anime.Name, epList[i].Number)

    if err != nil{
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

func SelectAnime(db *sql.DB,animeName string){
  animeName = utils.TreatingAnimeName(animeName)
  animeUrl, animeSelectedName , err := searchAnime(animeName)
  
  SelectedAnime := Anime{Name: animeSelectedName, URL: animeUrl}

  if err != nil{
    log.Fatalf("Failed to locale anime: %v\n", err)
    os.Exit(1)
  }
  
  treatedName := utils.DatabaseFormatter(SelectedAnime.Name)
  destPath := filepath.Join(utils.GetFolder(), "anime", treatedName)
  fmt.Println(destPath)

  epList , err := getAnimeEpisodes(animeUrl)

  if err != nil{
    log.Fatalf("Failed to get animes episodes: %v\n", err)
  }
  
  downloadAll(db, destPath, SelectedAnime ,epList)
}
