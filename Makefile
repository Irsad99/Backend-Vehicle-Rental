APP = backendgo
APP_EXE = "./build/$(APP)"

build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=windows go build -o ${APP_EXE}

test:
	go test -cover -v ./...