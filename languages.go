package main

import (
	"os/exec"

	"github.com/Iilun/survey/v2"
)

var languageList = []string{
	"bun",
	"deno",
	"go",
	"zig",
}

func createLanguageProject() error {
	var cmd *exec.Cmd

	// NOTE: if you contribute to this, the case statements must be in
	// alphabetical order. If they are out of order, I will request a
	// change.

	switch languageProject {
	case "c":
		// generate makefile
		// write hello world to src/main.c
	case "bun":
		// --yes accepts all defaults, which is needed to avoid
		// interrupting the program.
		cmd = exec.Command("bun", "init", "--yes")
	case "deno":
		cmd = exec.Command("deno", "init")
	case "go":
		var pkgName string

		err := survey.AskOne(&survey.Input{Message: "Package name?"}, &pkgName)
		if err != nil {
			return err
		}

		cmd = exec.Command("go", "mod", "init", pkgName)
	case "zig":
		cmd = exec.Command("zig", "init")
	}

	return cmd.Run()
}
