package models

import "github.com/jinzhu/gorm"

// Crop represents a crop entity in the system
type Crop struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Health       string `json:"health"`
	GrowthStage  string `json:"growth_stage"`
	Description  string `json:"description"`
}
