package structure

import "errors"

type Anime struct {
	Name     string
	Url      string
	Episodes []Episode
}

func NewAnime(name string, url string) (*Anime, error) {
	if name == "" {
		return nil, errors.New("Anime name cannot be empty")
	}

	if url == "" {
		return nil, errors.New("Anime URL cannot be empty")
	}

	return &Anime{
		Name:     name,
		Url:      url,
		Episodes: make([]Episode, 0),
	}, nil
}

func (a *Anime) AddEpisode(episode Episode) error {
	if !IsValidEpisode(episode) {
		return errors.New("invalid episode")
	}

	a.Episodes = append(a.Episodes, episode)
	return nil
}

func (a *Anime) GetAnimeName() (*string, error) {
	if a == nil {
		return nil, errors.New("Anime is nil")
	}

	return &a.Name, nil
}

func (a *Anime) GetAnimeUrl() (*string, error) {
	if a == nil {
		return nil, errors.New("Anime is nil")
	}
	return &a.Url, nil
}

func (a *Anime) GetAnimeEpisodes() (*[]Episode, error) {

	if a == (*Anime)(nil) {
		return nil, errors.New("Anime is nil")
	}

	episodes := make([]Episode, len(a.Episodes))
	copy(episodes, a.Episodes)

	return &episodes, nil
}
