package flat

import "github.com/mufteev/mitre/attack/stix"

// Technique - описание техники и подтехники
type Technique struct {
	stix.Technique
	ParentID *string
}
