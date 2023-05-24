package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

// Video represents the video data structure
type Video struct {
	ID        int
	Title     string
	DateTime  string
	FilePath  string
}

// Config represents the MySQL configuration structure
type Config struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

var tmpl *template.Template
var config Config

func main() {
	// Load MySQL configuration from config.json file
	err := loadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to load MySQL configuration:", err)
	}

	http.HandleFunc("/", indexHandler)
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Host, config.DBName))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	query := "SELECT ID, Title, DateTime, FilePath FROM Videos"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer rows.Close()

	videos := []Video{}
	for rows.Next() {
		video := Video{}
		err := rows.Scan(&video.ID, &video.Title, &video.DateTime, &video.FilePath)
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		videos = append(videos, video)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Failed to fetch rows:", err)
	}

	if tmpl == nil {
		tmpl, err = template.ParseFiles("templates/index.html")
		if err != nil {
			log.Fatal("Failed to parse HTML template:", err)
		}
	}

	err = tmpl.Execute(w, videos)
	if err != nil {
		log.Fatal("Failed to render HTML template:", err)
	}
}

func loadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	return nil
}

func urlEncode(s string) string {
	return url.QueryEscape(s)
}
