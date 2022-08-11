package main

import (
	"context"
	"errors"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
)

var dbpool *sqlitex.Pool

func initDbPool() {
	poolSize := 10 // this is the number of connections in the pool
	var err error
	dbpool, err = sqlitex.Open("db/go-game.db", sqlite.SQLITE_OPEN_READWRITE, poolSize)
	if err != nil {
		panic(err)
	}
}

func selectAllUsers(c context.Context) ([]User, error) {
	var registeredUsers []User = make([]User, 0)
	conn := dbpool.Get(c)
	if conn == nil {
		return registeredUsers, errors.New("connection to db failed")
	}
	// query all registered users
	defer dbpool.Put(conn) // put the conn back in the pool
	stmt := conn.Prep("SELECT id, handle FROM users;")
	var user User
	for {
		if hasRow, err := stmt.Step(); err != nil {
			return registeredUsers, err
		} else if !hasRow {
			break
		}
		user = User{
			Id:     int(stmt.GetInt64("id")),
			Handle: stmt.GetText("handle"),
		}
		registeredUsers = append(registeredUsers, user)
	}
	return registeredUsers, nil
}

func selectAllGames(c context.Context) ([]Game, error) {
	var games []Game = make([]Game, 0)
	conn := dbpool.Get(c)
	if conn == nil {
		return games, errors.New("connection to db failed")
	}
	// query all registered users
	defer dbpool.Put(conn) // put the conn back in the pool
	stmt := conn.Prep("SELECT id, black_user_id, white_user_id FROM games;")
	var game Game
	for {
		if hasRow, err := stmt.Step(); err != nil {
			return games, err
		} else if !hasRow {
			break
		}
		game = Game{
			Id:          int(stmt.GetInt64("id")),
			BlackUserId: int(stmt.GetInt64("black_user_id")),
			WhiteUserId: int(stmt.GetInt64("white_user_id")),
		}
		games = append(games, game)
	}
	return games, nil
}
