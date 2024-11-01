# Makefile for running the Go application

# Set the path to your main Go file
MAIN_FILE="c:/Users/LENOVO/Documents/File E/kk-requester/main.go"

# The default target
run:
	go run $(MAIN_FILE)

# Optional: Add a target to build the application
build:
	go build -o kk-requester $(MAIN_FILE)

# Clean up the build
clean:
	del kk-requester.exe
