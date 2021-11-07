package entity

import (
	"html"
	"strings"
	"swapbackendtest/infrastructure/security"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	FullName   string     `gorm:"size:100;not null;" json:"fullname" validate:"required,lte=100"`
	Email      string     `gorm:"size:100;not null;unique" json:"email" validate:"required,email,lte=100"`
	Password   string     `gorm:"size:100;not null;" json:"password"`
	UserStatus string     `gorm:"size:1;not null;default:'1'" json:"user_status" validate:"required,lte=1"`
	UserRole   string     `gorm:"size:1;not null;default:'0'" json:"user_role" validate:"required,lte=1"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

type PublicUser struct {
	ID         uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FullName   string `gorm:"size:100;not null;" json:"fullname"`
	UserStatus string `gorm:"size:1;not null;" json:"user_status"`
	UserRole   string `gorm:"size:1;not null;" json:"user_role"`
}

//BeforeSave is a gorm hook
func (u *User) BeforeSave() error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

type Users []User

//So that we dont expose the user's email address and password to the world
func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}

//So that we dont expose the user's email address and password to the world
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:         u.ID,
		FullName:   u.FullName,
		UserStatus: u.UserStatus,
		UserRole:   u.UserRole,
	}
}

func (u *User) Prepare() {
	u.FullName = html.EscapeString(strings.TrimSpace(u.FullName))
	u.UserStatus = html.EscapeString(strings.TrimSpace(u.UserStatus))
	u.UserRole = html.EscapeString(strings.TrimSpace(u.UserRole))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) string {
	var err error
	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			return "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				return "email email"
			}
		}

	case "login":
		if u.Password == "" {
			return "password is required"
		}
		if u.Email == "" {
			return "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				return "please provide a valid email"
			}
		}
	case "forgotpassword":
		if u.Email == "" {
			return "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				return "please provide a valid email"
			}
		}
	default:
		if u.FullName == "" {
			return "full name is required"
		}
		if u.UserStatus == "" {
			return "status user is required"
		}
		if u.UserRole == "" {
			return "role user is required"
		}
		if u.Password == "" {
			return "password is required"
		}
		if u.Password != "" && len(u.Password) < 6 {
			return "password should be at least 6 characters"
		}
		if u.Email == "" {
			return "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				return "please provide a valid email"
			}
		}
	}
	return ""
}
