/*
   types.go
       Contains all return and request types
*/
package idex

type Currency struct {
	Decimals int64  `json:"decimals"`
	Address  string `json:"address"`
	Name     string `json:"name"`
}
