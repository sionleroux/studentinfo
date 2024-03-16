package main

import (
	"embed"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
)

const DefaultPort = "8080"

//go:embed *.html favicon.png
var static embed.FS

type Student struct {
	ID         string
	Name       string
	Instrument string
	Teacher    string
}

type Students []Student

var students Students

func (ss Students) IDs() []string {
	var ids = make([]string, len(ss))
	for k, v := range ss {
		ids[k] = v.ID
	}
	return ids
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not set, defaulting to ", DefaultPort)
		port = DefaultPort
	}
	setupHandlers(http.DefaultServeMux)
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, Student{
			ID:         record[0],
			Name:       record[1],
			Instrument: record[2],
			Teacher:    record[3],
		})
	}
	log.Println("Server listening")
	http.ListenAndServe(":"+port, nil)
}
