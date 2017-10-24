buildDir=./build

build:
	go build -v -o ${buildDir}/api cmd/api/main.go

run:
	go build -v -o ${buildDir}/api cmd/api/main.go
	${buildDir}/api

test:
	go test -v ./pkg/...

dep:
	dep ensure -v

clean:
	rm -r vendor