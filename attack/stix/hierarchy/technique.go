package hierarchy

import "github.com/mufteev/mitre/attack/stix"

// Technique - описание техники и подтехники
type Technique struct {
	stix.Technique
	SubTechniques []*Technique
}
