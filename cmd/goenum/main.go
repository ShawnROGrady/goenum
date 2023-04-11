package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"os"

	"github.com/ShawnROGrady/goenum"
	"golang.org/x/tools/go/packages"
)

func main() {
	var (
		typeName             = flag.String("type", "", "name of the type to generate enum boilerplate for")
		defaultAsEmptyString = flag.Bool("default-as-empty-string", false, "treat the default value as an empty string")
		output               = flag.String("o", "", "file to write output to; defaults to stdout")
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

	var generatorOpts []goenum.GeneratorOpt
	if *defaultAsEmptyString {
		generatorOpts = append(generatorOpts, goenum.WithDefaultAsEmptyString())
	}

	var (
		dst     io.Writer = os.Stdout
		cleanup func() error
	)

	if *output != "" {
		outFile, err := os.OpenFile(*output, os.O_RDWR, 0)
		if err != nil {
			if !errors.Is(err, fs.ErrNotExist) {
				fmt.Fprintf(os.Stderr, "failed to open existing destination file %q: %v\n", *output, err)
				os.Exit(1)
			}

			// file doesn't exist, create.
			outFile, err = os.Create(*output)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to open new destination file %q: %v\n", *output, err)
				os.Exit(1)
			}

			defer outFile.Close()

			cleanup = func() error { return os.Remove(outFile.Name()) }
		} else {
			defer outFile.Close()

			// read original contents, so we can restore following
			// an error.
			origContents, err := io.ReadAll(outFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to read original contents of %q: %v\n", *output, err)
				os.Exit(1)
			}

			if err := outFile.Truncate(0); err != nil {
				fmt.Fprintf(os.Stderr, "failed to truncate destination file %q: %v\n", *output, err)
				os.Exit(1)
			}
			if _, err := outFile.Seek(0, 0); err != nil {
				fmt.Fprintf(os.Stderr, "failed to set destination file offset: %v\n", err)
				os.Exit(1)
			}

			cleanup = func() error {
				if _, err := outFile.Write(origContents); err != nil {
					return err
				}
				return nil
			}
		}

		dst = outFile
	}

	if err := goenum.NewGenerator(*typeName, generatorOpts...).Run(dst, pkgs[0]); err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate: %v\n", err)
		if cleanup != nil {
			if err := cleanup(); err != nil {
				fmt.Fprintf(os.Stderr, "cleanup failed: %v\n", err)
			}
		}
		os.Exit(1)
	}
}
