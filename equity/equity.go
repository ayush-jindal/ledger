package equity

type Share struct {
	id string
	name string
}

func NewShare(id, name string) Share {
	return Share{id, name}
}

func (s Share) GetPrice(xchng Exchange) float32 {
	return 0
}

func (s Share) GetDetails() string {
	return s.name
}
