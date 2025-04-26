package builders

import (
	"fmt"
	"item_insanity/cmds/build/data"
	"item_insanity/common"
)

const (
	COLLECT_ANY = "any"
)

type CollectBuilder struct{}

func (b CollectBuilder) BuildCollect(dir string, data *data.Collect) Advancement {
	criteria, requirements := b.buildCriteria(data.Type, data.Items)

	return Advancement{
		Parent:       AdvancementBuilder{}.BuildParent(dir, data.Parent),
		Display:      b.buildDisplay(data),
		Criteria:     criteria,
		Requirements: requirements,
	}
}

func (b CollectBuilder) buildDisplay(data *data.Collect) AdvancementDisplay {
	builder := DisplayBuilder{}

	return AdvancementDisplay{
		Display: Display{
			Icon:        builder.BuildIcon(data.Display.Item),
			Title:       data.Display.Title,
			Description: b.buildDescription(data),
		},
		Frame:          data.Display.Frame,
		ShowToast:      true,
		AnnounceToChat: true,
	}
}

func (b CollectBuilder) buildDescription(data *data.Collect) []ColoredText {
	prefix := "All the"
	if data.Type == COLLECT_ANY {
		prefix = "Any"
	}

	return []ColoredText{
		DisplayBuilder{}.BuildText(fmt.Sprintf("%s %s", prefix, common.ToUpperSpaced(data.Name)), b.frameColor(data.Type, data.Display.Frame)),
	}
}

func (CollectBuilder) frameColor(collectType, frame string) string {
	if collectType == COLLECT_ANY {
		return COLOR_GREEN
	}

	switch frame {
	case FRAME_CHALLENGE:
		return COLOR_LIGHT_PURPLE
	default:
		return COLOR_YELLOW
	}
}

func (b CollectBuilder) buildCriteria(collectType string, items []string) (map[string]any, [][]string) {
	criteria := map[string]any{}

	builder := AdvancementBuilder{}
	for _, item := range items {
		criteria[item] = builder.BuildCriteria(item, COLLECT_ITEM_TRIGGER)
	}

	if collectType == COLLECT_ANY {
		return criteria, b.buildAnyRequirements(criteria)
	} else {
		return criteria, b.buildAllRequirements(criteria)
	}
}

func (CollectBuilder) buildAllRequirements(criteria map[string]any) [][]string {
	requirements := make([][]string, len(criteria))
	index := 0

	for key := range criteria {
		requirements[index] = []string{key}
		index++
	}

	return requirements
}

func (CollectBuilder) buildAnyRequirements(criteria map[string]any) [][]string {
	requirements := make([]string, len(criteria))
	index := 0

	for key := range criteria {
		requirements[index] = key
		index++
	}

	return [][]string{requirements}

}
