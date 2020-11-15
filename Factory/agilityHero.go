package Factory

type agilityHero struct {
	hero
}

func newAgilityHero() iHero {
	return &agilityHero{
		hero{
			name:      "name",
			attribute: "agility",
		},
	}
}
