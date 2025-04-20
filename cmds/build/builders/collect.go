package builders

import (
	"fmt"
	"item_insanity/cmds/build/data"
	"strings"
)

type CollectBuilder struct{}

func (b CollectBuilder) BuildCollect(dir string, data *data.Collect) Advancement {
	criteria, requirements := b.buildCriteria(data.Items)

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
	desc := make([]ColoredText, 1+len(data.Items))
	builder := DisplayBuilder{}

	// header
	desc[0] = builder.BuildText(fmt.Sprintf("All the %s\n\n", idToUpperSpaced(data.Name)), b.frameColor(data.Display.Frame))

	for i, item := range data.Items {
		desc[i+1] = builder.BuildText(fmt.Sprintf("- %s\n", idToLowerSpaced(item)), COLOR_WHITE)
	}

	desc[len(desc)-1].Text = strings.TrimSuffix(desc[len(desc)-1].Text, "\n")
	return desc
}

func (CollectBuilder) frameColor(frame string) string {
	switch frame {
	case FRAME_CHALLENGE:
		return COLOR_LIGHT_PURPLE
	default:
		return COLOR_YELLOW
	}
}

func (CollectBuilder) buildCriteria(items []string) (map[string]any, [][]string) {
	criteria := map[string]any{}

	builder := AdvancementBuilder{}
	for _, item := range items {
		criteria[item] = builder.BuildCriteria(item, COLLECT_ITEM_TRIGGER)
	}

	requirements := make([][]string, len(criteria))
	index := 0

	for key := range criteria {
		requirements[index] = []string{key}
		index++
	}

	return criteria, requirements
}
