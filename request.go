package rem

type IRequest interface {
	GetMethod() string
	GetHeaders() map[string]string
	GetCookies() map[string]string
	GetBody() string
}