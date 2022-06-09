package expansion

import (
	"fmt"
	"sync"
)

var (
	mu             sync.RWMutex
	expansionNames = make(map[int]string)
)

// ExpansionNameByID returns an expansion's name when provided an id
func ExpansionNameByID(id int) (string, error) {
	var err error
	mu.Lock()
	defer mu.Unlock()

	name, ok := expansionNames[id]
	if ok {
		return name, nil
	}
	name, err = expansionName(id)
	if err != nil {
		return "", fmt.Errorf("expansionName: %w", err)
	}
	expansionNames[id] = name
	return name, nil
}

func expansionName(id int) (string, error) {
	if id == 0 {
		return "Classic", nil
	}
	if id == 1 {
		return "Ruins of Kunark", nil
	}
	if id == 2 {
		return "Scars of Velious", nil
	}
	if id == 3 {
		return "Shadows of Luclin", nil
	}
	if id == 4 {
		return "Planes of Power", nil
	}
	if id == 5 {
		return "Legacy of Ykesha", nil
	}
	if id == 6 {
		return "Dungeons of Norrath", nil
	}
	if id == 7 {
		return "Gates of Discord", nil
	}
	if id == 8 {
		return "Omens of War", nil
	}
	if id == 9 {
		return "Dragons of Norrath", nil
	}
	if id == 10 {
		return "Depths of Darkhollow", nil
	}
	if id == 11 {
		return "Prophecy of Ro", nil
	}
	if id == 12 {
		return "Serpent's Spine", nil
	}
	if id == 13 {
		return "Buried Sea", nil
	}
	if id == 14 {
		return "Secrets of Faydwer", nil
	}
	if id == 15 {
		return "Seeds of Destruction", nil
	}
	if id == 16 {
		return "Underfoot", nil
	}
	if id == 17 {
		return "House of Thule", nil
	}
	if id == 18 {
		return "Veil of Alaris", nil
	}
	if id == 19 {
		return "Rain of Fear", nil
	}
	if id == 20 {
		return "of the Forsaken", nil
	}
	if id == 21 {
		return "Darkened Sea", nil
	}
	if id == 22 {
		return "Broken Mirror", nil
	}
	if id == 23 {
		return "Empires of Kunark", nil
	}
	if id == 24 {
		return "Ring of Scale", nil
	}
	if id == 25 {
		return "Burning Lands", nil
	}
	if id == 26 {
		return "Torment of Velious", nil
	}
	if id == 27 {
		return "Claws of Veeshan", nil
	}
	if id == 28 {
		return "Terror of Luclin", nil
	}
	if id == 98 {
		return "Not in game/Disabled", nil
	}

	return "", fmt.Errorf("unknown expansion: %d", id)
}
