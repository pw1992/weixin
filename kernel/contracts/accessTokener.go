package contracts

type AccessTokener interface {
	GetToken() string
	Refresh() string
}
