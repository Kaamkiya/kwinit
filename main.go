package main

import (
	"os"
	"net/http"
	"io"
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
	gitAttributesTemplates string

	// License specific variables.
	addLicense  bool
	licenseType string
)

func main() {
	survey.AskOne(&survey.Input{Message: "What's your project's name?"}, &projectName)

	// The following questions are whther ot not the user is using Git and other Git-specific questions.
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
			survey.AskOne(&survey.Input{
				Message: "Which templates?",
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
				"mpl-2.0",
				"unlicense",
			},
		}, &licenseType)
	}

	// Add robots.txt?
	// yes -> block what? enter empty line to stop.
	// Init language project?
	// yes -> which language? (bun, go, cargo, node, zig, py, c, etc)

	// This section is for writing files.
	if gitAddIgnore {
		if err := createGitIgnore(); err != nil {
			log.Println("Failed to create gitignore: %v\n", err)
		}
	}

	if gitAddAttributes {
		if err := createGitAttributes(); err != nil {
			log.Println("Failed to create gitattributes: %v\n", err)
		}
	}
}

func createGitIgnore() error {
	res, err := http.Get("https://www.toptal.com/developers/gitignore/api/"+gitIgnoreTemplates)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.Create(".gitignore")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	return err
}

func createGitAttributes() error {
	res, err := http.Get("https://gitattributes.com/api/"+gitAttributesTemplates)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	f, err := os.Create(".gitattributes")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	return err
}
