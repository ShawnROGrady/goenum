package goenum_test

import (
	"bytes"
	"embed"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/fs"
	"testing"

	"github.com/ShawnROGrady/goenum"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/packages"
)

var (
	//go:embed testdata/*
	testData embed.FS
)

func mustReadFile(t testing.TB, fileSystem fs.FS, name string) []byte {
	t.Helper()
	file, err := fileSystem.Open(name)
	if err != nil {
		t.Fatalf("open file %q: %v", name, err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("open file %q: %v", name, err)
	}
	return b
}

func readTestData(t testing.TB, name string) []byte {
	t.Helper()
	return mustReadFile(t, testData, name)
}

func mustParseFileContents(t testing.TB, contents []byte) *ast.File {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", contents, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse file: %v", err)
	}
	return f
}

func mustReadAndParseFile(t testing.TB, fileSystem fs.FS, name string) *ast.File {
	t.Helper()
	contents := mustReadFile(t, fileSystem, name)
	f := mustParseFileContents(t, contents)
	return f
}

func readAndParseTestData(t testing.TB, name string) *ast.File {
	t.Helper()
	return mustReadAndParseFile(t, testData, name)
}

func TestGenerateSuccess(t *testing.T) {
	testCases := map[string]struct {
		typeName       string
		inputFile      *ast.File
		expectedOutput []byte
	}{
		"animal": {
			typeName: "Animal",
			inputFile: readAndParseTestData(
				t, "testdata/good/animal.go",
			),
			expectedOutput: readTestData(
				t, "testdata/good/animal_out.go",
			),
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			runGeneratorTest(
				t,
				testCase.typeName,
				testCase.inputFile,
				testCase.expectedOutput,
			)
		})
	}
}

func runGeneratorTest(t testing.TB, typeName string, inputFile *ast.File, expectedOutput []byte) {
	t.Helper()

	var dst bytes.Buffer

	generator := goenum.NewGenerator(typeName)

	if err := generator.Run(&dst, &packages.Package{
		Syntax: []*ast.File{inputFile},
	}); err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	/*
		if dst.String() != string(expectedOutput) {
			t.Errorf("generated:\n%s\n\nwant:\n%s", dst.String(), expectedOutput)
		}
	*/
	assert.Equal(t, string(expectedOutput), dst.String())
}
