package zone

import (
	"fmt"
	"testing"

	"github.com/xackery/req-web/testdata"
)

func TestZoneGet(t *testing.T) {
	_, c, w, err := testdata.InitEcho()
	if err != nil {
		t.Fatalf("testdata.InitEcho: %s", err)
	}
	c.SetParamNames("id")
	c.SetParamValues("3")
	err = ZoneGet(c)
	if err != nil {
		t.Fatalf("ZoneGet: %s", err)
	}

	fmt.Printf("%s", w.Body.Bytes())
	t.Fatalf("temporary error")
}
