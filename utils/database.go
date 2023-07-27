package utils

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func InitializeDB() (*sql.DB, error){
  path := GetFolder()
  
  _, exists := os.Stat(path)

  if os.IsNotExist(exists){  
    err := os.MkdirAll(path, os.ModePerm)
    if err != nil{
      return nil, err
    }
  }

  path = filepath.Join(path, "downloaded.db")

  db, err := sql.Open("sqlite3", path)

  if err != nil{
    return nil, err
  }

  createTableSQL := `
    CREATE TABLE IF NOT EXISTS anime(
      id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
      name TEXT,
      episode TEXT,
      path TEXT
  );
  `

  _, err = db.Exec(createTableSQL)

  if err != nil{
    return nil, err
  }

  return db, nil
}

func AddAnimeToTB(db *sql.DB, animeName string, episode string, path string) error {
  insertSQL := `
    INSERT INTO anime (name, episode, path) VALUES (?, ?, ?);
  `
  
  animeName = DatabaseFormatter(animeName)
  episode = EpisodeFormatter(episode)

  _, err := db.Exec(insertSQL,animeName, episode, path)

  if err != nil {
    return err
  }

  return nil
}
