package entity

import (
	"html"
	"strings"
	"time"
)

type Airport struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:100;not null;unique" json:"name" validate:"required,lte=100"`
	Code      string     `gorm:"size:4;not null;unique" json:"code" validate:"required,lte=4"`
	City      string     `gorm:"size:100;not null; ; index:idx_airport_city" json:"city" validate:"required,lte=100"`
	Country   string     `gorm:"size:100;not null;" json:"country" validate:"required,lte=100"`
	IsActive  string     `gorm:"size:1;not null;default:'1'; index:idx_airport_active_status;" json:"is_active" validate:"required,lte=1"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	//Flights   []Flight   `gorm:"ForeignKey:OriginID; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

//BeforeSave is a gorm hook
func (a *Airport) BeforeSave() {
	a.Name = html.EscapeString(strings.TrimSpace(a.Name))
}

type Airports []Airport

func (a *Airport) Prepare() {
	a.Name = html.EscapeString(strings.TrimSpace(a.Name))
	a.Code = html.EscapeString(strings.TrimSpace(a.Code))
	a.City = html.EscapeString(strings.TrimSpace(a.City))
	a.Country = html.EscapeString(strings.TrimSpace(a.Country))
	a.IsActive = html.EscapeString(strings.TrimSpace(a.IsActive))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Airport) Validate(action string) string {
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
		if a.Code == "" {
			return "code required"
		}
		if a.City == "" {
			return "city required"
		}
		if a.Country == "" {
			return "country required"
		}
		if a.IsActive == "" {
			return "status active required"
		}
	}
	return ""
}
