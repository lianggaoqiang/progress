BuildCmd = CGO_ENABLED=0 GOARCH=amd64 go build -o _ ./p/t.go

.PHONY: build
build: build-sys build-win build-linux build-darwin

.PHONY: build-sys
build-sys:
	go build -o bin-sys ./p/t.go

.PHONY: build-win
build-win:
	GOOS=windows $(patsubst _, bin-win.exe, $(BuildCmd))

.PHONY: build-linux
build-linux:
	GOOS=linux $(patsubst _, bin-linux, $(BuildCmd))

.PHONY: build-darwin
build-darwin:
	GOOS=darwin $(patsubst _, bin-darwin, $(BuildCmd))

.PHONY: clean
clean:
	rm -f ./bin-*