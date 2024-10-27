# Go compiler
GO := go

# Directory where the source files are located
SRCDIR := cmd

# Output directory for built binaries
OUTDIR := bin

# List of source files (assuming all .go files in the src directory)
SOURCES := $(wildcard $(SRCDIR)/*.go)

# Name of the final executable (you can change this to your desired name)
TARGET := server 

# Build command
build:
	@echo "Building $(TARGET)..."
	$(GO) build -o $(OUTDIR)/$(TARGET) $(SOURCES)

# Run command (depends on build)
run: build
	@echo "Running $(TARGET)..."
	$(OUTDIR)/$(TARGET)

# Clean command
clean:
	@echo "Cleaning up..."
	rm -rf $(OUTDIR)

# Migration command to create a new migration
migration:
	@echo "Creating migration..."
	@migrate create -ext sql -dir $(SRCDIR)/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

# Migrate up command
migrate-up:
	@echo "Applying migrations up..."
	$(GO) run $(SRCDIR)/migrate/main.go up

# Migrate down command
migrate-down:
	@echo "Applying migrations down..."
	$(GO) run $(SRCDIR)/migrate/main.go down

# Code generation command
codegen: crud

SCHEMA_PACKAGE := schemas

crud:
	@echo "Running code generation..."
	$(GO) run codegen/generate_crud.go crud $(SCHEMA_PACKAGE) $(SCHEMA_NAME)

templ:
	@echo "Running templ generate"
	@templ generate

# Help command (to show available commands)
help:
	@echo "Available commands:"
	@echo "  make build       - Build the project"
	@echo "  make run         - Build and run the project"
	@echo "  make clean       - Clean up built binaries"
	@echo "  make migration   - Create a new migration"
	@echo "  make migrate-up  - Apply migrations up"
	@echo "  make migrate-down- Apply migrations down"
	@echo "  make codegen
