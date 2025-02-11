// SPDX-License-Identifier: Unlicense
package main

import (
	"os/exec"

	"github.com/Iilun/survey/v2"
)

var languageList = []string{
	"Cancel.",
	"bun",
	"deno",
	"d",
	"go",
	"gradle",
	"node/npm",
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
	case "d":
		cmd = exec.Command("dub", "init")
	case "go":
		var pkgName string

		err := survey.AskOne(&survey.Input{Message: "Package name?"}, &pkgName)
		if err != nil {
			return err
		}

		cmd = exec.Command("go", "mod", "init", pkgName)
	case "gradle":
		questions := []*survey.Question{
			{
				Name: "projectType",
				Prompt: &survey.Select{
					Message: "Library or application?",
					Options: []string{"library", "application"},
				},
			},
			{
				Name: "dsl",
				Prompt: &survey.Select{
					Message: "Which DSL would you like to use?",
					Options: []string{"kotlin", "groovy"},
				},
			},
			{
				Name: "testFramework",
				Prompt: &survey.Select{
					Message: "Which test framework do you want?",
					Options: []string{"junit", "testng", "spock", "junit-jupiter"},
				},
			},
			{
				Name: "javaVersion",
				Prompt: &survey.Select{
					Message: "Which Java version?",
					// TODO: there must be a better way to do this.
					Options: []string{
						"7",
						"8",
						"9",
						"10",
						"11",
						"12",
						"13",
						"14",
						"15",
						"16",
						"17",
						"18",
						"19",
						"20",
						"21",
					},
				},
			},
			{
				Name: "projectName",
				Prompt: &survey.Input{
					Message: "What's your project's name?",
				},
			},
		}

		gradleConfig := struct {
			ProjectType   string
			Dsl           string
			TestFramework string
			JavaVersion   string
			ProjectName   string
		}{}

		err := survey.Ask(questions, &gradleConfig)
		if err != nil {
			return err
		}

		cmd = exec.Command(
			"gradle",
			"init",
			"--type",
			"java"+gradleConfig.ProjectType,
			"--dsl",
			gradleConfig.Dsl,
			"--test-framework",
			gradleConfig.TestFramework,
			"--java-version",
			gradleConfig.JavaVersion,
			"--project-name",
			gradleConfig.ProjectName,
		)
	case "node/npm":
		cmd = exec.Command("npm", "init", "-y")
	case "zig":
		cmd = exec.Command("zig", "init")
	case "Cancel.":
		return nil
	}

	return cmd.Run()
}
