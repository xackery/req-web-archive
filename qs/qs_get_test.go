package qs

import (
	"fmt"
	"testing"

	"github.com/xackery/req-web/testdata"
)

func TestQsGet(t *testing.T) {
	_, c, w, err := testdata.InitEcho()
	if err != nil {
		t.Fatalf("testdata.InitEcho: %s", err)
	}
	c.SetParamNames("id")
	c.SetParamValues("123")
	err = QsGet(c)
	if err != nil {
		t.Fatalf("QsGet: %s", err)
	}

	c.SetParamValues("abc")
	err = QsGet(c)
	if err == nil {
		t.Fatalf("expected nil, got error %s", err)
	}
	fmt.Printf("%s", w.Body.Bytes())
	t.Fatalf("temporary error")
}
