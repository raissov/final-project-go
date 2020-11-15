package Builder

type heroBuilder interface {

}

func getBuilder(builderType string) heroBuilder  {
	if builderType == "agility" {
		return &agilityBuilder{}
	}
	return nil
}
