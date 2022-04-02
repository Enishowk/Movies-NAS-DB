package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("start save...")
	start := time.Now()

	s, err := ConnectSMB()
	defer s.Logoff()
	if err != nil {
		panic(err)
	}

	fs, err := s.Mount(os.Getenv("NAS_SHARED_FOLDER"))
	defer fs.Umount()
	if err != nil {
		panic(err)
	}

	files, err := fs.ReadDir(os.Getenv("NAS_DIRECTORY"))
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	defer db.Close()

	movies, err := getAllMovies(db)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		title, year, quality := getFilenameData(file.Name())

		if _, found := movies[title]; !found {
			stmt, err := db.Prepare("INSERT INTO movies(id, title, year, quality, size, date) VALUES (?, ?, ?, ?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}

			_, err = stmt.Exec(uuid.New(), title, year, quality, file.Size(), file.ModTime().Format(time.RFC3339))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	duration := time.Since(start)
	fmt.Println("finished in :", duration)
}

type Movie struct {
	id    string
	title string
}

func getAllMovies(db *sql.DB) (map[string]string, error) {
	rows, err := db.Query("SELECT id, title FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := make(map[string]string)
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.id, &movie.title); err != nil {
			return nil, err
		}
		movies[movie.title] = movie.id
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func getFilenameData(filename string) (title string, year string, quality string) {
	fmtString := strings.FieldsFunc(filename, func(r rune) bool {
		return r == '(' || r == ')' || r == '.'
	})

	title = strings.TrimSpace(fmtString[0])
	year = fmtString[1]
	quality = strings.TrimSpace(fmtString[2])

	return title, year, quality
}
