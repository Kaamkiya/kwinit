package main

import (
	"strconv"
	"time"
	"net/http"
	"fmt"
	"io"
	"os"
	"strings"
)

func createLicense() error {
	res, err := http.Get("https://raw.githubusercontent.com/github/choosealicense.com/refs/heads/gh-pages/_licenses/"+licenseType+".txt")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 || res.StatusCode < 200 {
		return fmt.Errorf("Request failed with status %s", res.StatusCode)
	}

	f, err := os.Create("LICENSE")
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	text := strings.Split(string(data), "---")[2]
	text = strings.ReplaceAll(text, "[year]", strconv.Itoa(time.Now().Year()))
	text = strings.ReplaceAll(text, "[fullname]", licenseHolder)
	text = strings.ReplaceAll(text, "[email]", licenserEmail)

	f.WriteString(text)

	return err
}