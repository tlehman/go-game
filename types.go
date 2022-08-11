package main

type User struct {
	Id     int
	Handle string
}

type Game struct {
	Id          int
	BlackUserId int
	WhiteUserId int
}
