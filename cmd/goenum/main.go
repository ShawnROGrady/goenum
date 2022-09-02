package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/ShawnROGrady/goenum"
	"golang.org/x/tools/go/packages"
)

func main() {
	var (
		typeName = flag.String("type", "", "name of the type to generate enum boilerplate for")
	)

	flag.Parse()
	args := flag.Args()

	loadPkgConf := &packages.Config{
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			var mode parser.Mode
			mode |= parser.ParseComments
			return parser.ParseFile(fset, filename, src, mode)
		},
		Mode: packages.NeedSyntax | packages.NeedFiles,
	}

	var pkgPath = "."
	if len(args) != 0 {
		pkgPath = args[0]
	}

	if *typeName == "" {
		fmt.Fprintln(os.Stderr, "type is required")
		os.Exit(1)
	}

	pkgs, err := packages.Load(loadPkgConf, pkgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load package %q: %v\n", pkgPath, err)
		os.Exit(1)
	}

	if err := goenum.NewGenerator(*typeName).Run(os.Stdout, pkgs[0]); err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate: %v\n", err)
		os.Exit(1)
	}
}
