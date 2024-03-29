package structure

import (
	structure "Animatic/Structure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEpisode(t *testing.T) {
	t.Run("valid episode", func(t *testing.T) {
		episode := structure.Episode{
			Number: 1,
			Url:    "https://example.com/episode/1",
		}

		assert.Equal(t, 1, episode.Number)
		assert.Equal(t, "https://example.com/episode/1", episode.Url)
	})

	t.Run("empty episode", func(t *testing.T) {
		episode := structure.Episode{}

		assert.Zero(t, episode.Number)
		assert.Empty(t, episode.Url)
	})
}
func TestIsValidEpisode(t *testing.T) {
	t.Run("valid episode", func(t *testing.T) {
		episode := structure.Episode{
			Number: 1,
			Url:    "https://example.com/episode/1",
		}

		assert.True(t, structure.IsValidEpisode(episode))
	})

	t.Run("invalid episode with zero number", func(t *testing.T) {
		episode := structure.Episode{
			Number: 0,
			Url:    "https://example.com/episode/1",
		}

		assert.False(t, structure.IsValidEpisode(episode))
	})

	t.Run("invalid episode with empty url", func(t *testing.T) {
		episode := structure.Episode{
			Number: 1,
			Url:    "",
		}

		assert.False(t, structure.IsValidEpisode(episode))
	})
}
func TestEpisodeAdditional(t *testing.T) {

	t.Run("episode with negative number", func(t *testing.T) {
		episode := structure.Episode{
			Number: -1,
			Url:    "https://example.com/episode/-1",
		}

		assert.False(t, structure.IsValidEpisode(episode))
	})

	t.Run("episode with extremely large number", func(t *testing.T) {
		episode := structure.Episode{
			Number: 999999,
			Url:    "https://example.com/episode/999999",
		}

		assert.True(t, structure.IsValidEpisode(episode))
	})

	t.Run("episode with invalid url", func(t *testing.T) {
		episode := structure.Episode{
			Number: 1,
			Url:    "invalidurl",
		}

		assert.False(t, structure.IsValidEpisode(episode))
	})

	t.Run("episode with empty number and url", func(t *testing.T) {
		episode := structure.Episode{}

		assert.False(t, structure.IsValidEpisode(episode))
	})
}
