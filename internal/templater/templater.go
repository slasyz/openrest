package templater

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Templater struct {
}

func New() *Templater {
	return &Templater{}
}

func (t *Templater) Write(outFile string, tpl *template.Template, data interface{}) (bool, error) {
	err := os.MkdirAll(filepath.Dir(outFile), 0744)
	if err != nil {
		return false, fmt.Errorf("cannot create output directory %s: %w", filepath.Dir(outFile), err)
	}

	f, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_EXCL, 0644)
	if err != nil {
		if os.IsExist(err) {
			return false, nil
		}

		return false, fmt.Errorf("cannot create output file %s: %w", outFile, err)
	}
	defer f.Close()

	err = tpl.Execute(f, data)
	if err != nil {
		return false, fmt.Errorf("cannot execute template for %s: %w", outFile, err)
	}

	return true, nil
}
