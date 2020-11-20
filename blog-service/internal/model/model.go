package model

import "time"

type Model struct {
	ID           int       `gorm:"primary_key" json:"id"`
	CreatorId    int       `json:"creator_id"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedName  string    `json:"created_name"`
	ModifierId   int       `json:"modifier_id"`
	ModifierName string    `json:"modifier_name"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsRemoved    int       `json:"is_removed"`
}
