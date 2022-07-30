run: game
	./bin/go-game

bin:
	mkdir bin

game: bin
	go build -o bin/go-game

clean:
	rm -rf bin/