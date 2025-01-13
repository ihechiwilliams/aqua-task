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

- Set up and populate.env file:

    ```bash
      make env
    ```

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
- Start the  notification service
  ```bash
    make start-notification
  ```

## HTTP API Endpoints

API documentation for the Rest endpoints can be found inside openapi folder, then you will see the openapi.yml file

## gRPC API

The gRPC server provides the following methods:

- `DeleteNotification`: Deletes a notification by ID.
- `DeleteNotificationByUser`: Deletes all notifications for a user.
- `GetNotificationsByUser`: Gets all notifications for a user.