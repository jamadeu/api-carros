package schemas

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	CarModel     string
	Manufacturer string
	Color        string
	Value        float64
}
