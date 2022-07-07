# List available recipes
list:
  @just --list

run:
  @go run main.go

test:
  @echo "Make sure that bricebug/ping-pong is running with 'just start-ping'"
  @go test ./...

# Start the dummy ping pong service for testing
start-ping:
  @docker run -p 7777:80 -d --rm --name ping briceburg/ping-pong

# Stop the dummy ping pong service for testing
stop-ping:
  @docker kill ping
