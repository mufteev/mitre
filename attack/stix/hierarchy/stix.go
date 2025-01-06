package hierarchy

import (
	"context"
	"fmt"
	"mitre/attack/stix"
)

func LoadFromByteAssociate(ctx context.Context, data []byte) ([]*Tactic, error) {
	tactics, techniques, relationships, err := stix.LoadFromByte(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}

	techniquesHierarchy := associateTechnique(techniques, relationships)
	tacticsHierarchy := associateTactic(tactics, techniquesHierarchy)

	return tacticsHierarchy, nil
}
