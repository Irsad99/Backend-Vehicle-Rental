APP = BackendGo
APP_EXE = "./build/$(APP)"

build:
	go build -o bin/main

run:
	go run main.go server

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

test:
	go test -cover -v ./...