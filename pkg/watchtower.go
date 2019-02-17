package root

type WatchtowerAttributes struct {
	Address string
	Email   string
	Phone   string
}

type WatchtowerService interface {
	CreateWatchtower(a WatchtowerAttributes) error
}
