package coin

type Coin struct {
	ID     string
	Symbol string
	Name   string
}

func NewCoin(id string) *Coin {
	return &Coin{ID: id}
}
