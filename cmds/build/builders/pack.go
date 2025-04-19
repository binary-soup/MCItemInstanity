package builders

import "item_insanity/cmds/build/data"

const (
	PACK_FORMAT = 48
)

type MCMeta struct {
	Pack map[string]any `json:"pack"`
}

type PackBuilder struct{}

func (PackBuilder) BuildMeta(pack *data.Pack) MCMeta {
	return MCMeta{
		Pack: map[string]any{
			"description": pack.Description,
			"pack_format": PACK_FORMAT,
		},
	}
}
