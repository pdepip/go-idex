/*
   idex.go
       Wrapper for idex Exchange API
*/
package idex

const (
	BaseUrl = "https://api.idex.market"
)

type Idex struct {
	client *Client
}

func New(address, privKey string) *Idex {
	client := NewClient(address, privKey)
	return &Idex{client}
}
