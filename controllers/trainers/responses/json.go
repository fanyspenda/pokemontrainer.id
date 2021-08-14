package responses

import (
	"pokemontrainer/business/trainers"
	"time"

	"gorm.io/gorm"
)

// TrainerResponse response to send back to FE
type TrainerResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Username  string         `json:"username"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// FromDomain convert domain data to response data
func FromDomain(domain trainers.Domain) TrainerResponse {
	return TrainerResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Address:   domain.Address,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

// FromSliceDomain convert slice of Domains to Slice of Trainer Responses
func FromSliceDomain(domains []trainers.Domain) []TrainerResponse {
	var convertedDomains = []TrainerResponse{}
	for _, v := range domains {
		convertedDomains = append(convertedDomains, FromDomain(v))
	}
	return convertedDomains
}
