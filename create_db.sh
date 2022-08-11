#!/bin/sh

(cd db && (sqlite3 go-game.db < schema.sql))
