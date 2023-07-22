package config

import (
	"bufio"
	"fmt"
  "path/filepath"
	"os"
  "runtime"
	"strconv"
	"strings"

  "golang.org/x/text/language"
)

type Config struct {
  parallelValue int
  portugueseSearch bool
  englishSearch bool
}

func writeConfigToFile(file *os.File, config *Config) error {
	writer := bufio.NewWriter(file)

	_, err := fmt.Fprintf(writer, "parallelValue: %d\n", config.parallelValue)

	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "portugueseSearch: %v\n", config.portugueseSearch)

	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(writer, "englishSearch: %v\n", config.englishSearch)

	if err != nil {
		return err
	}

	err = writer.Flush() // Flush the buffered writer to ensure data is written to the file.
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(path string) (*Config, error) {
	Configpath := filepath.Join(path, "settings.txt")

	_, err := os.Stat(Configpath)
	if err == nil {
		// The file exists, let's read it and load the Config struct.
		file, err := os.Open(Configpath)
		if err != nil {
			return nil, fmt.Errorf("error opening the configuration file: %v", err)
		}
		defer file.Close()

		config := &Config{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue // Skip lines that are not in the format "key: value"
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "parallelValue":
				val, err := strconv.Atoi(value)
				if err != nil {
					return nil, fmt.Errorf("error parsing parallelValue: %v", err)
				}
				config.parallelValue = val
			case "portugueseSearch":
				bval, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("error parsing portugueseSearch: %v", err)
				}
				config.portugueseSearch = bval
			case "englishSearch":
				bval, err := strconv.ParseBool(value)
				if err != nil {
					return nil, fmt.Errorf("error parsing englishSearch: %v", err)
				}
				config.englishSearch = bval
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading the configuration file: %v", err)
		}

		return config, nil
	} else if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return nil, fmt.Errorf("error creating the configuration folder: %v", err)
		}

		file, err := os.Create(Configpath)
		if err != nil {
			return nil, fmt.Errorf("error creating the configuration file: %v", err)
		}
		defer file.Close()

    
    // Get the current system locale
		locale := language.MustParse("en-US")

		// Check if the locale matches either Brazilian Portuguese or European Portuguese
		matcher := language.NewMatcher([]language.Tag{
			language.BrazilianPortuguese,
			language.EuropeanPortuguese,
		})
		tag, _ := language.MatchStrings(matcher, locale.String())
		portugueseSearch := tag == language.BrazilianPortuguese || tag == language.EuropeanPortuguese
    
    numCores := runtime.NumCPU() / 2
    
		// Populate the default configuration values based on the locale and pc Power.
		defaultConfig := &Config{
			parallelValue:   numCores,
			portugueseSearch: portugueseSearch,
			englishSearch:    !portugueseSearch,
		}

		// Write the default configuration to the new file.
		err = writeConfigToFile(file, defaultConfig)
		if err != nil {
			return nil, fmt.Errorf("error writing default configuration to file: %v", err)
		}

		return defaultConfig, nil
	} else {
		return nil, fmt.Errorf("error while checking the configuration file: %v", err)
	}
}
