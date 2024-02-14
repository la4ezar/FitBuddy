package coach

// Coach represents a coach in the application.
type Coach struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	ImageURL  string `json:"image_url"`
}
