package formatters

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

type GoImports struct {
	cmd string
}

func NewGoImports(cmd string) *GoImports {
	return &GoImports{
		cmd: cmd,
	}
}

func (g *GoImports) Format(filename string) error {
	out, err := exec.Command(
		g.cmd,
		filename,
	).Output()
	if err != nil {
		return fmt.Errorf("error running %s %s: %w", g.cmd, filename, err)
	}

	err = ioutil.WriteFile(filename, out, 0644)
	if err != nil {
		return fmt.Errorf("error writing goimports out to %s: %w", filename, err)
	}

	return nil
}
