package stix

// Relationship - связь между объектами
type Relationship struct {
	SourceRef string
	TargetRef string
	Type      string
}

func parseRelationship(obj map[string]interface{}) *Relationship {
	return &Relationship{
		SourceRef: obj["source_ref"].(string),
		TargetRef: obj["target_ref"].(string),
		Type:      obj["relationship_type"].(string),
	}
}
