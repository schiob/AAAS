build:
	goagen app -d github.com/schiob/AAAS/design
	goagen main -d github.com/schiob/AAAS/design
	goagen swagger -d github.com/schiob/AAAS/design
	go build .
