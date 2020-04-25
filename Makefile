default: build

build:
	@go build -o ./bin/toothbox cmd/toothbox/main.go

build.linux:
	@GOOS=linux go build -o ./bin/linux/toothbox cmd/toothbox/main.go

build.mac:
	@GOOS=linux go build -o ./bin/darwin/toothbox cmd/toothbox/main.go

build.windows:
	@GOOS=linux go build -o ./bin/windows/toothbox.exe cmd/toothbox/main.go

clean:
	@rm -f ./bin/toothbox

clean.full:
	@rm -rf ./bin/*
