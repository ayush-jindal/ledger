package equity

type ISIN int32

type Share struct {
	id ISIN
	name string
}

func NewShare(name string) Share {
	return Share{0, name}
}

func (s Share) GetID() ISIN { return s.id; }

func (s Share) GetPrice(xchng Exchange) float32 {
	return 100
}

func (s Share) GetDetails() string {
	return s.name
}
