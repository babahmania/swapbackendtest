package persistence

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"swapbackendtest/domain/entity"
	"swapbackendtest/domain/repository"
	"swapbackendtest/infrastructure/security"
	"swapbackendtest/infrastructure/validator"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

//UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, error) {
	// Create a new validator for a User model.
	//signUp := &entity.User{}
	validate := validator.NewValidator()
	//fmt.Println(user)
	// Validate sign up fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		msgNotValid := validator.ValidatorErrors(err)
		b := new(bytes.Buffer)
		for _, value := range msgNotValid {
			//fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
			fmt.Fprintf(b, "%s. ", value)
		}
		msgError := b.String()
		return nil, errors.New(msgError)
		//return nil, errors.New("input data not validate")
	}
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("email already taken")
		}
		return nil, errors.New("database error")
	}
	return user, nil
}

func (r *UserRepo) GetUser(id uint64) (*entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *UserRepo) GetUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Debug().Find(&users).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return users, nil
}

func (r *UserRepo) GetUserByEmailAndPassword(u *entity.User) (*entity.User, error) {
	var user entity.User
	err := r.db.Debug().Where("email = ?", u.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, errors.New("database error")
	}
	//Verify the password
	err = security.VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("incorrect password")
	}
	return &user, nil
}
