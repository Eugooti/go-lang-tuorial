# Golang Basics: Packages and Modules

## Overview

Packages and modules are fundamental concepts in Go that allow developers to organize code, reuse functionality, and manage dependencies effectively. Understanding how to create, import, and use packages and modules is essential for writing modular and maintainable Go applications.

## 1. Packages

- **Definition**: A package is a collection of related Go files in the same directory that provides specific functionality. Every Go file belongs to a package, and each Go program is made up of one or more packages.
- **Types**:
    - **Main Package**: A special package that defines a standalone executable program. It must contain a `main()` function.
    - **Library Packages**: Packages that provide reusable functionality. They do not have a `main()` function.
- **Naming Conventions**: The package name is usually the same as the last element of the directory path. For example, files in the `mathutil` directory will have `package mathutil`.

## 2. Modules

- **Definition**: A module is a collection of related Go packages that are versioned together as a unit. Modules are how Go manages dependencies.
- **`go.mod` File**: The `go.mod` file defines the module's path and its dependencies. It is created using the `go mod init <module-name>` command.
- **Example**:
  ```bash
  go mod init example.com/myapp
