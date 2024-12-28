package main

import (
	"os"
	"os/exec"
	"net/http"
	"io"
)

func gitInit() error {
	cmd := exec.Command("git", "init")
	cmd.Stdout = os.Stdout
	return cmd.Run()
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

func createGitRemote() error {
	cmd := exec.Command("git", "remote", "add", "origin", gitRemoteAddr)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}