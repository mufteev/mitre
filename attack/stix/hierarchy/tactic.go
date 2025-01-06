package hierarchy

import "github.com/mufteev/mitre/attack/stix"

// Tactic - описание тактики
type Tactic struct {
	stix.Tactic
	Techniques []*Technique
}
