BINARY_NAME=go_steam_assetclassinfo

build:
	go build  -o dist/${BINARY_NAME} ./main.go

run: build
	dist/${BINARY_NAME}

clean:
	go clean
