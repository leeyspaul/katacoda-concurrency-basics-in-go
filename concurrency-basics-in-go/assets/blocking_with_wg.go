package main

import "fmt"

func main() {
	var paul User

	paul.UserName = "Paul"
	paul.UserId = 137

	paul = updateUserId(paul, 200)

	fmt.Printf("id update for user paul is %d\n", paul.UserId)
}

type User struct {
	UserName string
	UserId int
}

func updateUserId(user User, userId int) (updatedUser User) {

	go func() {
		user.UserId = userId
	}()

	return user
}

func updateUserIdFixed(user User, userId int) (updatedUser User) {
	// to be filled!
	return
}