package main

import (
	"os/exec"

	"github.com/Iilun/survey/v2"
)

var languageList = []string{
	"bun",
	"go",
	"zig",
}

func createLanguageProject() error {
	var cmd *exec.Cmd

	switch languageProject {
	case "c":
		// generate makefile
		// write hello world to src/main.c
	case "bun":
		// --yes accepts all defaults, which is needed to avoid
		// interrupting the program.
		cmd = exec.Command("bun", "init", "--yes")
	case "go":
		var packageName string

		err := survey.AskOne(&survey.Input{Message: "Package name?"}, &packageName)
		if err != nil {
			return err
		}

		cmd = exec.Command("go", "mod", "init", packageName)
	case "zig":
		cmd = exec.Command("zig", "init")
	}

	return cmd.Run()
}
