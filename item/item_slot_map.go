package item

type itemSlotDefinition struct {
	id    int
	name  string
	short string
	icon  string
}

var (
	itemSlots = []itemSlotDefinition{
		{
			id:    1, //bit 1
			name:  "Charm",
			short: "CHARM",
			icon:  "xa-shield",
		},
		{
			id:    2, //bit 4
			name:  "Head",
			short: "HEAD",
			icon:  "xa-shield",
		},
		{
			id:    3, //bit 8
			name:  "Face",
			short: "FACE",
			icon:  "xa-shield",
		},
		{
			id:    4, //bit 18
			name:  "Ears",
			short: "EARS",
			icon:  "xa-shield",
		},
		{
			id:    5, //bit 32
			name:  "Neck",
			short: "NECK",
			icon:  "xa-shield",
		},
		{
			id:    6, //bit 64
			name:  "Shoulder",
			short: "SHOULDER",
			icon:  "xa-shield",
		},
		{
			id:    7, //bit 128
			name:  "Arms",
			short: "ARMS",
			icon:  "xa-shield",
		},
		{
			id:    8, //bit 256
			name:  "Back",
			short: "BACK",
			icon:  "xa-shield",
		},
		{
			id:    9, //bit 1536
			name:  "Bracers",
			short: "BRACERS",
			icon:  "xa-shield",
		},
		{
			id:    10, //bit 2048
			name:  "Range",
			short: "RANGE",
			icon:  "xa-shield",
		},
		{
			id:    11, //bit 4096
			name:  "Hands",
			short: "HANDS",
			icon:  "xa-shield",
		},
		{
			id:    12, //bit 8192
			name:  "Primary",
			short: "PRIMARY",
			icon:  "xa-shield",
		},
		{
			id:    13, //bit 16384
			name:  "Second",
			short: "SECONDARY",
			icon:  "xa-shield",
		},
		{
			id:    14, //bit 98304
			name:  "Rings",
			short: "RINGS",
			icon:  "xa-shield",
		},
		{
			id:    15, //bit 131072
			name:  "Chest",
			short: "CHEST",
			icon:  "xa-shield",
		},
		{
			id:    16, //bit 262144
			name:  "Legs",
			short: "LEGS",
			icon:  "xa-shield",
		},
		{
			id:    17, //bit 524288
			name:  "Feet",
			short: "FEET",
			icon:  "xa-shield",
		},
		{
			id:    18, //bit 1048576
			name:  "Waist",
			short: "WAIST",
			icon:  "xa-shield",
		},
		{
			id:    19, //bit 2097152
			name:  "Ammo",
			short: "AMMO",
			icon:  "xa-shield",
		},
		{
			id:    20, //bit 4194304
			name:  "PowerSource",
			short: "POWERSOURCE",
			icon:  "xa-shield",
		},
	}
)
