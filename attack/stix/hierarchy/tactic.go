package hierarchy

import "mitre/attack/stix"

// Tactic - описание тактики
type Tactic struct {
	stix.Tactic
	Techniques []*Technique
}
