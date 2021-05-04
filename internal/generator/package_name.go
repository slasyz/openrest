package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

func generateOutputPackageName(outDir string) (string, error) {
	outDir, err := filepath.Abs(outDir)
	if err != nil {
		return "", fmt.Errorf("cannot get output directory absolute path: %w", err)
	}

	var elements []string
	for {
		if outDir == "/" {
			return "", errNoGoMod
		}

		filename := filepath.Join(outDir, "go.mod")
		data, err := ioutil.ReadFile(filename)
		if os.IsNotExist(err) {
			elements = append([]string{filepath.Base(outDir)}, elements...)
			outDir = filepath.Dir(outDir)
			continue
		}

		goMod, err := modfile.Parse(filename, data, nil)
		if err != nil {
			return "", fmt.Errorf("error parsing go.mod: %w", err)
		}

		return goMod.Module.Mod.Path + "/" + strings.Join(elements, "/"), nil
	}
}
