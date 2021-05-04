package generator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v2"

	"github.com/slasyz/openrest/internal/formatters"
	"github.com/slasyz/openrest/internal/generator/templates"
	"github.com/slasyz/openrest/internal/parsers/swagger"
	"github.com/slasyz/openrest/internal/templater"
)

var (
	errNoGoMod = errors.New("go.mod not found")
)

type Generator struct {
	swagger   *swagger.Parser
	templater *templater.Templater
	formatter formatters.Formatter
	logger    *log.Logger
}

func New(
	swagger *swagger.Parser,
	templater *templater.Templater,
	formatter formatters.Formatter,
	logger *log.Logger,
) *Generator {
	return &Generator{
		swagger:   swagger,
		templater: templater,
		formatter: formatter,
		logger:    logger,
	}
}

func (g *Generator) Do(srcFile string, outDir string) error {
	data, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return fmt.Errorf("error reading source file: %w", err)
	}

	var document swagger.Document
	err = yaml.Unmarshal(data, &document)
	if err != nil {
		return fmt.Errorf("error unmarshalling yaml: %w", err)
	}

	outputPackage, err := generateOutputPackageName(outDir)
	if err != nil {
		return fmt.Errorf("cannot generate output package name: %w", err)
	}

	parsed, err := g.swagger.Parse(&document)
	if err != nil {
		return fmt.Errorf("error parsing swagger file: %w", err)
	}

	err = g.cleanGeneratedDir(outDir)
	if err != nil {
		return fmt.Errorf("error cleaning generated dir: %w", err)
	}

	type fileInfo struct {
		tpl      *template.Template
		data     interface{}
		filename string
	}

	files := []fileInfo{
		{
			tpl: templates.TplDTO,
			data: templater.TemplateDTO{
				Structs: parsed.Structs,
			},
			filename: filepath.Join("generated", "dto.go"),
		},
		{
			tpl: templates.TplHandler,
			data: templater.TemplateHandler{
				OutputPackage: outputPackage,
				Methods:       parsed.Methods,
			},
			filename: filepath.Join("generated", "handler.go"),
		},
		{
			tpl:      templates.TplMethods,
			data:     templater.TemplateMethods{},
			filename: filepath.Join("methods", "methods.go"),
		},
		{
			tpl:      templates.TplError,
			data:     templater.TemplateError{},
			filename: filepath.Join("methods", "error.go"),
		},
	}
	for _, method := range parsed.Methods {
		files = append(files, fileInfo{
			tpl: templates.TplStub,
			data: templater.TemplateStub{
				OutputPackage: outputPackage,
				Method:        &method,
			},
			filename: filepath.Join("methods", method.File),
		})
	}

	for _, file := range files {
		filename := filepath.Join(outDir, file.filename)

		created, err := g.templater.Write(filename, file.tpl, file.data)
		if err != nil {
			return fmt.Errorf("cannot create %s: %w", file.filename, err)
		}

		if !created {
			continue
		}

		g.logger.Println(filename)

		err = g.formatter.Format(filename)
		if err != nil {
			return fmt.Errorf("cannot format %s: %w", file.filename, err)
		}
	}

	return nil
}

func (g *Generator) cleanGeneratedDir(outDir string) error {
	content, err := os.ReadDir(filepath.Join(outDir, "generated"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return fmt.Errorf("cannot list files in generated directory: %w", err)
	}

	g.logger.Println("Removing old files")

	for _, file := range content {
		filename := filepath.Join(filepath.Join(outDir, "generated", file.Name()))
		g.logger.Printf(" - %s\n", filename)

		err = os.RemoveAll(filename)
		if err != nil {
			return fmt.Errorf("cannot remove %s: %w", file.Name(), err)
		}
	}

	return nil
}
