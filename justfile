# List available recipes
list:
  @just --list

# Run the program
run:
  @go run main.go

# Start the dummy ping pong service for testing
start-ping:
  @docker run -p 8080:80 -d --rm --name ping briceburg/ping-pong

# Stop the dummy ping pong service for testing
stop-ping:
  @docker kill ping
