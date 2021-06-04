all: clean build run
build:
	go build -o bin/main.exe cmd/api/main.go
run:
	./bin/main.exe
clean:
	powershell if (Test-Path ./bin/) {rm -r ./bin/}
docker:
	docker compose up -d --build