package item

// SlotShortNames returns a spaced string with what slots a bit fits on
func SlotShortNames(slotBit int) string {
	mu.RLock()
	defer mu.RUnlock()
	value := ""
	for _, slot := range itemSlots {
		//fmt.Printf("%d vs %d\n", 1<<slot.id&slotBit, 1<<slot.id)
		if 1<<slot.id&slotBit != 1<<slot.id {
			continue
		}
		value += " " + slot.short
	}
	if value == "" {
		value = "NONE"
	} else {
		value = value[1:]
	}
	return value
}
