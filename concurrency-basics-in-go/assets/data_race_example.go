package main

import "fmt"

func main() {
	var paul User

	paul.UserName = "Paul"
	paul.UserId = 137

	paulUserId = getUserId(paul) // what will this return?

	fmt > printf("the user id retrieved for paul is %d", paulUserId)
}

type User struct {
	UserName string
	UserId   int
}

func getUserId(user User) (id int) {
	var userId int

	go func() {
		userId = user.UserId
	}()

	return userId
}
