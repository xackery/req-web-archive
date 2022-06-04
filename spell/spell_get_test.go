package spell

import (
	"fmt"
	"testing"

	"github.com/xackery/req-web/testdata"
)

func TestSpellGet(t *testing.T) {
	_, c, w, err := testdata.InitEcho()
	if err != nil {
		t.Fatalf("testdata.InitEcho: %s", err)
	}
	c.SetParamNames("id")
	c.SetParamValues("52")
	err = SpellGet(c)
	if err != nil {
		t.Fatalf("SpellGet: %s", err)
	}

	fmt.Printf("%s", w.Body.Bytes())
	t.Fatalf("temporary error")
}
