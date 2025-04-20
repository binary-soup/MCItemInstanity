package builders

import (
	"fmt"
	"item_insanity/cmds/build/data"
	"strings"
)

const (
	PACK_FORMAT = 71 //1.21.5
)

type MCMeta struct {
	Pack map[string]any `json:"pack"`
}

type PackBuilder struct{}

func (PackBuilder) BuildMeta(pack *data.Pack) MCMeta {
	return MCMeta{
		Pack: map[string]any{
			"description": fmt.Sprintf("[%s] %s", pack.Version, pack.Description),
			"pack_format": PACK_FORMAT,
		},
	}
}

func (b PackBuilder) BuildAuthor(pack data.PackAuthor) Advancement {
	desc := make([]ColoredText, 2+len(pack.Socials))
	builder := DisplayBuilder{}

	desc[0] = builder.BuildText("Created by ", COLOR_WHITE)
	desc[1] = builder.BuildText(fmt.Sprintf("%s\n\n", pack.Creator), COLOR_YELLOW)

	maxLen := 0

	for i, social := range pack.Socials {
		desc[i+2] = builder.BuildText(fmt.Sprintf("%s\n", social.Handle), social.Color)
		maxLen = max(maxLen, len(social.Handle))
	}

	desc[len(desc)-1].Text = strings.TrimSuffix(desc[len(desc)-1].Text, "\n")

	info := data.InfoDisplay{
		Item:  "player_head",
		Title: fmt.Sprintf("%-*s", maxLen, "Author"),
	}
	return InfoBuilder{}.Build(info, AdvancementBuilder{}.BuildParent(".", "root"), desc)
}
