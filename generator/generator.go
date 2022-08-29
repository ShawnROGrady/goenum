package generator

import (
	"fmt"
	"io"

	"github.com/ShawnROGrady/goenum/model"
	"github.com/ShawnROGrady/goenum/templates"
)

type Generator struct {
	spec *model.EnumSpec
}

func New(spec *model.EnumSpec) *Generator {
	return &Generator{spec: spec}
}

func (g *Generator) writeEnumNames(dst io.Writer) error {
	if err := templates.WriteEnumNames(dst, g.spec); err != nil {
		return fmt.Errorf("writeEnumNames: %w", err)
	}

	return nil
}

func (g *Generator) writeEnumValues(dst io.Writer) error {
	if err := templates.WriteEnumValues(dst, g.spec); err != nil {
		return fmt.Errorf("writeEnumValues: %w", err)
	}

	return nil
}

func (g *Generator) writeStringMethod(dst io.Writer) error {
	if err := templates.WriteStringMethod(dst, g.spec); err != nil {
		return fmt.Errorf("writeStringMethod: %w", err)
	}

	return nil
}

func (g *Generator) writeInvalidNameError(dst io.Writer) error {
	if err := templates.WriteInvalidNameError(dst, g.spec); err != nil {
		return fmt.Errorf("writeInvalidNameError: %w", err)
	}

	return nil
}

func (g *Generator) writeFromString(dst io.Writer) error {
	if err := templates.WriteFromString(dst, g.spec); err != nil {
		return fmt.Errorf("writeFromString: %w", err)
	}

	return nil
}

func (g *Generator) writeImplTextMarshaler(dst io.Writer) error {
	if err := templates.WriteImplTextMarshaler(dst, g.spec); err != nil {
		return fmt.Errorf("writeImplTextMarshaler: %w", err)
	}

	return nil
}

func (g *Generator) writeImplTextUnmarshaler(dst io.Writer) error {
	if err := templates.WriteImplTextUnmarshaler(dst, g.spec); err != nil {
		return fmt.Errorf("writeImplTextUnmarshaler: %w", err)
	}

	return nil
}

func (g *Generator) writeHeader(dst io.Writer) error {
	if err := templates.WriteHeader(dst, g.spec); err != nil {
		return fmt.Errorf("writeHeader: %w", err)
	}

	return nil
}

func (g *Generator) Generate(dst io.Writer) error {
	fns := []func(io.Writer) error{
		g.writeHeader,
		g.writeEnumNames,
		g.writeEnumValues,
		g.writeStringMethod,
		g.writeInvalidNameError,
		g.writeFromString,
		g.writeImplTextMarshaler,
		g.writeImplTextUnmarshaler,
	}

	for _, fn := range fns {
		if err := fn(dst); err != nil {
			return err
		}
	}

	return nil
}
