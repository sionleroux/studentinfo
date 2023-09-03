package main

import (
	"embed"
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

var students = Students{
	Student{
		ID:         "1234",
		Name:       "Alice",
		Instrument: "Intermediate Guitar",
		Teacher:    "Professor Porcupine",
	},
	Student{
		ID:         "4444",
		Name:       "Bob",
		Instrument: "Beginner Piano",
		Teacher:    "Master Manatee",
	},
}

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
	setupHandlers()
	log.Println("Server listening")
	http.ListenAndServe(":"+port, nil)
}
