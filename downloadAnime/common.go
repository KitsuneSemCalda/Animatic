package downloadanime

type Episode struct {
	Number string
	URL    string
}

type Anime struct {
	Name     string
	URL      string
	Episodes []Episode
}

type VideoResponse struct {
	Data []VideoData `json:"data"`
}

type VideoData struct {
	Src   string `json:"src"`
	Label string `json:"label"`
}
