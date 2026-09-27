# Go Handy Notes

> A lightweight reference for Go concepts, terminology, and commands learned during development.
> This is intentionally separate from the project's main `README.md`.

---

## 1. Multiplexer (`mux`)

### What is a multiplexer?

In general, a **multiplexer** (or **mux**) is something that allows multiple signals or inputs to share a common channel.

In Go web development, **mux** usually refers to an **HTTP request multiplexer**.

### HTTP request multiplexer

An HTTP mux matches an incoming HTTP request against a set of registered routes and dispatches the request to the handler associated with the matching route.

Simply put:

> **A mux is the component that decides which handler should handle an incoming HTTP request.**

For example:

```text
GET /users
       │
       ▼
     Mux
       │
       ▼
UsersHandler
```

The mux is therefore an important entry point into an HTTP application.

---

## 2. Go Modules

A **Go module** is a collection of related Go packages that are versioned and distributed together. A module is defined by a `go.mod` file at its root.

A Go module helps Go:

- identify the module's root
- resolve packages within the module
- manage the module's dependencies
- record the module's module path and required Go version

### Initialize a module

For example:

```bash
go mod init github.com/goodness/go-mastery
```

This creates a `go.mod` file and establishes:

```text
Module path: github.com/goodness/go-mastery
Module root: directory containing go.mod
```

**Important:** the module path is an identifier used for importing and resolving packages. It does **not** mean that the source code must physically be located in a directory such as:

```text
github.com/goodness/go-mastery
```

The module root is simply the directory containing the `go.mod` file.

Example:

```text
go-mastery/
├── go.mod          ← module root
├── main.go
├── internal/
└── api/
```

---

## 3. Adding External Dependencies

A Go project can use packages from outside its own module.

For example, after importing a package:

```go
import "github.com/gorilla/mux"
```

you can add the dependency with:

```bash
go get github.com/gorilla/mux
```

The `go` command updates the module's dependency information in `go.mod` and may update `go.sum`.

### `go.mod`

`go.mod` describes the module and its dependencies, including the versions required by the module.

Example:

```go
module github.com/goodness/go-mastery

go 1.XX

require (
    github.com/gorilla/mux v1.x.x
)
```

### `go.sum`

`go.sum` contains cryptographic checksums for module versions that Go has downloaded or otherwise recorded.

It helps Go verify that downloaded module content matches the expected content.

> **Important:** `go.sum` is not a dependency lock file in the same sense as a typical JVM lock file. Dependency requirements and versions are declared in `go.mod`; `go.sum` primarily provides integrity checksums.

---

## Quick Reference

| Concept | Meaning |
|---|---|
| **mux** | Matches HTTP requests to registered routes/handlers |
| **module** | A versioned collection of Go packages |
| **`go.mod`** | Defines the module and its dependency requirements |
| **module root** | Directory containing `go.mod` |
| **`go get`** | Adds or updates module dependencies |
| **`go.sum`** | Stores checksums used to verify module content |

---

## Commands

```bash
# Initialize a new module
go mod init <module-path>

# Add or update a dependency
go get <module-path>

# Synchronize module dependencies
go mod tidy
```

> Keep this file focused on **definitions, mental models, terminology, and useful commands**. Put project-specific architecture and setup instructions in the main `README.md`.
