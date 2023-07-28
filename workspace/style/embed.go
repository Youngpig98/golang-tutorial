package main

import "fmt"

type notifier interface {
	notify()
}
type user struct {
	name  string
	email string
}
type admin struct {
	user
	level string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}
	// We can access the inner type's method directly.
	ad.user.notify()
	// The inner type's method is promoted.
	ad.notify()

	user := ad.user

	sendNotification(&user)
	sendNotification(&ad)
}
