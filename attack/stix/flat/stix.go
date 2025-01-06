package flat

import (
	"context"
	"fmt"
	"io"

	"github.com/mufteev/mitre/attack/stix"
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

func LoadFromReaderAssociate(ctx context.Context, r io.Reader) ([]*stix.Tactic, []*Technique, []*TechniqueTactic, error) {
	tactics, techniques, relationships, err := stix.LoadFromReader(ctx, r)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load: %w", err)
	}

	techniquesFlat := associateTechnique(techniques, relationships)
	techToTactic := associateTactic(tactics, techniquesFlat)

	return tactics, techniquesFlat, techToTactic, nil
}
