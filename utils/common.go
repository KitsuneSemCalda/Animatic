package utils

import(
  "regexp"
  "strings"
  "path/filepath"
  "os/user"
  "log"
)

func NameTreating(str string) string {
  regexdName := DatabaseFormatter(str)
  return TreatingAnimeName(regexdName)
}

func EpisodeFormatter(str string) string {
	regex := regexp.MustCompile(`\d+$`)
  result := regex.FindString(str)
	return result
}

func DatabaseFormatter(str string) string {
	regex := regexp.MustCompile(`\s+(\d+(\.\d+)?)`)
	result := regex.ReplaceAllString(str, "")
	result = strings.TrimSpace(result)
	result = strings.ToLower(result)
	return result
}

func DownloadFolderFormatter(str string) string {
  regex := regexp.MustCompile(`https:\/\/animefire\.net\/video\/([^\/?]+)`)
	match := regex.FindStringSubmatch(str)
	if len(match) > 1 {
		finalStep := match[1]
    return finalStep
  }
  return ""
}

func TreatingAnimeName(animeName string) string {
	loweredName := strings.ToLower(animeName)
	spacelessName := strings.ReplaceAll(loweredName, " ", "-")
	return spacelessName
}

func GetFolder() string{
  currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}
  return filepath.Join(currentUser.HomeDir, "/.local/Animatic")
}
