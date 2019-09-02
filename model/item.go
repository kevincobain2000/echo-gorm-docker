package model

import (
	"database/sql"
	"time"
)

/**
 * Item table
 */
type Item struct {
	Address   string         `gorm:"column:address"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	ID        int64          `gorm:"column:id;primary_key"`
	Keywords  sql.NullString `gorm:"column:keywords"`
	Name      string         `gorm:"column:name"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}
