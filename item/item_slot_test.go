package item

import "testing"

func TestSlotShortNames(t *testing.T) {
	slot := 32
	response := SlotShortNames(slot)
	if response != "NECK" {
		t.Fatalf("unexpected response to %d, got %s", slot, response)
	}
}
