package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	Cli "Animatic/Cli"
)

type MockCli struct {
	mock.Mock
}

func (m *MockCli) GetUserInput(label string) (string, error) {
	args := m.Called(label)
	return args.String(0), args.Error(1)
}

func TestGetUserInput_ValidInput(t *testing.T) {
	mockCli := new(MockCli)

	input := "Test Anime"
	label := "Anime Name"
	mockCli.On("GetUserInput", label).Return(input, nil)

	result, err := mockCli.GetUserInput(label)

	require.NoError(t, err)

	assert.Equal(t, input, result)
}

func TestGetUserInput(t *testing.T) {
	t.Run("empty label", func(t *testing.T) {
		label := ""

		_, err := Cli.GetUserInput(label)

		assert.Error(t, err)
	})

	t.Run("empty input", func(t *testing.T) {
		label := "Anime Name"

		_, err := Cli.GetUserInput(label)

		assert.Error(t, err)
	})
}
