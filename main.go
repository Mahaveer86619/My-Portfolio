// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Outyet is a web server that announces whether or not a particular Go version
// has been tagged. modified to be a web resume.
package main

import (	
	"encoding/json"
	"context"
	"html/template"
	"log"	
	"os"
	"time"
	"net/http"


	"portfolio-app/resume"	
)

var mongo *resume.Mongo

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Initialize MongoDB connection
	var err error
	mongo, err = resume.ConnectDB(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer resume.DisconnectDB(ctx, mongo)

	if os.Getenv("POPULATE_DB") == "true" {
		populateDB(ctx, mongo)
	}

	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
	
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/resume", resumeHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/experience", experienceHandler)
	http.HandleFunc("/education", educationHandler)
	http.HandleFunc("/skills", skillsHandler)
	http.HandleFunc("/contacts", contactsHandler)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	resumeData, err := mongo.GetResume(context.Background())
	if err == nil {
		renderTemplate(w, "templates/index.tmpl", resumeData)
	} else {
		log.Printf("Error getting resume: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {
	resumeData, err := mongo.GetResume(context.Background())
	if err == nil {
		renderJSON(w, resumeData)
	} else {
		log.Printf("Error getting resume: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	resumeData, _ := mongo.GetResume(context.Background())
	renderTemplate(w, "templates/projects.tmpl", resumeData.Projects)
}

func experienceHandler(w http.ResponseWriter, r *http.Request) {
	resumeData, _ := mongo.GetResume(context.Background())
	renderTemplate(w, "templates/experience.tmpl", resumeData.Experience)
}

func educationHandler(w http.ResponseWriter, r *http.Request) {
	resumeData, _ := mongo.GetResume(context.Background())
	renderTemplate(w, "templates/education.tmpl", resumeData.Education)
}

func skillsHandler(w http.ResponseWriter, r *http.Request) {
	resumeData, _ := mongo.GetResume(context.Background())
	renderTemplate(w, "templates/skills.tmpl", resumeData.Skills)
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	resumeData, _ := mongo.GetResume(context.Background())
	renderTemplate(w, "templates/contacts.tmpl", resumeData.Contacts)
}

func populateDB(ctx context.Context, mongo *resume.Mongo) {
	resumeData := resume.GetSampleResume()
	_, err := mongo.Collection.InsertOne(ctx, resumeData)
	if err != nil {
		log.Fatalf("Error inserting sample data: %v", err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl, "templates/base.tmpl")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "base", data)
}

func renderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}
