// +build tools

// This file used for attaching tools dependencies to the project
package main

import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
