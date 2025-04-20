package builders

import "item_insanity/cmds/build/data"

type Info struct {
	Display Display
}

type InfoBuilder struct{}

func (InfoBuilder) Build(info data.InfoDisplay, parent string, desc []ColoredText) Advancement {
	builder := DisplayBuilder{}

	return Advancement{
		Parent: parent,
		Display: AdvancementDisplay{
			Display: Display{
				Icon:        builder.BuildIcon(info.Item),
				Title:       info.Title,
				Description: desc,
				Background:  builder.BuildBackground(info.Background),
			},
			Frame:          FRAME_TASK,
			ShowToast:      false,
			AnnounceToChat: false,
		},
		Criteria: map[string]any{
			"tick": map[string]string{
				"trigger": AUTO_TRIGGER,
			},
		},
	}
}

func (b InfoBuilder) BuildRoot(info data.InfoDisplay) Advancement {
	return b.Build(info, "", []ColoredText{
		DisplayBuilder{}.BuildText(info.Description, COLOR_GOLD),
	})
}
