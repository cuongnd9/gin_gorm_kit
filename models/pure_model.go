// Package models gorm's models
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Pure pure/base model
type Pure struct {
	ID        uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
