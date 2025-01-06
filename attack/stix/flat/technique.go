package flat

import "mitre/attack/stix"

// Technique - описание техники и подтехники
type Technique struct {
	stix.Technique
	ParentID *string
}
