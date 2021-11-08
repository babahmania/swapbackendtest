package persistence

import (
	"swapbackendtest/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = entity.User{}
	user.Email = "maniababah@gmail.com"
	user.FullName = "maniababah"
	user.UserStatus = "1"
	user.UserRole = "0"
	user.Password = "maniababah@gmail.com"

	repo := NewUserRepository(conn)

	u, saveErr := repo.SaveUser(&user)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, u.Email, "maniababah@gmail.com")
	assert.EqualValues(t, u.FullName, "maniababah")
	assert.EqualValues(t, u.UserStatus, "1")
	assert.EqualValues(t, u.UserRole, "0")
	//The pasword is supposed to be hashed, so, it should not the same the one we passed:
	//assert.NotEqual(t, u.Password, "maniababah@gmail.com")
}

//Failure can be due to duplicate email, etc
//Here, we will attempt saving a user that is already saved
func TestSaveUser_Failure(t *testing.T) {

	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the user
	_, err = seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	var user = entity.User{}
	user.Email = "babahmania@gmail.com"
	user.FullName = "babahmania@gmail.com"
	user.UserStatus = "1"
	user.Password = "babahmania@gmail.com"

	repo := NewUserRepository(conn)
	u, saveErr := repo.SaveUser(&user)
	dbMsg := map[string]string{
		"email_taken": "email already taken",
	}
	assert.Nil(t, u)
	assert.EqualValues(t, dbMsg, saveErr)
}

func TestGetUser_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	//seed the user
	user, err := seedUser(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUser(user.ID)

	assert.Nil(t, getErr)
	assert.EqualValues(t, u.Email, "babahmania@gmail.com")
	assert.EqualValues(t, u.FullName, "babah mania")
	assert.EqualValues(t, u.UserStatus, "1")
}

func TestGetUsers_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}

	//seed the users
	_, err = seedUsers(conn)
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	repo := NewUserRepository(conn)
	users, getErr := repo.GetUsers()

	assert.Nil(t, getErr)
	assert.EqualValues(t, len(users), 2)
}

func TestGetUserByEmailAndPassword_Success(t *testing.T) {
	conn, err := DBConn()
	if err != nil {
		t.Fatalf("want non error, got %#v", err)
	}
	/*
		//seed the user
		_, err = seedUser(conn)
		if err != nil {
			t.Fatalf("want non error, got %#v", err)
		}
	*/
	var user = &entity.User{
		Email:    "babahmania@gmail.com",
		Password: "babahmania@gmail.com",
	}
	repo := NewUserRepository(conn)
	u, getErr := repo.GetUserByEmailAndPassword(user)

	assert.Nil(t, getErr)
	//assert.EqualValues(t, u.Email, user.Email)
	//Note, the user password from the database should not be equal to a plane password, because that one is hashed
	assert.NotEqual(t, u.Password, user.Password)
}
