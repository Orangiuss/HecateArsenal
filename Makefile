# Makefile for HecateArsenal

# Compiler options
GO := go
GOFLAGS :=

# Source file
SRC := src/main.go

# Output binary
BIN := bin/hecatearsenal

# Default target
all: $(BIN)

# Compile the source file
$(BIN): $(SRC)
	$(GO) build $(GOFLAGS) -o $(BIN) $(SRC)

# Clean the generated binary
clean:
	rm -f $(BIN)