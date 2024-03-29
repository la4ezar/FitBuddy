// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type Coach struct {
	ID        string `json:"ID"`
	ImageURL  string `json:"ImageUrl"`
	Name      string `json:"Name"`
	Specialty string `json:"Specialty"`
}

type Exercise struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type Goal struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	StartDate   string `json:"StartDate"`
	EndDate     string `json:"EndDate"`
	Completed   bool   `json:"Completed"`
}

type LeaderboardUser struct {
	ID        string `json:"ID"`
	UserEmail string `json:"UserEmail"`
	Score     int    `json:"Score"`
}

type Meal struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type Mutation struct {
}

type Nutrition struct {
	ID        string `json:"ID"`
	UserEmail string `json:"UserEmail"`
	MealName  string `json:"MealName"`
	Grams     int    `json:"Grams"`
	Calories  int    `json:"Calories"`
	Date      string `json:"Date"`
}

type Post struct {
	ID        string `json:"ID"`
	UserEmail string `json:"UserEmail"`
	Title     string `json:"Title"`
	Content   string `json:"Content"`
	CreatedAt string `json:"CreatedAt"`
}

type Query struct {
}

type SleepLog struct {
	ID        string `json:"ID"`
	SleepTime string `json:"SleepTime"`
	WakeTime  string `json:"WakeTime"`
}

type User struct {
	ID    string `json:"ID"`
	Email string `json:"Email"`
}

type Workout struct {
	ID           string  `json:"ID"`
	UserEmail    string  `json:"UserEmail"`
	ExerciseName string  `json:"ExerciseName"`
	Sets         int     `json:"Sets"`
	Reps         int     `json:"Reps"`
	Weight       float64 `json:"Weight"`
	Date         string  `json:"Date"`
}
