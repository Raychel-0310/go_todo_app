package data

import (
	"sync"
	"todo_app/models"
)

var (
	users []models.User
	mu sync.Mutex
)

func GetUsers() []models.User {
	return users
}

func AddUser(user models.User) {
	users = append(users, user)
}

func GetUserByID(id int) *models.User {
	for _, u := range users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}