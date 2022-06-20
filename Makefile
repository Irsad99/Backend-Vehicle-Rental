APP = backendgo
APP_EXE = "bin/$(APP)"

build:
	mkdir -p ./bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP_EXE}

test:
	go test -cover -v ./...