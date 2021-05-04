package formatters

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

type GoFmt struct {
}

func NewGoFmt() *GoFmt {
	return &GoFmt{}
}

func (g *GoFmt) Format(filename string) error {
	err := exec.Command(
		"go",
		"fmt",
		filepath.Join(filename, "generated"),
	).Run()
	if err != nil {
		return fmt.Errorf("error running go fmt %s: %w", filename, err)
	}

	return nil
}
