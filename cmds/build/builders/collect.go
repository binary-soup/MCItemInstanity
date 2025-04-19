package builders

import (
	"fmt"
	"item_insanity/cmds/build/data"
	"path/filepath"
)

type CollectBuilder struct{}

func (b CollectBuilder) BuildCollect(dir string, data *data.Collect) Advancement {
	criteria, requirements := b.buildCriteria(data.Items)

	return Advancement{
		Parent:       b.buildParent(dir, data.Parent),
		Display:      b.buildDisplay(data.Display),
		Criteria:     criteria,
		Requirements: requirements,
	}
}

func (b CollectBuilder) buildParent(dir, parent string) string {
	return fmt.Sprintf("%s:%s", PACK_NAMESPACE, filepath.Join(dir, parent))
}

func (b CollectBuilder) buildDisplay(data data.CollectDisplay) AdvancementDisplay {
	builder := DisplayBuilder{}

	return AdvancementDisplay{
		Display: Display{
			Icon:  builder.BuildIcon(data.Item),
			Title: data.Title,
			//TODO: add description builder using name and items
			Description: builder.BuildText("TODO", b.chooseDescriptionColor(data.Frame)),
		},
		Frame:          data.Frame,
		ShowToast:      true,
		AnnounceToChat: true,
	}
}

func (CollectBuilder) chooseDescriptionColor(frame string) string {
	switch frame {
	case FRAME_GOAL:
		return COLOR_YELLOW
	case FRAME_CHALLENGE:
		return COLOR_LIGHT_PURPLE
	}

	return COLOR_YELLOW
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
