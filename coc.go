package main

import (
	"bytes"
	"os"

	_ "embed"
)

//go:embed CODE_OF_CONDUCT.md
var codeOfConductText []byte

func createCodeOfConduct() error {
	codeOfConductText = bytes.ReplaceAll(codeOfConductText, []byte("[contact_method]"), []byte(cocContact))
	f, err := os.Create("CODE_OF_CONDUCT.md")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(codeOfConductText)
	return err
}
