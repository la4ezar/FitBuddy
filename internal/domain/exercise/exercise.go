package exercise

// Exercise represents a fitness exercise in the application.
type Exercise struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
