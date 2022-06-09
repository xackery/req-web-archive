package zone

import (
	"context"
	"testing"

	"github.com/xackery/req-web/db"
)

func TestZoneNameByID(t *testing.T) {
	err := db.Init(context.Background())
	if err != nil {
		t.Fatalf("db.Init: %s", err)
	}

	name, err := ZoneShortNameByID(1)
	if err != nil {
		t.Fatalf("ZoneShortNameByID: %s", err)
	}
	if name != "abysmal" {
		t.Fatalf("expected abysmal, got %s", name)
	}

	name, err = ZoneShortNameByID(1)
	if err != nil {
		t.Fatalf("ZoneShortNameByID: %s", err)
	}
	if name != "abysmal" {
		t.Fatalf("expected abysmal, got %s", name)
	}

}
