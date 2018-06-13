package public_test

import (
	"github.com/pdepip/go-idex"
	"testing"
)

func TestReturnCurrencies(t *testing.T) {
	client := idex.New("", "")

	res, err := client.ReturnCurrencies()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", res)
}
