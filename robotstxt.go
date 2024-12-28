package main

import (
	"io"
	"fmt"
	"os"
	"net/http"
)

func createRobotsTxt() error {
	res, err := http.Get("https://raw.githubusercontent.com/ai-robots-txt/ai.robots.txt/refs/heads/main/robots.txt")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 || res.StatusCode < 200 {
		return fmt.Errorf("failed with status %d", res.StatusCode)
	}

	f,err := os.Create("robots.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	return err
}
