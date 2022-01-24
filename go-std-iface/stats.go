package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type iface struct {
	path  string
	line  int
	count int
}

func numMethods(i *ast.InterfaceType) int {

	count := 0
	for _, m := range i.Methods.List {
		if len(m.Names) == 0 {
			break
		}
		if ast.IsExported(m.Names[0].Name) {
			count++
		}
	}
	return count

}

func fileIfaces(path string) ([]iface, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, string(data), 0)
	if err != nil {
		return nil, err
	}

	// START_IFACES OMIT
	var ifaces []iface
	ast.Inspect(f, func(n ast.Node) bool {
		i, ok := n.(*ast.InterfaceType)
		if !ok {
			return true
		}

		if len(i.Methods.List) == 0 { // interface{}
			return true
		}

		loc := fset.Position(n.Pos())
		ifaces = append(ifaces, iface{
			path:  path,
			line:  loc.Line,
			count: numMethods(i),
		})

		return true
	})
	// END_IFACES OMIT

	return ifaces, nil
}

func main() {
	root := fmt.Sprintf("%s/src", runtime.GOROOT())
	total, count, max := 0, 0, 0
	var maxIface iface

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			if strings.Contains(path, "internal") || strings.Contains(path, "test") {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		ifaces, err := fileIfaces(path)
		if err != nil {
			return err
		}
		if len(ifaces) == 0 {
			return nil
		}
		for _, iface := range ifaces {
			count++
			total += iface.count
			if iface.count > max {
				max = iface.count
				maxIface = iface
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	avg := float64(total) / float64(count)
	fmt.Printf("total: %d, max: %d, avg: %.2f\n", count, max, avg)
	// fmt.Printf("%+v\n", maxIface)
}
