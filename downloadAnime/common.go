package downloadanime

import (
  "regexp"
  "strings"
)

type Episode struct {
  Number string
  URL string
}

type VideoData struct {
  Src string `json:"src"`
  Label string `json:"label"`
}

type VideoResponse struct {
  Data []VideoData `json:"data"`
}

type Anime struct {
  Name string
  Url string
  Episodes []Episode
}

func DatabaseFormatter(str string) string{
  regex := regexp.MustCompile(`\s*\([^)]*\)|\bn/a\b|\s+\d+(\.\d+)?$`)
	result := regex.ReplaceAllString(str, "")
	result = strings.TrimSpace(result)
	result = strings.ToLower(result)
	return result
}

func TreatingAnimeName(animeName string) string {
	loweredName := strings.ToLower(animeName)
	spacelessName := strings.ReplaceAll(loweredName, " ", "-")
	return spacelessName
}
