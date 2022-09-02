package goenum

import (
	"fmt"
	"io"

	"github.com/ShawnROGrady/goenum/finder"
	"github.com/ShawnROGrady/goenum/generator"
	"golang.org/x/tools/go/packages"
)

type Generator struct {
	typeName string
}

func NewGenerator(typeName string) *Generator {
	return &Generator{typeName: typeName}
}

func (g *Generator) Run(out io.Writer, pkg *packages.Package) error {
	enumFinder := finder.New(g.typeName)

	spec, err := enumFinder.FindFromFiles(pkg.Syntax)
	if err != nil {
		return fmt.Errorf("find enum: %w", err)
	}

	generator := generator.New(spec)
	if err := generator.Generate(out); err != nil {
		return err
	}

	return nil
}
