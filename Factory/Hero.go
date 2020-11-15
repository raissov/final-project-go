package Factory

type iHero interface {
	setName(name string)
	setAttribute(attribute string)
	getName() string
	getAttribute() string
}
