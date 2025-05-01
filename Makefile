
APP_NAME = vis

build:
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

clean: rm -f $(APP_NAME)