package entity

import (
	"html"
	"strings"
	"time"
)

type Airline struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:100;not null;unique" json:"name" validate:"required,lte=100"`
	ImageName string     `gorm:"size:255;not null" json:"image" validate:"required,lte=255"`
	IsActive  string     `gorm:"size:1;not null;default:'1'; index:idx_airline_active_status;" json:"is_active" validate:"required,lte=1"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

//BeforeSave is a gorm hook
func (a *Airline) BeforeSave() {
	a.Name = html.EscapeString(strings.TrimSpace(a.Name))
}

type Airlines []Airline

func (a *Airline) Prepare() {
	a.Name = html.EscapeString(strings.TrimSpace(a.Name))
	a.ImageName = html.EscapeString(strings.TrimSpace(a.ImageName))
	a.IsActive = html.EscapeString(strings.TrimSpace(a.IsActive))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Airline) Validate(action string) string {
	switch strings.ToLower(action) {
	case "update":
		if a.Name == "" {
			return "name required"
		}
		if a.IsActive == "" {
			return "status active required"
		}

	default:
		if a.Name == "" {
			return "name required"
		}
		if a.ImageName == "" {
			return "image name required"
		}
		if a.IsActive == "" {
			return "status active required"
		}
	}
	return ""
}
