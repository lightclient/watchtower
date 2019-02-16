package root

type WatchtowerAttributes struct {
	Address string
}

type WatchtowerService interface {
	CreateWatchtower(a WatchtowerAttributes) error
}
