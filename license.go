// SPDX-License-Identifier: Unlicense
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func createLicense() error {
	res, err := http.Get("https://raw.githubusercontent.com/github/choosealicense.com/refs/heads/gh-pages/_licenses/" + licenseType + ".txt")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 || res.StatusCode < 200 {
		return fmt.Errorf("Request failed with status %d", res.StatusCode)
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

	_, err = f.WriteString(text)
	return err
}
