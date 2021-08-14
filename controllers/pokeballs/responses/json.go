package responses

import (
	"pokemontrainer/business/pokeballs"
	"time"

	"gorm.io/gorm"
)

// Response basic
type Response struct {
	Name        string         `json:"name"`
	SuccessRate float32        `json:"success_rate"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// FromDomain convert domain to response data
func FromDomain(domainData pokeballs.Domain) Response {
	return Response{
		Name:        domainData.Name,
		SuccessRate: domainData.SuccessRate,
		CreatedAt:   domainData.CreatedAt,
		UpdatedAt:   domainData.UpdatedAt,
		DeletedAt:   domainData.DeletedAt,
	}
}
