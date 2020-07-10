package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/parser"
	"github.com/z7zmey/php-parser/printer"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {
	filename := "examples/class-wp-site.php"
	src, _ := ioutil.ReadFile(filename)
	parser, _ := parser.NewParser(src, "5.4")
	parser.Parse(true)
	for _, err := range parser.GetErrors() {
		fmt.Println(err)
	}
	rootNode := parser.GetRootNode()
	pp.Println(rootNode.GetPosition())
	pp.Println(rootNode.GetFreeFloating())

	out, _ := os.Create("examples/class-wp-site2.txt")
	rootNode.Walk(&visitor.GoDumper{Writer: out})

	dst, _ := os.Create("examples/class-wp-site2.php")
	ppt := printer.NewPrettyPrinter(dst, "\t")
	ppt.Print(rootNode)

	var klass *stmt.Class
	if r, ok := rootNode.(*node.Root); ok && len(r.Stmts) > 0 {
		if klass, ok = r.Stmts[0].(*stmt.Class); !ok {
			return
		}
	}
	pp.Println(klass.Position)
	pp.Println(klass.ClassName)
	pp.Println(klass.PhpDocComment)
}
