package item

import (
	"fmt"
	"testing"

	"github.com/xackery/req-web/testdata"
)

func TestItemGet(t *testing.T) {
	_, c, w, err := testdata.InitEcho()
	if err != nil {
		t.Fatalf("testdata.InitEcho: %s", err)
	}
	c.SetParamNames("id")
	c.SetParamValues("1001")
	err = ItemGet(c)
	if err != nil {
		t.Fatalf("ItemGet: %s", err)
	}

	fmt.Printf("%s", w.Body.Bytes())
	//t.Fatalf("temporary error")
}
