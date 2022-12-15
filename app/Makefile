clean:
	rm -rf tmp/*

dependencies:
	go mod download

build:
	go build -o tmp/app ./main.go

build\:playground:
	go build -o tmp/playground ./playground/main.go

migrate\:build:
	go build -o tmp/migrate cmd/migration/main.go
