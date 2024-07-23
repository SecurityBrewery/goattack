package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"slices"
	"strings"
)

//go:embed foot.gotempl
var foot string

//go:embed head.gotempl
var head string

//go:embed enterprise-attack.json
var enterpriseAttackJSON []byte

type Object struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func main() {
	var data enterpriseAttack

	if err := json.Unmarshal(enterpriseAttackJSON, &data); err != nil {
		log.Fatal(err)
	}

	fmt.Print(head)

	var ids []string
	objects := make(map[string]*Object)

	for _, object := range data.Objects {
		if !(object.Type == "attack-pattern" || object.Type == "x-mitre-tactic") {
			continue
		}

		var externalId, url string

		for _, externalReference := range object.ExternalReferences {
			if externalReference.SourceName == "mitre-attack" {
				externalId = externalReference.ExternalId
				url = externalReference.URL

				break
			}
		}

		ids = append(ids, externalId)
		objects[externalId] = &Object{
			ID:          externalId,
			Type:        object.Type,
			Name:        object.Name,
			Description: strings.TrimSpace(object.Description),
			URL:         url,
		}
	}

	slices.Sort(ids)

	for _, id := range ids {
		if object, ok := objects[id]; ok {
			fullName := object.Name
			parent, _, found := strings.Cut(object.ID, ".")
			if found {
				if parent, ok := objects[parent]; ok {
					fullName = parent.Name + ": " + object.Name
				}
			}

			fmt.Printf("\n\t%q: {ID: %q, Type: %q, Name: %q, FullName: %q, Description: %q, URL: %q},", object.ID, object.ID, object.Type, object.Name, fullName, object.Description, object.URL)
		}
	}

	fmt.Print(foot)
}
