## Go Boiler Plate Service

Initialize project for Go Boiler Plate Service
this boilerplate is made based on [Golang Standard Project Layout](https://github.com/golang-standards/project-layout)

## Prerequisites

1. Go 1.17
2. PostgreSQL 11
3. UNIX Shell
   > Use `wsl` in Windows 10
4. Git
5. Make
6. Docker CE (Optional)

## Set-up

1. Configure Project

   ```sh
   # Run scripts to make env from .env-example and grant permission
   make setup
   
   # Run scripts to get all app dependencies that needed.
   make configure
   
   # Run scripts to check all prerequisites for development is available
   make doctor
   ```

2. Configure project. see [Configuration Section](#Configuration) for details:

3. Init Database

   Once `.env` has been configured, initiate database:
   ```bash
   # Create database if not exists
   make db
   
   # Upgrade database to next version
   make db-up
   ```

## Configuration

   PDS service are configurable from `.env` file

### Run Development

   ```sh
   # Run Service
   make serve
   ```

### Make command available
   ```sh
   # Run make help

   Choose a command run in Go Boilerplate Service:

   help                 Show command help
   clean                Clean everything
   doctor               Check for prerequisites
   setup                Make env from env example and grant permission.
   configure            Configure project
   configure-swag       Configure install swag app
   configure-ginkgo     Configure install ginkgo app
   configure-mockgen    Configure install mockgen app
   configure-golangci   Configure install golangci app
   configure-with-cli   Configure project with cli apps dependent
   lint                 Run golang linter
   testing              Run automation test with ginkgo
   serve                Run server in development mode
   serve-with-doc       Run server in development mode with the swagger doc
   release-dev          Run server in development mode with the swagger doc
   vendor               Download dependencies to vendor folder
   release              Compile binary for deployment.
   image                Build a docker image from release
   image-push           Push app image
   docker-serve         Run application with docker compose
   db-configure         Generate a configuration for database migration tool
   db-status            Prints the details and status information about all the migration.
   db-generate          create a new migration file version
   db-up                Upgrade database
   db-down              (Experimental) undo to previous migration version
   db-clean             Clean database

   ```