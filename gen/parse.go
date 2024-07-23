package main

import (
	_ "embed"
)

type enterpriseAttack struct {
	Objects []attackPattern `json:"objects"`
}

type attackPattern struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	ExternalReferences []struct {
		SourceName  string `json:"source_name"`
		ExternalId  string `json:"external_id,omitempty"`
		URL         string `json:"url"`
		Description string `json:"description,omitempty"`
	} `json:"external_references"`
}
