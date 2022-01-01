package database

import (
	"fmt"

	"github.com/arthur-rl/applyx/util"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserPermissons struct {
	ADMIN bool
}

type User struct {
	Id                         int            `json:"id"`
	Username                   string         `json:"username"`
	Password                   string         `json:"password"`
	Permissions                datatypes.JSON `json:"permissions"`
	Name                       string         `json:"name"`
	Email                      string         `json:"email"`
	Enabled                    bool           `json:"enabled"`
	Change_password_next_login bool           `json:"change_password_next_login"`
}

var db *gorm.DB = Connect()

func CreateNewUser(username string, password string, email string, name string, permissions datatypes.JSON) User {
	hashedPassword, _ := util.HashPassword(password)
	fmt.Println(username, password, email, name)
	user := User{
		Username:                   username,
		Password:                   hashedPassword,
		Permissions:                datatypes.JSON([]byte(permissions)),
		Email:                      email,
		Name:                       name,
		Enabled:                    true,
		Change_password_next_login: false,
	}
	db.Omit("created_at").Create(&user)
	return user
}

func GetUserById(id int) User {
	user := User{}
	db.First(&user, id)
	return user
}

func GetUserByUsername(username string) User {
	user := User{}
	db.First(&user, "username = ?", username)
	return user
}

func DoesUserExistByUsername(username string) bool {
	user := User{}
	var exists bool
	db.Model(user).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists)
	return exists
}
