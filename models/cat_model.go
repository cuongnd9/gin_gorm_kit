// Package models gorm's models
package models

// Cat cat model struct
type Cat struct {
	Pure       `gorm:"embedded"`
	Name       string   `json:"name"`
	Color      string   `json:"color"`
	CategoryID string   `json:"category_id"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
}

// TableName get table name
func (cat *Cat) TableName() string {
	return "cats"
}
