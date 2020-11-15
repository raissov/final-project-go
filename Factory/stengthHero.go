package Factory

type strengthHero struct {
	hero
}

func newStrengthHero() iHero {
	return &strengthHero{
		hero{
			name:      "name",
			attribute: "strength",
		},
	}
}

