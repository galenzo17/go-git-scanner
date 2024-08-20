.PHONY: build run clean

build:
	go build -o git-scanner src/*.go

run: build
	./git-scanner

clean:
	rm -f git-scanner