package hierarchy

import (
	"context"
	"fmt"
	"io"

	"github.com/mufteev/mitre/attack/stix"
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

func LoadFromReaderAssociate(ctx context.Context, r io.Reader) ([]*Tactic, error) {
	tactics, techniques, relationships, err := stix.LoadFromReader(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}

	techniquesHierarchy := associateTechnique(techniques, relationships)
	tacticsHierarchy := associateTactic(tactics, techniquesHierarchy)

	return tacticsHierarchy, nil
}
