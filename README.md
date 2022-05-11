## Go Boiler Plate Service

Initialize project for Go Boiler Plate Service

## Prerequisites

1. Go 1.16
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