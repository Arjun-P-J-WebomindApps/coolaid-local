# CoolAid Backend Full Overview

This document provides a comprehensive overview of the CoolAid backend project, including its folder structure, key components, and their responsibilities. The goal is to document all aspects of the architecture and implementation for better understanding and maintainability.

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
- **connect.go**: Handles database connection setup and provides a global `DBContext` for queries.

---

## internal/
### api/
#### graphql/
- **handler.go**: HTTP handler for GraphQL requests. Sets up context and serves requests using the GraphQL server.
- **playground.go**: Serves the GraphQL Playground UI.

---

### domain/
#### product/
- **models.go**: Defines core domain models for products, including `Product`, `ProductDetails`, and `ModelVariant`.

---

## middleware/
- **cors.go**: Implements CORS middleware for handling cross-origin requests.

---

## utils/
- **jwt.go**: Handles JWT operations, including token generation and validation.

---

## internal/repository/product/
- **offer.go**: Manages database operations related to product offers, including fetching, creating, and mapping offers.

---

## internal/
### csv_schema/
- **basic.go**: Contains basic schema definitions for CSV processing.
- **categories/**: Contains schema definitions for various product categories, such as:
  - `actuator.go`, `blower_motor.go`, `cabin_filter.go`, `chiller_unit.go`, `clutch_assy.go`, etc.
- **inventory.go**: Schema definitions for inventory-related CSV data.
- **pricing.go**: Schema definitions for pricing-related CSV data.
- **registry.go**: Schema registry for managing CSV schemas.

### csv_util/
- **buildCSV.go**: Utility for building CSV files.
- **error.go**: Handles errors related to CSV operations.
- **export.go**: Handles exporting data to CSV format.
- **header.go**: Manages CSV headers.
- **mapping.go**: Handles mapping of data to CSV format.
- **reader.go**: Reads data from CSV files.
- **rebuild.go**: Rebuilds CSV files.
- **rows.go**: Manages rows in CSV files.
- **write.go**: Handles writing data to CSV files.

### generated/
#### sqlc/
- Contains SQLC-generated code for database queries and models, including:
  - `actuator.sql.go`, `blower_motors.sql.go`, `brand.sql.go`, `categories.sql.go`, etc.
  - `db.go`: Centralized file for SQLC database queries.
  - `models.go`: SQLC-generated models for database tables.

### repository/
#### auth/
- **adaptor.go**: Handles authentication-related database operations.
- **mapping.go**: Maps database rows to domain models for authentication.
- **otp.go**: Manages OTP-related database operations.
- **refreshToken.go**: Handles refresh token operations.
- **session.go**: Manages user sessions in the database.
- **toptp.go**: Handles TOTP (Time-based One-Time Password) operations.
- **user.go**: Manages user-related database operations.

#### csv/
- **basic.go**: Handles basic CSV-related repository operations.
- **category/**: Manages category-related CSV repository operations.
- **inventory.go**: Handles inventory-related CSV repository operations.
- **pricing.go**: Manages pricing-related CSV repository operations.

---

## internal/
### domain/
#### auth/
- **db_models.go**: Defines database models for authentication.
- **db_params.go**: Defines database query parameters for authentication.
- **errors.go**: Contains error definitions for authentication.
- **inputs.go**: Defines input structures for authentication operations.
- **interfaces.go**: Defines interfaces for authentication services.
- **login.go**: Handles user login logic.
- **models.go**: Defines core domain models for authentication.
- **password.go**: Manages password-related operations (e.g., hashing, validation).
- **refresh.go**: Handles token refresh logic.
- **register.go**: Manages user registration logic.
- **service.go**: Implements authentication services.
- **session.go**: Manages user sessions.
- **toptp.go**: Handles TOTP (Time-based One-Time Password) operations.

#### master/
##### model/
- **create.go**: Handles creation of model-related data.
- **db_model.go**: Defines database models for models.
- **db_params.go**: Defines database query parameters for models.
- **delete.go**: Handles deletion of model-related data.
- **errors.go**: Contains error definitions for models.
- **inputs.go**: Defines input structures for model operations.
- **interfaces.go**: Defines interfaces for model services.
- **mapper.go**: Maps database rows to domain models for models.
- **models.go**: Defines core domain models for models.
- **queries.go**: Contains database queries for models.
- **service.go**: Implements model-related services.
- **update.go**: Handles updates to model-related data.

##### vendors/
- Similar structure to `model/` with files for managing vendor-related data.

---

### repository/
#### master/
- **brand/**: Repository for brand-related database operations.
- **category/**: Repository for category-related database operations.
- **company/**: Repository for company-related database operations.
- **customer/**: Repository for customer-related database operations.
- **model/**: Repository for model-related database operations.
- **vendor/**: Repository for vendor-related database operations.

#### techspec/
- Contains repository files for technical specifications, including:
  - `actuator.go`, `adapter.go`, `blower_motor.go`, `cabin_filter.go`, `chiller_unit.go`, etc.

---

## internal/
### domain/
#### master/
##### model/
- **create.go**: Handles creation of model-related data.
- **db_model.go**: Defines database models for models.
- **db_params.go**: Defines database query parameters for models.
- **delete.go**: Handles deletion of model-related data.
- **errors.go**: Contains error definitions for models.
- **inputs.go**: Defines input structures for model operations.
- **interfaces.go**: Defines interfaces for model services.
- **mapper.go**: Maps database rows to domain models for models.
- **models.go**: Defines core domain models for models.
- **queries.go**: Contains database queries for models.
- **service.go**: Implements model-related services.
- **update.go**: Handles updates to model-related data.

##### vendors/
- Similar structure to `model/` with files for managing vendor-related data.

---

### repository/
#### product/
- **adaptor.go**: Handles product-related database operations.
- **inventory.go**: Manages inventory-related database operations.
- **offer.go**: Manages database operations related to product offers.
- **pricing.go**: Handles pricing-related database operations.
- **product.go**: Manages core product-related database operations.
- **variants.go**: Handles operations related to product variants.

#### csv/
- **basic.go**: Handles basic CSV-related repository operations.
- **category/**: Manages category-related CSV repository operations.
- **inventory.go**: Handles inventory-related CSV repository operations.
- **pricing.go**: Manages pricing-related CSV repository operations.

---

## middleware/
- **cors.go**: Implements CORS middleware for handling cross-origin requests.

---

## utils/
- **jwt.go**: Handles JWT operations, including token generation and validation.

---

## internal/repository/product/
- **offer.go**: Manages database operations related to product offers, including fetching, creating, and mapping offers.

---

## internal/
### csv_schema/
- **basic.go**: Contains basic schema definitions for CSV processing.
- **categories/**: Contains schema definitions for various product categories, such as:
  - `actuator.go`, `blower_motor.go`, `cabin_filter.go`, `chiller_unit.go`, `clutch_assy.go`, etc.
- **inventory.go**: Schema definitions for inventory-related CSV data.
- **pricing.go**: Schema definitions for pricing-related CSV data.
- **registry.go**: Schema registry for managing CSV schemas.

### csv_util/
- **buildCSV.go**: Utility for building CSV files.
- **error.go**: Handles errors related to CSV operations.
- **export.go**: Handles exporting data to CSV format.
- **header.go**: Manages CSV headers.
- **mapping.go**: Handles mapping of data to CSV format.
- **reader.go**: Reads data from CSV files.
- **rebuild.go**: Rebuilds CSV files.
- **rows.go**: Manages rows in CSV files.
- **write.go**: Handles writing data to CSV files.

### generated/
#### sqlc/
- Contains SQLC-generated code for database queries and models, including:
  - `actuator.sql.go`, `blower_motors.sql.go`, `brand.sql.go`, `categories.sql.go`, etc.
  - `db.go`: Centralized file for SQLC database queries.
  - `models.go`: SQLC-generated models for database tables.

### repository/
#### auth/
- **adaptor.go**: Handles authentication-related database operations.
- **mapping.go**: Maps database rows to domain models for authentication.
- **otp.go**: Manages OTP-related database operations.
- **refreshToken.go**: Handles refresh token operations.
- **session.go**: Manages user sessions in the database.
- **toptp.go**: Handles TOTP (Time-based One-Time Password) operations.
- **user.go**: Manages user-related database operations.

#### csv/
- **basic.go**: Handles basic CSV-related repository operations.
- **category/**: Manages category-related CSV repository operations.
- **inventory.go**: Handles inventory-related CSV repository operations.
- **pricing.go**: Manages pricing-related CSV repository operations.

---

## internal/
### domain/
#### auth/
- **db_models.go**: Defines database models for authentication.
- **db_params.go**: Defines database query parameters for authentication.
- **errors.go**: Contains error definitions for authentication.
- **inputs.go**: Defines input structures for authentication operations.
- **interfaces.go**: Defines interfaces for authentication services.
- **login.go**: Handles user login logic.
- **models.go**: Defines core domain models for authentication.
- **password.go**: Manages password-related operations (e.g., hashing, validation).
- **refresh.go**: Handles token refresh logic.
- **register.go**: Manages user registration logic.
- **service.go**: Implements authentication services.
- **session.go**: Manages user sessions.
- **toptp.go**: Handles TOTP (Time-based One-Time Password) operations.

#### master/
##### model/
- **create.go**: Handles creation of model-related data.
- **db_model.go**: Defines database models for models.
- **db_params.go**: Defines database query parameters for models.
- **delete.go**: Handles deletion of model-related data.
- **errors.go**: Contains error definitions for models.
- **inputs.go**: Defines input structures for model operations.
- **interfaces.go**: Defines interfaces for model services.
- **mapper.go**: Maps database rows to domain models for models.
- **models.go**: Defines core domain models for models.
- **queries.go**: Contains database queries for models.
- **service.go**: Implements model-related services.
- **update.go**: Handles updates to model-related data.

##### vendors/
- Similar structure to `model/` with files for managing vendor-related data.

---

### repository/
#### master/
- **brand/**: Repository for brand-related database operations.
- **category/**: Repository for category-related database operations.
- **company/**: Repository for company-related database operations.
- **customer/**: Repository for customer-related database operations.
- **model/**: Repository for model-related database operations.
- **vendor/**: Repository for vendor-related database operations.

#### techspec/
- Contains repository files for technical specifications, including:
  - `actuator.go`, `adapter.go`, `blower_motor.go`, `cabin_filter.go`, `chiller_unit.go`, etc.

---

## typesense/
- **connect.go**: Manages the connection to the Typesense search engine. It:
  - Initializes a `TypesenseContext` with a `typesense.Client`.
  - Performs a health check on the Typesense server.
  - Logs the connection status using the `oplog` package.

---

## migrations/
### db/
- **queries/**: Contains SQL queries for database operations.
- **schema/**: Contains database schema files.

---

## Observations and Recommendations

### Strengths
1. **Modular Design:**
   - The project is well-structured with clear separation of concerns.
   - Each folder and file has a specific responsibility, making the codebase easier to navigate and maintain.

2. **Scalability:**
   - The use of interfaces, domain models, and repository patterns ensures the architecture is extensible and scalable.

3. **Third-Party Libraries:**
   - The project leverages robust libraries like `gin`, `robfig/cron`, and `sqlc` to reduce boilerplate code and improve reliability.

4. **Security:**
   - JWT-based authentication is implemented, which is a good practice for stateless authentication.

5. **Background Jobs:**
   - The `cronjob` package is well-structured and uses the `robfig/cron` library for managing background tasks.

---

### Limitations
1. **Error Handling:**
   - Error handling is basic and scattered. A centralized error-handling mechanism would improve consistency and maintainability.

2. **Testing:**
   - There is no visible testing strategy or test files in the workspace. Adding unit and integration tests is crucial for ensuring code quality.

3. **Documentation:**
   - While the architecture is modular, inline comments and function documentation are limited. Adding more comments and documentation would help new developers understand the codebase more easily.

4. **Configuration Management:**
   - Configuration files are spread across multiple files in the `config/` folder. Consolidating them into a single file or struct could simplify management.

5. **Missing Files:**
   - Some files referenced in the folder structure (e.g., `service/container/container.go`, `service/mailer/send.go`) are missing or could not be resolved. This might indicate incomplete implementation or outdated documentation.

---

This document will be updated as more files and folders are analyzed.