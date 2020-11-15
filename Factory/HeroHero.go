package Factory

type hero struct {
	name string
	attribute string
}

func (h *hero) setName(name string)  {
	h.name = name
}

func (h *hero) getName() string {
	return h.name
}

func (h *hero) setAttribute(attribute string)  {
	h.attribute = attribute
}

func (h *hero) getAttribute() string {
	return h.attribute
}