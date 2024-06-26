# Halo Suster

[Halo Suster](https://openidea-projectsprint.notion.site/HaloSuster-be1866776fe84c2d8d9eac08ce09b7a5) is a backend system designed for nurses to record patient medical records.

## 📜Table of Contents

- [Halo Suster](#halo-suster)
  - [📜Table of Contents](#table-of-contents)
  - [🔍Requirements](#requirements)
  - [🛠️Installation](#installation)
  - [🌟Features](#features)
  - [🚀Usage](#usage)
  - [⚙️Configuration](#configuration)
  - [💾Database Migration](#database-migration)
  - [🔨Build](#build)
  - [🧪Testing](#testing)

## 🔍Requirements

This application requires the following:

- [Golang](https://golang.org/dl/)
- [Git](https://git-scm.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Golang Migrate](https://github.com/golang-migrate/migrate)
- [AWS CLI](https://aws.amazon.com/cli/)

## 🛠️Installation

To install the boilerplate, follow these steps:

1. Make sure you have [Golang](https://golang.org/dl/), [Git](https://git-scm.com/), [PostgreSQL](https://www.postgresql.org/), [Golang Migrate](https://github.com/golang-migrate/migrate), and [AWS CLI](https://aws.amazon.com/cli/) installed and configured on your system.

2. Clone this repository:

   ```bash
   git clone https://github.com/malikfajr/halo-suster.git
   ```

3. Navigate to the project directory:

   ```bash
   cd halo-suster
   ```

4. Download the required dependencies:

   ```bash
   go mod tidy
   ```

5. Configure AWS CLI:

   ```bash
   aws configure
   ```

## 🌟Features

- Authentication:
- Nurse Management
- Manage Medical Record
- Image upload

  
## 🚀Usage

1. **Setting Up Environment Variables**

   Before starting the application, you need to set up the following environment variables:

   ```bash
   export DB_NAME=           # Name of the PostgreSQL database
   export DB_PORT=           # Port of the PostgreSQL database (default: 5432)
   export DB_HOST=           # Hostname or IP address of the PostgreSQL server
   export DB_USERNAME=       # Username for the PostgreSQL database
   export DB_PASSWORD=       # Password for the PostgreSQL database
   export DB_PARAMS=         # Additional connection parameters for PostgreSQL (e.g., sslmode=disable)
   export JWT_SECRET=        # Secret key used for generating JSON Web Tokens (JWT)
   export BCRYPT_SALT=       # Salt for password hashing (use a higher value than 8 in production!)
   
   # AWS S3 Configuration
   export AWS_ACCESS_KEY_ID=         # AWS Access Key ID for S3 bucket access
   export AWS_SECRET_ACCESS_KEY=     # AWS Secret Access Key for S3 bucket access
   export AWS_S3_BUCKET_NAME=        # Name of the AWS S3 bucket for file storage
   export AWS_REGION=                # AWS region where the S3 bucket is located
   ```


2. **Running the Application**

   Once you have set up the environment variables, you can start the application by running:

   ```bash
   go run main.go
   ```

   This will start the Halo Suster application on the default port (usually 8080). You can access the application in your web browser at `http://localhost:8080`.

## ⚙️Configuration

The application uses environment variables for configuration. You can configure the database connection, JWT secret, aws config, and bcrypt salt by setting the following environment variables:

- Refer to the [Usage](#usage) section for a detailed explanation of each environment variable.

## 💾Database Migration

Database migration must use golang-migrate as a tool to manage database migration

1. Direct your terminal to your project folder first

2. Initiate folder

   ```bash
   mkdir db/migrations
   ```

3. Create migration

   ```bash
   migrate create -ext sql -dir db/migrations add_user_table
   ```

   This command will create two new files named `add_user_table.up.sql` and `add_user_table.down.sql` inside the `db/migrations` folder

   - `.up.sql` can be filled with database queries to create / delete / change the table
   - `.down.sql` can be filled with database queries to perform a `rollback` or return to the state before the table from `.up.sql` was created

4. Execute migration

   ```bash
   migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
   ```

5. Rollback migration

   ```bash
   migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations down
   ```

6. View the current migration state

   ```bash
   migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" version 
   ```

## 🔨Build

To build app for different operating systems and architectures, you can use the following commands:

1. **Windows (amd64)**:

    ```bash
    GOOS=windows GOARCH=amd64 go build -o build/main.exe main.go
    ```

2. **Linux (amd64)**:

    ```bash
    GOOS=linux GOARCH=amd64 go build -o build/main main.go
    ```

3. **macOS (amd64)**:

    ```bash
    GOOS=darwin GOARCH=amd64 go build -o build/main main.go
    ```

4. **Linux (ARM)**:
    ```bash
    GOOS=linux GOARCH=arm go build -o build/main main.go
    ```

## 🧪Testing

To test the Cats Social API, you can use the testing [repository](https://github.com/nandanugg/HaloSusterTestCasesPSW3B2) provided.
