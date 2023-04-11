package finder

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"regexp"

	"github.com/ShawnROGrady/goenum/model"
)

type valueParser struct {
	pkgName     string
	values      map[string][]*ast.ValueSpec
	currentType *ast.Ident
}

func (p *valueParser) currentTypeName() string {
	if p.currentType == nil {
		return ""
	}

	return p.currentType.Name
}

func (p *valueParser) extract(valueSpec *ast.ValueSpec) {
	if typeIdent, ok := valueSpec.Type.(*ast.Ident); ok {
		typeName := typeIdent.Name
		p.values[typeName] = append(p.values[typeName], valueSpec)
		p.currentType = typeIdent
		return
	}

	if len(valueSpec.Values) != 0 {
		p.currentType = nil
		return
	}

	p.values[p.currentTypeName()] = append(p.values[p.currentTypeName()], valueSpec)
}

func (p *valueParser) parseSpecs(specs []ast.Spec) {
	if p.values == nil {
		p.values = map[string][]*ast.ValueSpec{}
	}

	p.currentType = nil

	for _, spec := range specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		p.extract(valueSpec)
	}

	p.currentType = nil
}

func (p *valueParser) parseGenDecl(genDecl *ast.GenDecl) {
	if genDecl.Tok != token.CONST {
		return
	}

	p.parseSpecs(genDecl.Specs)
}

var enumNameRe = regexp.MustCompile(`//goenum:"name=([a-zA-Z0-9]+)"`)

func (p *valueParser) specByName(name string) (*model.EnumSpec, error) {
	vals, ok := p.values[name]
	if !ok {
		return nil, fmt.Errorf("could not find any constant decls for %q", name)
	}

	spec := &model.EnumSpec{Name: name, Package: p.pkgName}
	variants := []model.Variant{}
	for _, val := range vals {
		var enumName string
		if val.Comment != nil {
			matches := enumNameRe.FindAllStringSubmatch(val.Comment.List[0].Text, -1)
			if len(matches) != 0 {
				enumName = matches[0][1]
			}
		}

		for _, ident := range val.Names {
			variants = append(variants, model.Variant{
				GoName:   ident.Name,
				EnumName: enumName,
			})
		}
	}

	spec.Variants = variants
	return spec, nil
}

func (p *valueParser) parseFile(f *ast.File) {
	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			p.parseGenDecl(genDecl)
		}
	}
}

// EnumFinder is responsible for extracting an EnumSpec from source code.
type EnumFinder struct {
	name string
}

// New returns a new EnumFinder which searches for the spec with the provided
// name.
func New(name string) *EnumFinder {
	return &EnumFinder{
		name: name,
	}
}

// FindFromFiles extraces the EnumSpec from the provided files.
func (finder *EnumFinder) FindFromFiles(fs []*ast.File) (*model.EnumSpec, error) {
	if len(fs) == 0 {
		return nil, errors.New("at least 1 file must be provided")
	}

	p := new(valueParser)
	p.pkgName = fs[0].Name.Name

	// TODO: make sure type is numeric.
	for _, f := range fs {
		p.parseFile(f)
	}

	return p.specByName(finder.name)
}
