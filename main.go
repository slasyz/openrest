package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/slasyz/openrest/internal/formatters"
	"github.com/slasyz/openrest/internal/generator"
	"github.com/slasyz/openrest/internal/parsers/schema"
	"github.com/slasyz/openrest/internal/parsers/swagger"
	"github.com/slasyz/openrest/internal/templater"
)

func main() {
	srcFile := flag.String("srcFile", "", "")
	outDir := flag.String("outDir", "", "")
	goImportsCmd := flag.String("goImportsCmd", "", "if not specified, go fmt will be used")
	flag.Parse()

	if *srcFile == "" {
		flag.Usage()
		os.Exit(1)
		return
	}

	lggr := log.Default()
	schm := schema.New()
	swgr := swagger.New(schm)
	tmpltr := templater.New()
	frmtr := initFormatter(*goImportsCmd)
	gnrtr := generator.New(swgr, tmpltr, frmtr, lggr)

	err := gnrtr.Do(*srcFile, *outDir)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
		return
	}
}

func initFormatter(goImportsCmd string) formatters.Formatter {
	if goImportsCmd == "" {
		return formatters.NewGoFmt()
	} else {
		return formatters.NewGoImports(goImportsCmd)
	}
}
