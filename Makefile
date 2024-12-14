# Variables
BINARY_NAME=gc_check
MAIN_FILE=main.go
PORT=8080

# Targets
build:
	@echo "Building the Go server..."
	go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "Build complete. Binary created: $(BINARY_NAME)"

run: build
	@echo "Running the Go server..."
	./$(BINARY_NAME)

curl:
	@echo "Enter N (in KB):"
	@read n_kb; \
	n_bytes=$$(($$n_kb * 1024)); \
	echo "Enter GC type (manual or auto):"; \
	read gc_type; \
	echo "Executing command: curl -X GET \"http://localhost:$(PORT)/allocate?N=$$n_bytes&gc=$$gc_type\""; \
	curl -X GET "http://localhost:$(PORT)/allocate?N=$$n_bytes&gc=$$gc_type"


help:
	@echo "Available Makefile targets:"
	@echo "  build  - Build the Go server binary."
	@echo "  run    - Build and run the Go server."
	@echo "  curl   - Prompt for parameters and send a request to the server."
	@echo "  help   - Show this help message."

