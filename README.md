YohGo database
===

YohGo database manages the database connection of an application so that the database connection is wrapped in a configurable container that can be utilised in various packages that require a connection to a database.

---------------------------------------

  * [Requirements](#requirements)
  * [Features](#features)
  * [Installation](#installation)
  * [Usage](#usage)
    * [Password Grant](#get-an-access-token-using-the-password-grant-type)
    * [Refresh Token Grant](#refresh-the-token-using-the-refresh-token-grant-type)

---------------------------------------

## Requirements
   * Go (1.5+)
   * MySQL (4.1+)

---------------------------------------

## Features

  * Lightweight and fast
  * Native Go implementation
  * Currently only supports `mysql` and `gorm mysql`

---------------------------------------

## Installation

Simply install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:

```bash
$ go get github.com/yohgo/database
```

Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

---------------------------------------

## Usage

_YohGo Database_ is a Go database container that. You only need to import the YohGo database package of choice and can use the full [`database/sql`](https://golang.org/pkg/database/sql/) API then.

Example of using the `gorm` package:
```go

import (
    "github.com/yohgo/database/gorm"
)

// Configure database connection
database, err := gorm.NewDatabase("DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME")
// Check if connection failed
if err != nil {
    log.Fatal("Failed to connect to database")
}
// Always ensure to close the connection
defer database.Close()
```

Example of using the `mysql` package:
```go

import (
    "github.com/yohgo/database/mysql"
)

// Configure database connection
database, err := mysql.NewDatabase("DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME")
// Check if connection failed
if err != nil {
    log.Fatal("Failed to connect to database")
}
// Always ensure to close the connection
defer database.Close()
```
