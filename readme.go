// SPDX-License-Identifier: Unlicense
package main

import (
	"fmt"
	"net/url"
	"os"
)

func createReadme() error {
	badgeURL := "https://img.shields.io/"

	readme := "# " + projectName

	if readmeDescription != "" {
		readme += "\n\n" + readmeDescription
	}

	if readmeAddBadges {
		readme += "\n\n"

		if licenseType != "" {
			readme += fmt.Sprintf(
				"[![License: %s](%sbadge/license-%s-blue.svg)](./LICENSE)",
				licenseType,
				badgeURL,
				licenseType,
			)
		}

		if readmeDocsURL != "" {
			readme += fmt.Sprintf(
				"\n[![Read the docs](%s)](%s)",
				badgeURL+"badge/read_the-docs-darkgreen.svg",
				readmeDocsURL,
			)
		}

		if readmeWebsiteURL != "" {
			readme += fmt.Sprintf(
				"\n[![Website](%s)](%s)",
				badgeURL+"website/?url="+url.QueryEscape(readmeWebsiteURL),
				readmeWebsiteURL,
			)
		}
	}

	if readmeInstallCommand != "" {
		readme += fmt.Sprintf(
			"\n\n## Installation\n\n```sh\n%s\n```",
			readmeInstallCommand,
		)
	}

	if readmeUsageCommand != "" {
		readme += fmt.Sprintf(
			"\n\n## Usage\n\n```sh\n%s\n```",
			readmeUsageCommand,
		)
	}

	if licenseType != "" {
		readme += fmt.Sprintf(
			"\n\n## License\n\nThis project is licensed under the %s license.",
			licenseType,
		)
	}

	f, err := os.Create("README.md")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(readme)
	return err
}
