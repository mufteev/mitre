package flat

import (
	"context"
	"fmt"
	"mitre/attack/stix"
)

func LoadFromByteAssociate(ctx context.Context, data []byte) ([]*stix.Tactic, []*Technique, []*TechniqueTactic, error) {
	tactics, techniques, relationships, err := stix.LoadFromByte(ctx, data)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load: %w", err)
	}

	techniquesFlat := associateTechnique(techniques, relationships)
	techToTactic := associateTactic(tactics, techniquesFlat)

	return tactics, techniquesFlat, techToTactic, nil
}
