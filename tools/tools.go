// +build tools

//go:generate go build -o . golang.org/x/tools/cmd/goimports

package tools

import (
	_ "golang.org/x/tools/cmd/goimports"
)
