run: game
	./bin/go-game

bin:
	mkdir bin

game: bin
	go build -o bin/go-game

db:
	./create_db.sh

query: db
	sqlite3 db/go-game.db

clean:
	rm -rf bin/