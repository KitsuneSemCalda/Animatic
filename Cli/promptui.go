package cli

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func GetUserInput(label string) (*string, error) {
	if label == "" {
		return nil, errors.New("label cannot be empty")
	}

	prompt := promptui.Prompt{
		Label: label,
	}

	animeName, err := prompt.Run()

	if err != nil {
		return nil, err
	}

	if animeName == "" {
		return nil, errors.New("user input cannot be empty")
	}

	return &animeName, nil
}
