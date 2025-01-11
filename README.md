# Aqua Security Assessment


## Cloud Resource API

Cloud Resource Inventory Management System [Go](https://golang.org)

### Prerequisites

Before diving in, make sure you have:

- A Unix-like OS (Linux/macOS)
- Docker & Docker Compose
- Go 1.23+
- make is installed

### Quick Start

Ready to launch? Follow these steps:

- Build app, Sync your local database with the schema and run migration:

    ```bash
      make build-app
    ```

- Run the migration

  ```bash
    make migrate
  ```

- Seed the database and start the server
  ```bash
    make seed
  ```

- Start the server without seeding
  ```bash
    make start
  ```
  