package main

import (
	"fmt"
	"log"

	"github.com/MakeNowJust/memefish/pkg/parser"
	"github.com/k0kubun/pp"
)

func main() {
	// Create a new Parser instance.
	file := &parser.File{
		Buffer: "SELECT * FROM customers",
	}
	p := &parser.Parser{
		Lexer: &parser.Lexer{File: file},
	}

	// Do parsing!
	stmt, err := p.ParseQuery()
	if err != nil {
		log.Fatal(err)
	}

	// Show AST.
	log.Print("AST")
	pp.Println(stmt)

	// Unparse AST to SQL source string.
	log.Print("Unparse")
	fmt.Println(stmt.SQL())
}
