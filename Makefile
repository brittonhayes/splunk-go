dependencies:
	go mod tidy

format:
	go fmt .

docs:
	go run main.go docs

security:
	gosec -severity high -confidence high -quiet ./...

build:
	go build -o ./bin/splunk-go main.go

cross-compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/splunk-go-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/splunk-go-windows-386 main.go

run:
	go run main.go

all: dependencies format security docs build