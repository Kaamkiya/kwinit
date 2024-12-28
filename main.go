package main

import (
	"log"

	"github.com/Iilun/survey/v2"
)

var (
	projectName string

	// Git specific variables.
	usingGit               bool
	gitAddRemote           bool
	gitRemoteAddr          string
	gitAddIgnore           bool
	gitIgnoreTemplates     string
	gitAddAttributes       bool
	gitAttributesTemplates []string

	// License specific variables.
	addLicense    bool
	licenseType   string
	licenseHolder string
	licenserEmail string
)

func main() {
	survey.AskOne(&survey.Input{Message: "What's your project's name?"}, &projectName)

	// The following questions are whether ot not the user is using Git and other Git-specific questions.
	survey.AskOne(&survey.Confirm{
		Message: "Are you using Git?",
		Default: true,
	}, &usingGit)
	if usingGit {
		survey.AskOne(&survey.Confirm{
			Message: "Add a remote?",
			Default: false,
		}, &gitAddRemote)
		if gitAddRemote {
			survey.AskOne(&survey.Input{
				Message: "Where are you hosting your project? Enter the full URL:",
			}, &gitRemoteAddr)
		}

		survey.AskOne(&survey.Confirm{
			Message: "Add .gitignore?",
			Default: true,
		}, &gitAddIgnore)
		if gitAddIgnore {
			survey.AskOne(&survey.Input{
				Message: "Which templates?",
			}, &gitIgnoreTemplates)
		}

		survey.AskOne(&survey.Confirm{
			Message: "Add .gitattributes?",
			Default: true,
		}, &gitAddAttributes)
		if gitAddAttributes {
			survey.AskOne(&survey.MultiSelect{
				Message: "Which templates?",
				Options: []string{
					"actionscript",
					"ada",
					"ballerina",
					"c++",
					"csharp",
					"common",
					"deplhi",
					"drupal",
					"dyalogapl",
					"elixir",
					"flutter",
					"fortran",
					"fountain",
					"fsharp",
					"go",
					"hashicorp",
					"java",
					"lua",
					"markdown",
					"mathematica",
					"matlab",
					"microsoftshell",
					"objectivec",
					"php",
					"pascal",
					"perl",
					"powershell",
					"python",
					"r",
					"rails",
					"ruby",
					"servoy",
					"swift",
					"sql",
					"tinacms",
					"unity",
					"vim",
					"web",
					"devcontainer",
					"visualstudio",
					"visualstudiocode",
				},
			}, &gitAttributesTemplates)
		}
	}

	// The following questions are about licenses.
	survey.AskOne(&survey.Confirm{Message: "Add license?"}, &addLicense)
	if addLicense {
		survey.AskOne(&survey.Select{
			Message: "Which license would you like to add?",
			Options: []string{
				"0bsd",
				"agpl-3.0",
				"apache-2.0",
				"artistic-2.0",
				"bsd-2-clause",
				"bsd-3-clause",
				"cc0-1.0",
				"gpl-2.0",
				"gpl-3.0",
				"lgpl-3.0",
				"isc",
				"mit",
				"mpl-2.0",
				"unlicense",
			},
		}, &licenseType)
		survey.AskOne(&survey.Input{
			Message: "What should the copyright holder's name be?",
		}, &licenseHolder)
		survey.AskOne(&survey.Input{
			Message: "What should the copyright holder's email be?",
		}, &licenserEmail)
	}

	// Add robots.txt?
	// yes -> block what? enter empty line to stop.
	// Init language project?
	// yes -> which language? (bun, go, cargo, node, zig, py, c, etc)

	// This section is for writing files.
	if usingGit {
		if err := gitInit(); err != nil {
			log.Printf("Failed to initialize git: %v\n", err)
		}
	}
	if gitAddIgnore {
		if err := createGitIgnore(); err != nil {
			log.Printf("Failed to create gitignore: %v\n", err)
		}
	}
	if gitAddAttributes {
		if err := createGitAttributes(); err != nil {
			log.Printf("Failed to create gitattributes: %v\n", err)
		}
	}
	if gitAddRemote {
		if err := createGitRemote(); err != nil {
			log.Printf("Failed to add git remote: %v\n", err)
		}
	}

	if addLicense {
		if err := createLicense(); err != nil {
			log.Printf("Failed to create LICENSE: %v\n", err)
		}
	}
}
