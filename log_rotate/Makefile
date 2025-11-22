APP_NAME := log_rotate

MAIN := ./main.go

.PHONY: build
build:
	@if [ ! -f go.mod ]; then \
		echo "==> go.mod not found — initializing module"; \
		go mod init $(APP_NAME); \
		go mod tidy; \
	fi
	@echo "==> Building $(APP_NAME)"
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME) $(MAIN)

.PHONY: clean
clean:
	@echo "==> Cleaning build files"
	@rm -f $(APP_NAME)
	@rm -f *.tar.gz
