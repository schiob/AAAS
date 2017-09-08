build:
	goagen app -d github.com/schiob/dijkstra/design
	goagen main -d github.com/schiob/dijkstra/design
	goagen swagger -d github.com/schiob/dijkstra/design
	go build .
