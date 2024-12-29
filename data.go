package main

import (
	"encoding/json"

	_ "embed"
)

//go:embed data.json
var rawData []byte

var (
	gitignoreList []string
	gitattributesList []string
	licenseList []string
)

func initData() error {
	var data map[string][]string

	err := json.Unmarshal(rawData, &data)

	gitignoreList = data["gitignore"]
	gitattributesList = data["gitattributes"]
	licenseList = data["licenses"]

	return err
}
