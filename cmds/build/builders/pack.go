package builders

import (
	"fmt"
	"item_insanity/cmds/build/data"
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
