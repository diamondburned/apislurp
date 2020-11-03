package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/diamondburned/apislurp"
	"github.com/k0kubun/pp"
)

var docPaths = map[string][]string{
	"resources": {
		"Audit_Log",
		"Channel",
		"Emoji",
		"Guild",
		"Invite",
		"Template",
		"User",
		"Voice",
		"Webhook",
	},
	"topics": {
		"Gateway",
	},
}

func main() {
	pp.ColoringEnabled = false

	for dir, files := range docPaths {
		log.Println("Wiping directory", dir+"/")

		if err := os.RemoveAll(dir); err != nil {
			log.Fatalln("failed to rmdir:", err)
		}
		if err := os.Mkdir(dir, os.ModeDir|os.ModePerm); err != nil {
			log.Fatalln("failed to mkdir:", err)
		}

		for _, file := range files {
			generateTest(dir, file)
		}
	}
}

var quals = map[string]string{
	"apislurp": "github.com/diamondburned/apislurp",
	"testutil": "github.com/diamondburned/apislurp/internal/tests/testutil",
}

func generateTest(markdownDir, markdownFile string) {
	funcName := strings.ReplaceAll(markdownFile, "_", "")
	dataName := strings.ToLower(funcName) + "Data"
	fileName := filepath.Join(markdownDir, strings.ToLower(funcName)+"_test.go")
	log.Println("Writing to", fileName)

	j := jen.NewFile("apislurp_test")
	j.ImportName("path/filepath", "filepath")

	for name, path := range quals {
		j.ImportName(path, name)
	}

	j.Func().Id("Test"+funcName).Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		jen.Id("path").Op(":=").Qual("path/filepath", "Join").Call(
			jen.Lit(markdownDir), jen.Lit(markdownFile+".md"),
		),
		jen.Id("f").Op(":=").Qual(quals["testutil"], "ReadFile").Call(jen.Id("t"), jen.Id("path")),
		jen.Id("d").Op(":=").Qual(quals["apislurp"], "Parse").Call(jen.Id("f")),
		jen.Qual(quals["testutil"], "Equal").Call(jen.Id("t"), jen.Id("d"), jen.Id(dataName), jen.Lit(markdownFile)),
	)

	d := apislurp.Parse(readFile(markdownDir, markdownFile+".md"))
	v := pp.Sprint(d)

	j.Var().Id(dataName).Op("=").Id(v)

	code := j.GoString()
	writeFile(fileName, code)
}

func writeFile(fileName, code string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("failed to create Go file:", err)
	}
	defer f.Close()

	if _, err := f.WriteString(code); err != nil {
		log.Fatalln("failed to write Go code:", err)
	}
}

func readFile(dir, file string) []byte {
	path := filepath.Join("discord-api-docs", "docs", dir, file)

	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("failed to find file:", err)
	}

	return f
}
