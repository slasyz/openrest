package templater

import "github.com/slasyz/openrest/internal/parsers"

type TemplateDTO struct {
	Structs []parsers.Struct
}

type TemplateHandler struct {
	OutputPackage string
	Methods       []parsers.Method
}

type TemplateMethods struct {
}

type TemplateError struct {
}

type TemplateStub struct {
	OutputPackage string
	Method        *parsers.Method
}
