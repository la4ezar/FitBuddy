package coach

import "github.com/google/uuid"

// Coach represents a coach in the application.
type Coach struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	ImageURL  string `json:"image_url"`
}

// NewCoach creates a new Coach instance.
func NewCoach(name, specialty, imageURL string) *Coach {
	return &Coach{
		ID:        uuid.New().String(),
		Name:      name,
		Specialty: specialty,
		ImageURL:  imageURL,
	}
}
