package internal

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path"
)

var (
	useFloat    = flag.Bool("use-float", false, "By default, FIX float fields are represented as arbitrary-precision fixed-point decimal numbers.  Set to 'true' to instead generate FIX float fields as float64 values.")
	pkgRoot     = flag.String("pkg-root", "github.com/quickfixgo", "Set a string here to provide a custom import path for generated packages.")
	tabWidth    = 8
	printerMode = printer.UseSpaces | printer.TabIndent
)

//ParseError indicates generated go source is invalid
type ParseError struct {
	path string
	err  error
}

func (e ParseError) Error() string {
	return fmt.Sprintf("Error parsing %v: %v", e.path, e.err)
}

//ErrorHandler is a convenience struct for interpretting generation Errors
type ErrorHandler struct {
	ReturnCode int
}

//Handle interprets the generation error. Proceeds with setting returnCode, or panics depending on error type
func (h *ErrorHandler) Handle(err error) {
	switch err := err.(type) {
	case nil:
	//do nothing
	case ParseError:
		fmt.Println(err)
		h.ReturnCode = 1
	default:
		panic(err)
	}
}

func write(filePath string, fset *token.FileSet, f *ast.File, size int) error {
	if parentdir := path.Dir(filePath); parentdir != "." {
		if err := os.MkdirAll(parentdir, os.ModePerm); err != nil {
			return err
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	ast.SortImports(fset, f)
	addToBar(int64(size))
	err = (&printer.Config{Mode: printerMode, Tabwidth: tabWidth}).Fprint(file, fset, f)
	closeErr := file.Close()
	printBar(int64(size), filePath)

	if err != nil {
		fmt.Println(err)
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	return nil
}

//WriteFile parses the generated code in fileOut and writes the code out to filePath.
//Function performs some import clean up and gofmts the code before writing
//Returns ParseError if the generated source is invalid but is written to filePath
func WriteFile(filePath, fileOut string) error {
	fset := token.NewFileSet()
	f, pErr := parser.ParseFile(fset, "", fileOut, parser.ParseComments)
	if f == nil {
		return pErr
	}

	//write out the file regardless of parseFile errors
	if err := write(filePath, fset, f, len(fileOut)); err != nil {
		return err
	}

	if pErr != nil {
		return ParseError{path: filePath, err: pErr}
	}

	return nil
}
