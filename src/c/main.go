package main

import (
	"./parser"
	"bytes"
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/bachvtuan/mime2extension"
	flatbuffers "github.com/google/flatbuffers/go"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type TreeShapeListener struct {
	*parser.BaseCListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

var Parser *parser.CParser
var stack [][]flatbuffers.UOffsetT
var children []flatbuffers.UOffsetT
var builder *flatbuffers.Builder
var isOpening bool
var text flatbuffers.UOffsetT
var tail flatbuffers.UOffsetT

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	isOpening = true
	children := make([]flatbuffers.UOffsetT, 0)
	stack = append(stack, children)
	if AppConfig.Service.OutputFormat == "xml" {
		ruleIndex := ctx.GetRuleContext().GetRuleIndex()
		fmt.Printf("<%s>", Parser.GetRuleNames()[ruleIndex])
	}
}

func (this *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	isOpening = false
	n := len(stack)
	children := stack[n-1]
	ElementStartChildVector(builder, len(children))
	for _, c := range children {
		builder.PrependUOffsetT(c)
	}
	inv := builder.EndVector(len(children))
	kind := int32(ctx.GetRuleContext().GetRuleIndex())
	ElementStart(builder)
	ElementAddKind(builder, kind)
	if isOpening {
		ElementAddText(builder, text)
	} else {
		ElementAddTail(builder, tail)
	}
	ElementAddChild(builder, inv)
	element := ElementEnd(builder)
	if n > 0 {
		n := len(stack) - 1
		stack[n-1] = append(stack[n-1], element)
		stack = stack[:n]
	}
	if AppConfig.Service.OutputFormat == "xml" {
		ruleIndex := ctx.GetRuleContext().GetRuleIndex()
		fmt.Printf("</%s>", Parser.GetRuleNames()[ruleIndex])
	}
}

func (this *TreeShapeListener) VisitTerminal(node antlr.TerminalNode) {
	t := node.GetText()
	if t == "<EOF>" {
		t = ""
	}
	str := builder.CreateString(t)
	if isOpening {
		text = str
	} else {
		tail = str
	}
	if AppConfig.Service.OutputFormat == "xml" {
		fmt.Printf("%s", t)
	}
}

func printElement(element *Element) {
	if element == nil {
		return
	}
	k := element.Kind()
	kind := Parser.GetRuleNames()[int(k)]
	fmt.Printf("<%s>%s", kind, string(element.Text()))
	n := element.ChildLength()
	for i := 0; i < n; i++ {
		c := &Element{}
		if element.Child(c, i) {
			printElement(c)
		}
	}
	fmt.Printf("%s</%s>", string(element.Tail()), kind)
}

func printData(data *Data) {
	kind := data.RecordType(nil)
	element := kind.Element(nil)
	printElement(element)
}

func codeToFlatbuffers(filename string) []byte {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)
	Parser = p
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	builder = flatbuffers.NewBuilder(1024)
	stack = make([][]flatbuffers.UOffsetT, 1)
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	element := stack[0][0]
	Anonymous2Start(builder)
	Anonymous2AddElement(builder, element)
	record := Anonymous2End(builder)
	DataStart(builder)
	DataAddRecordType(builder, record)
	offset := DataEnd(builder)
	builder.Finish(offset)
	buf := builder.FinishedBytes()
	return buf
}

func mimeOfFile(filename string) string {
	cmd := exec.Command("file", "-b", "--mime-type", os.Args[1])
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		return "unknown"
	}
	mimetype := strings.TrimSpace(out.String())
	err, filetype := mime2extension.Extension(mimetype)
	if err != nil {
		log.Fatal(err)
		return "unknown"
	}
	return filetype
}

func main() {
	dir := flag.String("dir", ".", "config.yaml dir")
	flag.Parse()
	LoadConfig(*dir)
	filetype := mimeOfFile(os.Args[1])
	switch filetype {
	case "conf":
		buf := codeToFlatbuffers(os.Args[1])
		ioutil.WriteFile(os.Args[2], buf, 0644)
	case "bin":
		buf, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		data := GetRootAsData(buf, 0)
		if AppConfig.Service.OutputFormat == "text" {
			input := antlr.NewInputStream("1 + 2 * 3;")
			lexer := parser.NewCLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			Parser = parser.NewCParser(stream)
			printData(data)
		} else {
			fmt.Printf("%v\n", data)
		}
	}
}
