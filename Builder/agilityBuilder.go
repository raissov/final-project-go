package Builder

type agilityBuilder struct{
	damage int
	maxMana int
	maxHealth int
	armor int
}

func newAgilityBuilder() *agilityBuilder {
	return &agilityBuilder{}
}

func setDamage(a *agilityBuilder)  {
	a.damage = 10
}
