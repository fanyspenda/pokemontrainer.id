package responses

import (
	"pokemontrainer/business/trainers"
	"time"
)

// TrainerResponse response to send back to FE
type TrainerResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func fromDomain(domain trainers.Domain) TrainerResponse {
	return TrainerResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Address:   domain.Address,
		Username:  domain.Address,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt.Time,
	}
}
