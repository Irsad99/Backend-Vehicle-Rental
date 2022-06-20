APP = BackendGo
APP_EXE = "./build/$(APP)"
cmd = mkdir -p ./build && CGO_ENABLED=0 GOOS=linux go build -o ${APP_EXE}

build:
	go build -o bin/main

compile:
	echo "Compiling..."
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

test:
	go test -cover -v ./...