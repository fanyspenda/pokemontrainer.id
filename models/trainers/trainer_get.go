package trainers

import (
	"time"
)

// GetTrainer ...
type GetTrainer struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
