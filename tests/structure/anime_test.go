package structure

import (
	structure "Animatic/Structure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnime(t *testing.T) {
	t.Run("valid anime", func(t *testing.T) {
		anime := structure.Anime{
			Name: "Example Anime",
			Url:  "https://example.com/anime/1",
			Episodes: []structure.Episode{
				{Number: 1, Url: "https://example.com/anime/1/episode/1"},
			},
		}

		assert.Equal(t, "Example Anime", anime.Name)
		assert.Equal(t, "https://example.com/anime/1", anime.Url)
		assert.Equal(t, 1, len(anime.Episodes))
		assert.Equal(t, 1, anime.Episodes[0].Number)
		assert.Equal(t, "https://example.com/anime/1/episode/1", anime.Episodes[0].Url)
	})

	t.Run("empty anime", func(t *testing.T) {
		anime := structure.Anime{}

		assert.Empty(t, anime.Name)
		assert.Empty(t, anime.Url)
		assert.Empty(t, anime.Episodes)
	})
}
