package goenum

import (
	"fmt"
	"io"

	"github.com/ShawnROGrady/goenum/finder"
	"github.com/ShawnROGrady/goenum/generator"
	"golang.org/x/tools/go/packages"
)

type Generator struct {
	typeName             string
	defaultAsEmptyString bool
}

type GeneratorOpt interface {
	apply(*Generator)
}

type generatorOptFn func(*Generator)

func (fn generatorOptFn) apply(g *Generator) {
	fn(g)
}

func WithDefaultAsEmptyString() GeneratorOpt {
	return generatorOptFn(func(g *Generator) {
		g.defaultAsEmptyString = true
	})
}

func NewGenerator(typeName string, opts ...GeneratorOpt) *Generator {
	g := &Generator{typeName: typeName}
	for _, opt := range opts {
		opt.apply(g)
	}

	return g
}

func (g *Generator) Run(out io.Writer, pkg *packages.Package) error {
	enumFinder := finder.New(g.typeName)

	spec, err := enumFinder.FindFromFiles(pkg.Syntax)
	if err != nil {
		return fmt.Errorf("find enum: %w", err)
	}

	if g.defaultAsEmptyString {
		empty := ""
		for i := range spec.Variants {
			if spec.Variants[i].Data == 0 {
				spec.Variants[i].EnumName = &empty
			}
		}
	}

	generator := generator.New(spec)
	if err := generator.Generate(out); err != nil {
		return err
	}

	return nil
}
