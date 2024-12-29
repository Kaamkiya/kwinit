package main

import (
	"log"
	"time"

	"github.com/Iilun/survey/v2"
	"github.com/briandowns/spinner"
)

var (
	projectName string

	// Git specific variables.
	usingGit               bool
	gitAddRemote           bool
	gitRemoteAddr          string
	gitAddIgnore           bool
	gitIgnoreTemplates     []string
	gitAddAttributes       bool
	gitAttributesTemplates []string

	// License specific variables.
	addLicense    bool
	licenseType   string
	licenseHolder string
	licenserEmail string

	// Readme specific variables.
	addReadme            bool
	readmeDescription    string
	readmeDocsURL        string
	readmeWebsiteURL     string
	readmeInstallCommand string
	readmeUsageCommand   string
	readmeAddBadges      bool

	// Language project specific variables.
	initLanguage    bool
	languageProject string

	// Miscellaneous other variables.
	addRobotsTxt           bool
	addIssueTemplates      bool
	addPullRequestTemplate bool
)

func main() {
	err := initData()
	if err != nil {
		panic(err)
	}

	err = survey.AskOne(&survey.Input{Message: "What's your project's name?"}, &projectName)
	check(err)

	// The following questions are whether ot not the user is using Git and other Git-specific questions.
	err = survey.AskOne(&survey.Confirm{
		Message: "Are you using Git?",
		Default: true,
	}, &usingGit)
	check(err)

	if usingGit {
		err = survey.AskOne(&survey.Confirm{
			Message: "Add a remote?",
			Default: false,
		}, &gitAddRemote)
		check(err)

		if gitAddRemote {
			err = survey.AskOne(&survey.Input{
				Message: "Where are you hosting your project? Enter the full URL:",
			}, &gitRemoteAddr)
			check(err)
		}

		err = survey.AskOne(&survey.Confirm{
			Message: "Add .gitignore?",
			Default: true,
		}, &gitAddIgnore)
		check(err)

		if gitAddIgnore {
			err = survey.AskOne(&survey.MultiSelect{
				Message: "Which templates?",
				Options: gitignoreList,
			}, &gitIgnoreTemplates)
			check(err)
		}

		err = survey.AskOne(&survey.Confirm{
			Message: "Add .gitattributes?",
			Default: true,
		}, &gitAddAttributes)
		check(err)

		if gitAddAttributes {
			err = survey.AskOne(&survey.MultiSelect{
				Message: "Which templates?",
				Options: gitattributesList,
			}, &gitAttributesTemplates)
			check(err)
		}
	}

	// The following questions are about licenses.
	err = survey.AskOne(&survey.Confirm{Message: "Add license?"}, &addLicense)
	check(err)
	if addLicense {
		err = survey.AskOne(&survey.Select{
			Message: "Which license would you like to add?",
			Options: licenseList,
		}, &licenseType)
		check(err)

		err = survey.AskOne(&survey.Input{
			Message: "What should the copyright holder's name be?",
		}, &licenseHolder)
		check(err)

		err = survey.AskOne(&survey.Input{
			Message: "What should the copyright holder's email be?",
		}, &licenserEmail)
		check(err)
	}

	err = survey.AskOne(&survey.Confirm{Message: "Add readme?"}, &addReadme)
	check(err)
	if addReadme {
		check(survey.AskOne(&survey.Input{Message: "Project description"}, &readmeDescription))
		check(survey.AskOne(&survey.Input{Message: "Documentation URL (blank to skip)"}, &readmeDocsURL))
		check(survey.AskOne(&survey.Input{Message: "Website (blank to skip)"}, &readmeWebsiteURL))
		check(survey.AskOne(&survey.Input{Message: "Installation command? (blank to skip?)"}, &readmeInstallCommand))
		check(survey.AskOne(&survey.Input{Message: "Usage command? (blank to skip?)"}, &readmeUsageCommand))
		check(survey.AskOne(&survey.Confirm{Message: "Add badges to readme?"}, &readmeAddBadges))
	}

	check(survey.AskOne(&survey.Confirm{Message: "Add robots.txt?"}, &addRobotsTxt))

	check(survey.AskOne(&survey.Confirm{Message: "Add issue templates?"}, &addIssueTemplates))
	check(survey.AskOne(&survey.Confirm{Message: "Add pull request template?"}, &addPullRequestTemplate))
	// Add CONTRIBUTING.md?
	// Add CODE_OF_CONDUCT.md?

	// Init language project?
	check(survey.AskOne(&survey.Confirm{Message: "Init language project?"}, &initLanguage))
	if initLanguage {
		err = survey.AskOne(
			&survey.Select{
				Message: "Select the language",
				Options: languageList,
			},
			&languageProject,
		)
		check(err)
	}
	// yes -> which language? (bun, go, cargo, node, zig, py, c, etc)

	// This section is for writing files.
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	if usingGit {
		s.Suffix = " Initializing Git..."
		if err := gitInit(); err != nil {
			log.Printf("Failed to initialize git: %v\n", err)
		}

		s.Suffix = " Making .gitignore..."
		if gitAddIgnore {
			if err := createGitIgnore(); err != nil {
				log.Printf("Failed to create gitignore: %v\n", err)
			}
		}

		s.Suffix = " Making .gitattributes..."
		if gitAddAttributes {
			if err := createGitAttributes(); err != nil {
				log.Printf("Failed to create gitattributes: %v\n", err)
			}
		}

		s.Suffix = " Adding git remote..."
		if gitAddRemote {
			if err := createGitRemote(); err != nil {
				log.Printf("Failed to add git remote: %v\n", err)
			}
		}
	}

	if addLicense {
		s.Suffix = " Making LICENSE..."

		if err := createLicense(); err != nil {
			log.Printf("Failed to create LICENSE: %v\n", err)
		}
	}

	if addReadme {
		s.Suffix = " Making README..."

		if err := createReadme(); err != nil {
			log.Printf("Failed to create README: %v\n", err)
		}
	}

	if addRobotsTxt {
		s.Suffix = " Making robots.txt..."
		if err := createRobotsTxt(); err != nil {
			log.Printf("Failed to create robots.txt: %v\n", err)
		}
	}

	if addIssueTemplates {
		s.Suffix = " Adding issue templates..."
		if err := createIssueTemplates(); err != nil {
			log.Printf("Failed to add issue templates: %v\n", err)
		}
	}

	if addPullRequestTemplate {
		s.Suffix = " Adding pull request template..."
		if err := createPullRequestTemplate(); err != nil {
			log.Printf("Failed to add PR template: %v\n", err)
		}
	}

	if initLanguage {
		// This prevents the spinner from interfering with any
		// inputs we may need to take.
		s.Stop()
		if err := createLanguageProject(); err != nil {
			log.Printf("Failed to initialize language project: %v\n", err)
		}
		s.Start()
		s.Suffix = " Initializing language project..."
	}

	s.Stop()
}

func check(err error) {
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}
