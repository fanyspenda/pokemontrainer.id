package responses

import (
	"pokemontrainer/business/gyms"
	"time"
)

// GymResponse response data structure
type GymResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// FromDomain convert response from Domain to Response Format
func FromDomain(gymDomain gyms.Domain) GymResponse {
	return GymResponse{
		ID:        gymDomain.ID,
		Name:      gymDomain.Name,
		Address:   gymDomain.Address,
		CreatedAt: gymDomain.CreatedAt,
		UpdatedAt: gymDomain.UpdatedAt,
		DeletedAt: gymDomain.DeletedAt.Time,
	}
}
