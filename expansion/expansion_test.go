package expansion

import "testing"

func TestExpansionNameByID(t *testing.T) {
	name, err := ExpansionNameByID(0)
	if err != nil {
		t.Fatalf("ExpansionNameByID: %s", err)
	}
	if name != "Classic" {
		t.Fatalf("expected Classic, got %s", name)
	}

	name, err = ExpansionNameByID(0)
	if err != nil {
		t.Fatalf("ExpansionNameByID: %s", err)
	}
	if name != "Classic" {
		t.Fatalf("expected Classic, got %s", name)
	}
}
