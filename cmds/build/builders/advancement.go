package builders

const (
	AUTO_TRIGGER         = "minecraft:tick"
	COLLECT_ITEM_TRIGGER = "minecraft:inventory_changed"
)

type Advancement struct {
	Parent       string             `json:"parent,omitempty"`
	Display      AdvancementDisplay `json:"display"`
	Criteria     map[string]any     `json:"criteria"`
	Requirements [][]string         `json:"requirements,omitempty"`
}

type AdvancementDisplay struct {
	Display
	Frame          string `json:"frame"`
	ShowToast      bool   `json:"show_toast"`
	AnnounceToChat bool   `json:"announce_to_chat"`
}

type AdvancementCriteria struct {
	Conditions map[string]any `json:"conditions"`
	Trigger    string         `json:"trigger"`
}

type AdvancementBuilder struct{}

func (b AdvancementBuilder) BuildCriteria(item string, trigger string) AdvancementCriteria {
	return AdvancementCriteria{
		Conditions: map[string]any{
			"items": []map[string]string{{"items": joinNamespace(MINECRAFT_NAMESPACE, item)}},
		},
		Trigger: trigger,
	}
}
