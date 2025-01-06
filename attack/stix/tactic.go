package stix

import (
	"errors"
	"strings"
)

var errNotFoundID = errors.New("tactic ID not found")

// Tactic - описание тактики
type Tactic struct {
	ID          string
	UID         string
	Name        string
	PhaseName   string
	Description string
}

func parseTactic(obj map[string]interface{}) (*Tactic, error) {
	name := obj["name"].(string)
	phaseName := strings.ReplaceAll(strings.ToLower(name), " ", "-")

	t := &Tactic{
		Name:        name,
		PhaseName:   phaseName,
		UID:         obj["id"].(string),
		Description: obj["description"].(string),
	}

	rawExternalReferences, ok := obj["external_references"].([]interface{})
	if !ok {
		return nil, errNotFoundID
	}

	for _, ref := range rawExternalReferences {
		if sourceName := ref.(map[string]interface{})["source_name"].(string); sourceName == "mitre-attack" {
			externalID := ref.(map[string]interface{})["external_id"].(string)
			t.ID = externalID

			break
		}
	}

	if t.ID == "" {
		return nil, errNotFoundID
	}

	return t, nil
}
