package main

import (
	"encoding/json"

	_ "embed"
)

//go:embed data.json
var dataJSON []byte

var (
	gitignoreList []string
	gitattributesList []string
	licenseList []string
)

func initData() error {
	var d map[string][]string

	err := json.Unmarshal(dataJSON, &d)

	gitignoreList = d["gitignore"]
	gitattributesList = d["gitattributes"]
	licenseList = d["licenses"]

	return err
}
