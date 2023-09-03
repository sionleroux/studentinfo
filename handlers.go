package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func setupHandlers() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request to /, serving main page")

		if r.URL.Path != "/" {
			log.Println("[ERROR] URL doesn't match /, so 404")
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		file, err := static.ReadFile("index.html")
		if err != nil {
			log.Println("[ERROR] Can't read file:", err)
			http.Error(w, "Error accessing file", http.StatusInternalServerError)
			return
		}

		var t = template.Must(template.New("index").Parse(string(file)))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.Execute(w, strings.Join(students.IDs(), ", "))
		if err != nil {
			log.Println("[ERROR] Failed to execute student template: ", err)
			http.Error(w, "Error accessing file", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		log.Println("Received request to /students, fetching student data with ID: ", id)

		if id == "" {
			log.Println("[ERROR] No ID was provided in the request, aborting")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		validatorID := regexp.MustCompile("^[0-9]{4}$")
		if !validatorID.MatchString(id) {
			log.Println("[ERROR] An invalid ID was provided in the request, aborting")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "The student ID must be a 4-digit number unique to the student")
			return
		}

		file, err := static.ReadFile("student.html")
		if err != nil {
			log.Println("[ERROR] Can't read file:", err)
			http.Error(w, "Error accessing file", http.StatusInternalServerError)
			return
		}

		var data *Student
		for _, s := range students {
			if s.ID == id {
				data = &s
				break
			}
		}
		if data == nil {
			log.Println("[ERROR] Student not found with ID:", id)
		}

		var t = template.Must(template.New("student").Parse(string(file)))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.Execute(w, data)
		if err != nil {
			log.Println("[ERROR] Failed to execute student template: ", err)
			http.Error(w, "Error accessing file", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/favicon.png", func(w http.ResponseWriter, r *http.Request) {
		file, err := static.ReadFile("favicon.png")
		if err != nil {
			log.Println("[ERROR] Can't read file:", err)
			http.Error(w, "Error accessing file", http.StatusInternalServerError)
			return
		}
		w.Write(file)
	})

}
