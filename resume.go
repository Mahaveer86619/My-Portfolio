package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Resume represents the overall resume data structure.
type Resume struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Profile    ProfileInfo        `bson:"profile"`
	Projects   []Project          `bson:"projects"`
	Experience []Experience       `bson:"experience"`
	Education  []Education        `bson:"education"`
	Skills     []string           `bson:"skills"`
	Contacts   Contacts           `bson:"contacts"`
}

// ProfileInfo contains the profile information of the individual.
type ProfileInfo struct {
	Name        string `bson:"name"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	ImageURL    string `bson:"imageURL"`
}

// Project represents a single project in the resume.
type Project struct {
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	TechStack   []string `bson:"techStack"`
	RepoURL     string   `bson:"repoURL"`
}

// Experience represents a work experience entry.
type Experience struct {
	Company     string   `bson:"company"`
	Position    string   `bson:"position"`
	StartDate   string   `bson:"startDate"`
	EndDate     string   `bson:"endDate"`
	Achievements []string `bson:"achievements"`
}

// Education represents an educational entry.
type Education struct {
	Institution string `bson:"institution"`
	Degree      string `bson:"degree"`
	StartDate   string `bson:"startDate"`
	EndDate     string `bson:"endDate"`
	Location    string `bson:"location"`
}

// Contacts represents the contact information.
type Contacts struct {
	LinkedIn string `bson:"linkedIn"`
	GitHub   string `bson:"gitHub"`
	Threads  string `bson:"threads"`
	Email    string `bson:"email"`
	Phone    string `bson:"phone"`
}