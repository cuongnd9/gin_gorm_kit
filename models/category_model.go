// Package models gorm's models
package models

// Category category model struct
type Category struct {
	Pure `gorm:"embedded"`
	Name string `json:"name"`
	Cats []Cat  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cats"`
}

// TableName get table name
func (category *Category) TableName() string {
	return "categories"
}
