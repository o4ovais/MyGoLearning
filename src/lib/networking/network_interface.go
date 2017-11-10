package networking

type Network interface {
	GetReq(url string) Response
}