package trainers

import (
	"time"

	"gorm.io/gorm"
)

//Trainer ...
type Trainer struct {
	ID        uint           `gorm:"primarykey"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
