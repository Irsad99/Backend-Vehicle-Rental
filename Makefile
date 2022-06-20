APP = BackendGo
APP_EXE = "./build/$(APP)"

build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP_EXE}

test:
	go test -cover -v ./...