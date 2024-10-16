# Variables
GO ?= go

CMD_DIR = cmd
BIN_DIR = bin

MAIN = ${CMD_DIR}/${SERVICE}/main.go
OUTPUT = ${BIN_DIR}/${SERVICE}

# Check for SERVICE variable
ifeq ($(SERVICE),)
  $(error SERVICE variable is not set. Please run 'make <target> SERVICE=<service_name>')
endif

# Targets
.PHONY: all build clean run

all: build

build: | $(BIN_DIR)
	$(GO) build -o $(OUTPUT) $(MAIN)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

run:
	./$(OUTPUT)

clean:
	rm -f $(OUTPUT)
