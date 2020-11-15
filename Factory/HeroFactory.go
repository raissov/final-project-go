package Factory

import "fmt"

func getHero(heroAttribute string) (iHero, error) {
	if heroAttribute == "agility"{
		return newAgilityHero(), nil
	}
	if heroAttribute == "strength"{
		return newStrengthHero(), nil
	}
	if heroAttribute == "intelligence"{
		return newIntelligenceHero(), nil
	}
	return nil, fmt.Errorf("Wrong hero type")
}