package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	CarModel     string
	Manufacturer string
	Color        string
	Value        float64
}

type CarResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt,omitempty"`
	CarModel     string    `json:"carModel"`
	Manufacturer string    `json:"manufacturer"`
	Color        string    `json:"color"`
	Value        float64   `json:"value"`
}
