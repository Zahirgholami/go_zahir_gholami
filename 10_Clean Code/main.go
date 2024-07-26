package main

import "fmt"

type user struct {
	Email    string
	Password string
}

type userRepo struct {
	DB []user
}

func (r userRepo) Register(u user) {
	if u.Email == "" || u.Password == "" {
		fmt.Println("register failed")
	}

	r.DB = append(r.DB, u)
}
func (r userRepo) Login(u user) string {
	if u.Email == "" || u.Password == "" {
		fmt.Println("login failed")
	}
	for _, us := range r.DB {
		if us.Email == u.Email && us.Password == u.Password {
			return "auth token"
		}
	}
	return ""
}
