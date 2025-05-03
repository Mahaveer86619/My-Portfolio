package resume

import "time"

type Resume struct {
	Profile    ProfileInfo  `bson:"profile"`
	Projects   []Project    `bson:"projects"`
	Experience []Experience `bson:"experience"`
	Education  []Education  `bson:"education"`
	Skills     []string     `bson:"skills"`
	Contacts   Contacts     `bson:"contacts"`
}

type ProfileInfo struct {
	Name        string `bson:"name"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Image       string `bson:"image"`
}

type Project struct {
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	TechStack   []string `bson:"techStack"`
	RepoLink    string   `bson:"repoLink"`
}

type Experience struct {
	Company     string    `bson:"company"`
	Role        string    `bson:"role"`
	StartDate   time.Time `bson:"startDate"`
	EndDate     time.Time `bson:"endDate"`
	Achievements []string  `bson:"achievements"`
}

type Education struct {
	Institution string    `bson:"institution"`
	Degree      string    `bson:"degree"`
	StartDate   time.Time `bson:"startDate"`
	EndDate     time.Time `bson:"endDate"`
	Location    string    `bson:"location"`
}

type Contacts struct {
	LinkedIn string `bson:"linkedIn"`
	GitHub   string `bson:"gitHub"`
	Threads  string `bson:"threads"`
	Email    string `bson:"email"`
	Phone    string `bson:"phone"`
}

func GetSampleResume() Resume {
	return Resume{
		Profile: ProfileInfo{
			Name:        "John Doe",
			Title:       "Software Engineer",
			Description: "Experienced software engineer with a passion for building innovative solutions.",
			Image:       "/static/img/profile.jpg",
		},
		Projects: []Project{
			{
				Title:       "Project 1",
				Description: "Description for Project 1.",
				TechStack:   []string{"Go", "React", "MongoDB"},
				RepoLink:    "https://github.com/user/project1",
			},
			{
				Title:       "Project 2",
				Description: "Description for Project 2.",
				TechStack:   []string{"Python", "Django", "PostgreSQL"},
				RepoLink:    "https://github.com/user/project2",
			},
		},
		Experience: []Experience{
			{
				Company:     "Company A",
				Role:        "Software Engineer",
				StartDate:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				Achievements: []string{"Achievement 1", "Achievement 2"},
			},
			{
				Company:     "Company B",
				Role:        "Senior Software Engineer",
				StartDate:   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Now(),
				Achievements: []string{"Achievement 3", "Achievement 4"},
			},
		},
		Education: []Education{
			{
				Institution: "University X",
				Degree:      "Bachelor of Science in Computer Science",
				StartDate:   time.Date(2016, 9, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC),
				Location:    "City Y",
			},
		},
		Skills: []string{"Go", "Python", "JavaScript", "React", "MongoDB", "PostgreSQL"},
		Contacts: Contacts{
			LinkedIn: "https://www.linkedin.com/in/johndoe/",
			GitHub:   "https://github.com/johndoe",
			Threads:  "https://www.threads.net/@johndoe",
			Email:    "johndoe@example.com",
			Phone:    "+15551234567",
		},
	}
}