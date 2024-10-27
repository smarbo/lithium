package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/mgutz/ansi"
	g "github.com/smarbo/lithium/src/generation"
	p "github.com/smarbo/lithium/src/parser"
	t "github.com/smarbo/lithium/src/tokens"
)

func check(err error) {
	if err != nil {
		fmt.Print("lithium compilation error: ")
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(ansi.Color("Incorrect usage. Correct usage is:", "red"))
		fmt.Println(ansi.Color("lithium", "blue") + " " + ansi.Color("<input.li>", "magenta"))
		os.Exit(1)
	}

	input := os.Args[1]
	file, err := os.Open(input)
	check(err)
	defer file.Close()
	buf, err := io.ReadAll(file)
	check(err)
	data := string(buf)

	tokenizer := t.Init(data)
	err, tokens := tokenizer.Tokenize()
	check(err)

	parser := p.Init(tokens)
	err, prog := parser.ParseProg()
	if prog == nil {
		check(errors.New("invalid program"))
	}
	check(err)

	generator := g.Init(*prog)
	asm_data := generator.GenProg()

	file, err = os.Create("out.asm")
	check(err)
	defer file.Close()
	_, err = file.WriteString(asm_data)
	check(err)

	outPath := "out"
	cmd := exec.Command("nasm", "-felf64", "out.asm")
	cmd.Run()
	cmd = exec.Command("ld", "-o", outPath, "out.o")
	cmd.Run()
	fmt.Println(ansi.Color(input+" compiled to '"+outPath+"' successfully!", "green"))

	os.Exit(0)
}
