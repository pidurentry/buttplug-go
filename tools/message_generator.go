package main

import (
	"errors"
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type msg string
type msgs []msg

type vars struct {
	Package string
	Msgs    msgs
}

func main() {
	flag.Parse()
	if err := generate(flag.Arg(0)); err != nil {
		panic(err)
	}
}

func generate(filepath string) error {
	file, err := parser.ParseFile(token.NewFileSet(), filepath, nil, 0)
	if err != nil {
		return err
	}

	writer, err := writer(filepath)
	if err != nil {
		return err
	}
	defer writer.Close()

	templates, err := templates()
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	vars := vars{
		Package: file.Name.Name,
		Msgs:    parse(file),
	}

	return tmpl.Execute(writer, vars)
}

func writer(file string) (*os.File, error) {
	ext := filepath.Ext(file)
	return os.Create(strings.TrimSuffix(file, ext) + "_msg" + ext)
}

func parse(file *ast.File) msgs {
	msgs := make(msgs, 0)
	for name, obj := range file.Scope.Objects {
		decl, ok := obj.Decl.(*ast.TypeSpec)
		if !ok {
			continue
		}
		switch decl.Type.(type) {
		case *ast.StructType:
			msgs = append(msgs, msg(name))
		}
	}
	return msgs
}

func templates() ([]string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, errors.New("unable to load templates")
	}
	return filepath.Glob(filepath.Join(filepath.Dir(filepath.Dir(filename)), "message", "*.tmpl"))
}
