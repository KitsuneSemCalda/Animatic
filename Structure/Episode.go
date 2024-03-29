package structure

import "errors"

type Episode struct {
	Number int
	Url    string
}

func IsValidEpisode(episode Episode) bool {
	if episode.Number <= 0 {
		return false
	}

	if episode.Url == "" {
		return false
	}

	return true
}

func NewEpisode(number int, url string) (*Episode, error) {
	if number <= 0 {
		return nil, errors.New("Episode number must be greater than 0")
	}

	if url == "" {
		return nil, errors.New("Episode URL cannot be empty")
	}

	return &Episode{
		Number: number,
		Url:    url,
	}, nil
}

func (e *Episode) GetEpisodeNumber() (*int, error) {
	if e == nil {
		return nil, errors.New("Episode is nil")
	}

	return &e.Number, nil
}

func (e *Episode) GetEpisodeUrl() (*string, error) {
	if e == nil {
		return nil, errors.New("Episode is nil")
	}

	return &e.Url, nil
}
