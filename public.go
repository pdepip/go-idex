/*
   public.go
       Handles all public api routes
*/
package idex

import (
    "net/http"
)

func (i *Idex) ReturnCurrencies() (currencies map[string]Currency, err error) {

	currencies = make(map[string]Currency)

    _, err = i.client.do(http.MethodPost, "returnCurrencies", "", true, &currencies)
	return
}
