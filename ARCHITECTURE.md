# CoolAid Backend Architecture Documentation

## Overview
This document describes the **high‑level architecture** of the `coolaid-backend` service, where each major component lives in the repository, and how the pieces are wired together at runtime.

---

## Table of Contents
1. [Entry Point](#entry-point)
2. [Configuration](#configuration)
3. [Infrastructure Layer](#infrastructure-layer)
4. [Dependency‑Injection Container](#dependency-injection-container)
5. [Domain Services](#domain-services)
6. [API Layer (HTTP & GraphQL)](#api-layer)
7. [Middleware](#middleware)
8. [Cron Jobs & Background Workers](#cron-jobs)
9. [Observability & Graceful Shutdown](#observability)
10. [Directory Map](#directory-map)
11. [Future Improvements](#future-improvements)

---

## 1. Entry Point <a name="entry-point"></a>
- **File:** `main.go`
- **Location:** `d:/Replication/coolaid-backend-model/main.go`
- **Purpose:** Calls `app.StartApplication()` which boots the whole system.

---

## 2. Configuration <a name="configuration"></a>
- **Package:** `config`
- **Key Files:**
  - `config/app.go` – holds the `App` struct (port, DB DSN, etc.)
  - `config/auth.go`, `config/db.go`, `config/search_engine.go`, … – grouped config sections.
- **How it works:** `config.LoadConfigs()` reads `.env` (or `.env.prod`) and populates a global `config.App` struct used throughout the code.

---

## 3. Infrastructure Layer <a name="infrastructure-layer"></a>
These packages provide concrete adapters to external systems.
| Package | Path | Responsibility |
|---------|------|----------------|
| **db** | `d:/Replication/coolaid-backend-model/db` | PostgreSQL connection, migrations, repositories (e.g., `NewAuthRepository`). |
| **typesense** | `d:/Replication/coolaid-backend-model/typesense` | Search‑engine client wrapper. |
| **cron_job** | `d:/Replication/coolaid-backend-model/cron_job` | Scheduler (`cronjob.RegisterAll`). |
| **middleware** | `d:/Replication/coolaid-backend-model/middleware` | CORS, User‑Agent, request tracing, etc. |
| **oplog** | `d:/Replication/coolaid-backend-model/oplog` | Structured logging used across the service. |
| **mailer**, **whatsapp**, **smtp** (under `config/` and `utils/`) | Various | Email / WhatsApp notification helpers. |

---

## 4. Dependency‑Injection Container <a name="dependency-injection-container"></a>
- **File:** `internal/service/container/graphql.go`
- **Path:** `d:/Replication/coolaid-backend-model/internal/service/container/graphql.go`
- **Exports:**
  - `type Container struct { DB *db.DBContext; Typesense *typesense.TypesenseContext; Auth *auth.Service }`
  - `func NewContainer(dbCtx *db.DBContext, ts *typesense.TypesenseContext) *Container`
- **What it does:**
  1. Builds infra adapters (`authRepo`, `cryptoSvc`, `mailerSvc`).
  2. Constructs domain services (`auth.NewService`).
  3. Returns a container holding both infra contexts and domain services for the rest of the app.

---

## 5. Domain Services <a name="domain-services"></a>
Located under `internal/service/`.
- **Auth Service** – `internal/service/auth/`
- **Crypto Service** – `internal/service/crypto/`
- **Mailer Service** – `internal/service/mailer/`
- **Ticket Service** – `internal/service/ticket/`
- **Readme:** `internal/service/Readme.md` provides a brief description of each sub‑package.
- **How they are used:** The container injects these services into the GraphQL resolvers and cron jobs.

---

## 6. API Layer (HTTP & GraphQL) <a name="api-layer"></a>
### HTTP (Gin)
- **File:** `app/application.go`
- **Path:** `d:/Replication/coolaid-backend-model/app/application.go`
- **Key responsibilities:**
  - Initialise Gin router (`router = gin.Default()`).
  - Register global middleware (`registerMiddleware`).
  - Call `mapUrls` to mount the GraphQL endpoint (`/graphql`) and Playground.

### GraphQL
- **Package:** `internal/api/graphql`
- **Key files:**
  - `handler.go` – HTTP handler that forwards to the GraphQL server.
  - `server.go` – Creates a `*graphql.Server` using the DI container.
  - `resolver/` – Resolver implementations that call domain services.
  - `playground.go` – Serves the GraphQL Playground UI.

---

## 7. Middleware <a name="middleware"></a>
- **File:** `app/application.go` (lines 91‑105)
- **Implemented middleware:**
  - `middleware.CORSMiddleware` – CORS configuration (origins listed in code).
  - `middleware.UserAgentMiddleware` – Logs the user‑agent header.
- **Location of implementations:** `d:/Replication/coolaid-backend-model/middleware`.

---

## 8. Cron Jobs & Background Workers <a name="cron-jobs"></a>
- **File:** `app/application.go` – `connectCronJobs()` (lines 114‑123)
- **Workflow:**
  1. `ticketservice.NewService` creates a ticket service.
  2. `cronjob.RegisterAll(ticketSvc)` registers all scheduled jobs.
  3. `scheduler.Start()` runs them in a separate goroutine.
- **Source code:** `d:/Replication/coolaid-backend-model/cron_job`.

---

## 9. Observability & Graceful Shutdown <a name="observability"></a>
- **Logging:** All logs go through `oplog` (e.g., `oplog.Info`, `oplog.Error`).
- **Shutdown:** `os.Interrupt` signal handling (lines 74‑88) creates a 10‑second timeout context, calls `server.Shutdown`, and logs the result.

---

## 10. Directory Map <a name="directory-map"></a>
```
.
├─ .dockerignore
├─ .env, .env.example, .env.prod
├─ Dockerfile
├─ go.mod, go.sum
├─ gqlgen.yml
├─ main.go                     ← entry point
├─ app/
│   ├─ application.go          ← server bootstrap, router, middleware
│   └─ url_mappings.go         ← (if present) route definitions
├─ config/
│   ├─ app.go
│   ├─ auth.go
│   ├─ db.go
│   └─ …                       ← configuration structs
├─ internal/
│   ├─ api/
│   │   ├─ graphql/
│   │   │   ├─ handler.go
│   │   │   ├─ playground.go
│   │   │   ├─ resolver/…
│   │   │   └─ server.go
│   │   └─ http/…               ← (future REST handlers)
│   ├─ service/
│   │   ├─ auth/
│   │   ├─ crypto/
│   │   ├─ mailer/
│   │   ├─ ticket/
│   │   └─ container/graphql.go   ← DI container
│   ├─ assets/, csv_util/, crypto/, … (other utilities)
│   └─ validation/…
├─ middleware/…                ← CORS, UserAgent, etc.
├─ db/…                        ← DB context, migrations, repositories
├─ typesense/…                 ← Search engine client
├─ cron_job/…                  ← Scheduler & job definitions
└─ oplog/…                     ← Structured logging
```

---

## 11. Future Improvements <a name="future-improvements"></a>
- Replace the global `router` variable with a struct‑based server.
- Introduce interfaces for repositories and services to improve testability.
- Add a health‑check endpoint (`/healthz`).
- Move CORS origins and other environment‑specific values to the `config` package.
- Write unit tests for services and GraphQL resolvers.
- Consider a code‑generation DI framework (e.g., `wire`) if the container grows.

---

*Document generated on 2026‑01‑24.*
