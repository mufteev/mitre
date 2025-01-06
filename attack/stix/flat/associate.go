package flat

import "mitre/attack/stix"

// associateTechnique связывает техники с их родительскими техниками.
// Возвращает список техник с указанными родительскими техниками.
func associateTechnique(techniques []*stix.Technique, relationships []*stix.Relationship) []*Technique {
	mapTechnique := make(map[string]*Technique)

	for _, t := range techniques {
		mapTechnique[t.UID] = &Technique{
			Technique: *t,
		}
	}

	for _, rel := range relationships {
		if rel.Type != "subtechnique-of" {
			continue
		}

		child, ok := mapTechnique[rel.SourceRef]
		if !ok {
			continue
		}

		parent, ok := mapTechnique[rel.TargetRef]
		if !ok {
			continue
		}

		child.ParentID = &parent.ID
	}

	roots := make([]*Technique, len(mapTechnique))
	c := 0

	for _, t := range mapTechnique {
		roots[c] = t
		c++
	}

	return roots
}

// associateTactic связывает тактики с техниками.
// Возвращает список связей между техниками и тактиками.
func associateTactic(tactics []*stix.Tactic, techniques []*Technique) []*TechniqueTactic {
	// Карта для быстрого поиска тактики по имени
	tacticMap := make(map[string]*stix.Tactic, len(tactics))
	for _, tactic := range tactics {
		tacticMap[tactic.PhaseName] = tactic
	}

	techToTactic := make([]*TechniqueTactic, 0)

	for _, technique := range techniques {
		for _, phase := range technique.KillChainPahses {
			if _, ok := tacticMap[phase.TacticName]; !ok {
				continue
			}

			techToTactic = append(techToTactic, &TechniqueTactic{
				TechniqueID: technique.ID,
				TacticID:    tacticMap[phase.TacticName].ID,
			})
		}
	}

	cpTechToTactic := make([]*TechniqueTactic, len(techToTactic))
	copy(cpTechToTactic, techToTactic)

	return cpTechToTactic
}
