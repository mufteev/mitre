package hierarchy

import "github.com/mufteev/mitre/attack/stix"

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

		if _, ok := mapTechnique[rel.TargetRef]; ok {
			if mapTechnique[rel.TargetRef].SubTechniques == nil {
				mapTechnique[rel.TargetRef].SubTechniques = make([]*Technique, 0)
			}

			child := mapTechnique[rel.SourceRef]

			mapTechnique[rel.TargetRef].SubTechniques = append(mapTechnique[rel.TargetRef].SubTechniques, child)
		}

		delete(mapTechnique, rel.SourceRef)
	}

	roots := make([]*Technique, len(mapTechnique))
	c := 0

	for _, t := range mapTechnique {
		subTechs := make([]*Technique, len(t.SubTechniques))
		copy(subTechs, t.SubTechniques)

		t.SubTechniques = subTechs
		roots[c] = t
		c++
	}

	return roots
}

func associateTactic(tactics []*stix.Tactic, techniques []*Technique) []*Tactic {
	// Карта для быстрого поиска тактики по имени
	tacticMap := make(map[string]*Tactic)
	for _, tactic := range tactics {
		tacticMap[tactic.PhaseName] = &Tactic{
			Tactic: *tactic,
		}
	}

	for _, technique := range techniques {
		for _, phase := range technique.KillChainPahses {
			if _, ok := tacticMap[phase.TacticName]; !ok {
				continue
			}

			if tacticMap[phase.TacticName].Techniques == nil {
				tacticMap[phase.TacticName].Techniques = make([]*Technique, 0)
			}

			tacticMap[phase.TacticName].Techniques = append(tacticMap[phase.TacticName].Techniques, technique)
		}
	}

	roots := make([]*Tactic, len(tacticMap))
	c := 0

	for _, tactic := range tacticMap {
		techs := make([]*Technique, len(tactic.Techniques))
		copy(techs, tactic.Techniques)
		tactic.Techniques = techs

		roots[c] = tactic
		c++
	}

	return roots
}
