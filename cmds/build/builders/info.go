package builders

import "item_insanity/cmds/build/data"

type InfoBuilder struct{}

func (InfoBuilder) Build(info *data.Info) Advancement {
	builder := DisplayBuilder{}

	return Advancement{
		Display: AdvancementDisplay{
			Display: Display{
				Icon:        builder.BuildIcon(info.Display.Item),
				Title:       info.Display.Title,
				Description: builder.BuildText(info.Display.Description, COLOR_GOLD),
				Background:  builder.BuildBackground(info.Display.Background),
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
