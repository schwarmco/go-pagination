# go-pagination

[![Build Status](https://travis-ci.org/schwarmco/go-pagination.svg?branch=master)](https://travis-ci.org/schwarmco/go-pagination)
[![GoDoc](https://godoc.org/github.com/schwarmco/go-pagination?status.svg)](https://godoc.org/github.com/schwarmco/go-pagination)

a package for calculating values needed for pagination in golang

## Installation

In order to start, `go get` this repository:

```
go get github.com/schwarmco/go-pagination
```

## Usage

```go
import (
    "fmt"
    "github.com/schwarmco/go-pagination"
)

items := []string{"apple", "cherry", "pear", "coconut", "banana"}

// assume, we are on page 1 and want 3 items listed on each page
pager := pagination.Calculate(1, 3, len(items))

fmt.Println(pager.NumPages) // Output: 2
fmt.Println(pager.HasNext) // Output: true
fmt.Println(pager.HasPrev) // Output: false
```

## Todo

* implement pagination range pattern (like: 1 2 3 .. 7 8 9 ... 12 13 14)
