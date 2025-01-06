package stix

type KillChainPahses struct {
	Name       string `json:"kill_chain_name"`
	TacticName string `json:"phase_name"`
}

// Technique - описание техники и подтехники
type Technique struct {
	UID             string
	ID              string
	Name            string
	Description     string
	KillChainPahses []KillChainPahses
}

func parseTechnique(obj map[string]interface{}) (*Technique, error) {
	t := Technique{
		UID:         obj["id"].(string),
		Name:        obj["name"].(string),
		Description: obj["description"].(string),
	}

	if rawPhases, ok := obj["kill_chain_phases"].([]interface{}); ok {
		phases := make([]KillChainPahses, 0)

		for _, phase := range rawPhases {
			name := phase.(map[string]interface{})["kill_chain_name"].(string)
			tacticName := phase.(map[string]interface{})["phase_name"].(string)

			phases = append(phases, KillChainPahses{
				Name:       name,
				TacticName: tacticName,
			})
		}

		cpPhases := make([]KillChainPahses, len(phases))
		copy(cpPhases, phases)

		t.KillChainPahses = cpPhases
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

	return &t, nil
}
