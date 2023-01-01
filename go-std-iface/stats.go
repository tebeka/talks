package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type iface struct {
	path  string
	line  int
	count int
	name  string
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

func getLine(path string, line int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for i := 0; i < line; i++ {
		s.Scan()
	}

	if err := s.Err(); err != nil {
		return "", err
	}

	return s.Text(), nil
}

func ifaceName(path string, line int) string {
	code, err := getLine(path, line)
	if err != nil {
		return ""
	}
	fields := strings.Fields(code)
	// type Pusher interface {
	if len(fields) < 3 {
		return ""
	}
	if fields[0] != "type" {
		return ""
	}
	name := fields[1]
	if !ast.IsExported(name) {
		return ""
	}
	return name
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

	var ifaces []iface
	// START_IFACES OMIT
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
			if strings.Contains(path, "internal") || strings.Contains(path, "test") || strings.Contains(path, "vendor") {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
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
			if iface.count == 0 {
				continue
			}
			name := ifaceName(iface.path, iface.line)
			if name == "" {
				continue
			}
			fmt.Printf("%s:%d %s %d\n", iface.path, iface.line, name, iface.count)

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
	fmt.Printf("%+v\n", maxIface)
}
