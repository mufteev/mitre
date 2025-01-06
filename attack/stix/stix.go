package stix

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"
)

type Bundle struct {
	Objects []map[string]interface{} `json:"objects"`
}

func LoadFromByte(ctx context.Context, data []byte) ([]*Tactic, []*Technique, []*Relationship, error) {
	var bundle Bundle

	if err := json.UnmarshalContext(ctx, data, &bundle); err != nil {
		return nil, nil, nil, fmt.Errorf("decode: %w", err)
	}

	tactics, techniques, relationships, err := parseFlatObjects(bundle.Objects)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("parse flat objects: %w", err)
	}

	return tactics, techniques, relationships, nil
}

func parseFlatObjects(objs []map[string]interface{}) ([]*Tactic, []*Technique, []*Relationship, error) {
	var (
		tactics       = make([]*Tactic, 0)
		techniques    = make([]*Technique, 0)
		relationships = make([]*Relationship, 0)
	)

	for i := 0; i < len(objs); i++ {
		if isRevokedOrDeprecated(objs[i]) {
			continue
		}

		switch objs[i]["type"] {
		case "x-mitre-tactic":
			t, err := parseTactic(objs[i])
			if err != nil {
				return nil, nil, nil, fmt.Errorf("parse tactic: %w", err)
			}

			tactics = append(tactics, t)
		case "attack-pattern":
			t, err := parseTechnique(objs[i])
			if err != nil {
				return nil, nil, nil, fmt.Errorf("parse technique: %w", err)
			}

			techniques = append(techniques, t)
		case "relationship":
			r := parseRelationship(objs[i])
			relationships = append(relationships, r)
		}
	}

	var (
		cpTactics       = make([]*Tactic, len(tactics))
		cpTechniques    = make([]*Technique, len(techniques))
		cpRelationships = make([]*Relationship, len(relationships))
	)

	copy(cpTactics, tactics)
	copy(cpTechniques, techniques)
	copy(cpRelationships, relationships)

	return cpTactics, cpTechniques, cpRelationships, nil
}
