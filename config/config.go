package config

import (
	"Animatic/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/language"
)

type Config struct {
  downloadPath string
  portugueseSearch bool
  englishSearch bool
  downloadAll bool
}

// Define a constructor

func NewConfig(portugueseSearch, englishSearch, downloadAll bool, downloadPath string) *Config {
	return &Config{
		portugueseSearch: portugueseSearch,
		englishSearch:    englishSearch,
		downloadAll:       downloadAll,
    downloadPath: downloadPath,
	}
}
// Define getter methods

func (c *Config) DownloadPath() string{
  return c.downloadPath
}

func (c *Config) PortugueseSearch() bool {
	return c.portugueseSearch
}

func (c *Config) EnglishSearch() bool {
	return c.englishSearch
}

func (c *Config) DownloadAll() bool {
  return c.downloadAll
}

func GetPathSettings() string {
  userHomeDir, err := os.UserHomeDir()
  if err != nil {
    log.Fatalf("Error in push user directory: %v\n", err)
    return ""
  }

  return filepath.Join(userHomeDir, ".local/Animatic")
}

// Create a Exporter Function
func writeConfigToFile(file *os.File, config *Config) error {
	writer := bufio.NewWriter(file)

  _, err := fmt.Fprintf(writer, "portugueseSearch: %v\n", config.portugueseSearch)

	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(writer, "englishSearch: %v\n", config.englishSearch)
  
  if err != nil {
    return err
  }

  _, err = fmt.Fprintf(writer, "downloadPath: %v\n", config.downloadPath)

	if err != nil {
		return err
	}

  _, err = fmt.Fprintf(writer, "downloadAll : %v\n", config.downloadAll)

  if err != nil {
    return err
  }

	err = writer.Flush() // Flush the buffered writer to ensure data is written to the file.
	if err != nil {
		return err
	}

	return nil
}

// Create a importer Function
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
      case "downloadAll":
        bval, err := strconv.ParseBool(value)
        if err != nil {
          return nil, fmt.Errorf("error parsing downloadAll: %v", err)
        }
        config.downloadAll = bval
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
    
		// Populate the default configuration values based on the locale and pc Power.
		defaultConfig := &Config{
			portugueseSearch: portugueseSearch,
			englishSearch:    !portugueseSearch,
      downloadAll: true,
      downloadPath:  filepath.Join(utils.GetFolder(), "downloads/animes/"),
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
