# CoolAid Backend Folder Structure Documentation

This document provides an overview of the folder structure of the `coolaid-backend` project, describing the purpose of each folder and its key files.

---

## Root Directory
- **ARCHITECTURE.md**: High-level architecture documentation for the backend.
- **Dockerfile**: Docker configuration for containerizing the application.
- **go.mod**: Go module dependencies.
- **gqlgen.yml**: Configuration for generating GraphQL code.
- **main.go**: Entry point of the application. Calls `app.StartApplication()` to initialize the system.
- **sqlc.yaml**: Configuration for SQLC code generation.

---

## app/
- **application.go**: Server bootstrap, initializes the Gin router, middleware, and routes.
- **url_mappings.go**: Defines route mappings for the application.

---

## config/
- **app.go**: Defines the `App` struct for application-level configurations (e.g., port, DB connection string).
- **auth.go**: Configuration for authentication.
- **db.go**: Database configuration.
- **search_engine.go**: Configuration for the search engine.
- **smtp.go**: Configuration for SMTP (email).
- **whatsapp.go**: Configuration for WhatsApp integration.

---

## cron_job/
- **cron_job.go**: Main file for defining and managing cron jobs.
- **registry.go**: Registers all cron jobs.
- **ticket_cleanup.go**: Defines a specific cron job for cleaning up tickets.

---

## db/
- **auth_repo.go**: Repository for authentication-related database operations.
- **connect.go**: Handles database connection setup.

---

## internal/
### api/
#### graphql/
- **handler.go**: HTTP handler for GraphQL requests.
- **playground.go**: Serves the GraphQL Playground UI.
- **server.go**: Creates a `*graphql.Server` using the dependency injection container.
- **resolver/**: Contains resolver implementations that interact with domain services.

#### http/
- **handler/**: Placeholder for future REST API handlers.

### assets/
#### mail/
- **common.go**: Common utilities for email templates.
- **inventory_reorder.go**: Email template for inventory reorder notifications.
- **otp_forgot_password.go**: Email template for OTP and password recovery.

### csv_schema/
- **basic.go**: Basic CSV schema definitions.
- **inventory.go**: CSV schema for inventory data.
- **pricing.go**: CSV schema for pricing data.
- **registry.go**: Registry for CSV schemas.
- **categories/**: Contains CSV schemas for various product categories (e.g., `actuator.go`, `blower_motor.go`, etc.).

### csv_util/
- **buildCSV.go**: Utility for building CSV files.
- **error.go**: Error handling for CSV operations.
- **export.go**: Handles CSV export functionality.
- **header.go**: Manages CSV headers.
- **mapping.go**: Maps data to CSV format.
- **reader.go**: Reads CSV files.
- **rebuild.go**: Rebuilds CSV data.
- **rows.go**: Handles CSV rows.
- **write.go**: Writes data to CSV files.

### generated/
#### graphql/
- Contains auto-generated GraphQL files (e.g., `auth.go`, `brand.go`, `category.go`, etc.).

#### sqlc/
- Contains auto-generated SQLC files (e.g., `actuator.sql.go`, etc.).

### importer/
- Placeholder for data import utilities.

### repository/
- **basic.go**: Basic repository functions.
- **inventory.go**: Repository for inventory-related operations.
- **pricing.go**: Repository for pricing-related operations.
- **category/**: Contains repositories for various categories.

### service/
- **auth/**: Authentication service.
- **crypto/**: Cryptographic utilities and services.
- **mailer/**: Email sending service.
- **ticket/**: Ticket management service.
- **container/graphql.go**: Dependency injection container for services and infrastructure.
- **Readme.md**: Overview of the services provided by this package.

### validation/
- **engine.go**: Validation engine.
- **headers.go**: Header validation utilities.
- **README.md**: Documentation for the validation package.
- **types.go**: Validation-related types.
- **validators.go**: Custom validators.

---

## middleware/
- **cors.go**: Middleware for handling CORS.
- **user_agent.go**: Middleware for logging user-agent headers.

---

## migrations/
### db/
- **queries/**: SQL queries for database migrations.
- **schema/**: Database schema definitions.

### graphql/
- **schema/**: GraphQL schema definitions.

---

## oplog/
- **log.go**: Structured logging utilities.

---

## pkg/
### crypto/
- **jwt.go**: JWT-related cryptographic utilities.
- **opt.go**: Cryptographic options.
- **password.go**: Password hashing and validation.
- **token.go**: Token generation and validation.

### mailer/
- **send.go**: Email sending utilities.

---

## utils/
- **jwt.go**: JWT-related utility functions.

---

This document provides a comprehensive overview of the folder structure and the purpose of each file and directory in the `coolaid-backend` project.