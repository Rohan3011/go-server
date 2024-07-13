# GO Backend Template

Strong building blocks for robust and scalable backend.

## Requirements

- Go
- PostgreSQL

## Installation

1. Clone the repository:

```bash
git clone https://github.com/your/repository.git
cd repository
```

2.Install dependencies:

```bash
go mod tidy
```

## Features

- Project Structure
- Migration
- Code Generation
- Services Architecture

## Project Structure

```
root/
├───bin
├───cmd
│   ├───api
│   └───migrate
│       └───migrations
├───codegen # All codegen scripts and templates
├───config
├───db
├───schemas
├───services
│   ├───auth
│   ├───user
│   └───view
├───static
│   ├───css
│   └───js
├───templates
├───tmp
├───types
└───utils

```

## Migration

### Creating a New Migration

To create a new migration file:

```bash
make migration NAME=new_migration_name
```

Replace `new_migration_name` with a descriptive name for your migration.

### Applying Migrations

To apply migrations:

**Up (Apply migrations):**

```bash
make migrate-up
```

**Down (Rollback migrations):**

```bash
make migrate-down
```

## Code Generation

### Generating CRUD Code

Define the schema in `schemas/schema.go` file.

To generate CRUD code for a specific schema, specify `SCHEMA_NAME`:

```bash
make codegen crud SCHEMA_NAME=your_schema_name
```

Replace `your_schema_name` with the name of the schema defined in your schemas package.

## Running the Application

Build and Run
To build and run the application:

```bash
make run
```

This command will compile the project and execute the generated executable.

## Cleaning Up

Clean Built Binaries
To clean up built binaries:

```bash
make clean
```

This command will delete the bin directory.

## Contributing

If you would like to contribute to this project, please follow these steps:

Fork the repository
Create a new branch (git checkout -b feature/your-feature)
Commit your changes (git commit -am 'Add new feature')
Push to the branch (git push origin feature/your-feature)
Create a new Pull Request

## License

This project is licensed under the MIT License.
