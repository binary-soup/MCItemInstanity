package builders

const (
	COLOR_WHITE        = "white"
	COLOR_YELLOW       = "yellow"
	COLOR_LIGHT_PURPLE = "light_purple"
	COLOR_GOLD         = "gold"
)

type Display struct {
	Icon        map[string]string `json:"icon"`
	Title       string            `json:"title"`
	Description []ColoredText     `json:"description"`
	Background  string            `json:"background,omitempty"`
}

type ColoredText struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}

type DisplayBuilder struct{}

func (DisplayBuilder) BuildIcon(id string) map[string]string {
	return map[string]string{"id": joinNamespace(MINECRAFT_NAMESPACE, id)}
}

func (DisplayBuilder) BuildBackground(id string) string {
	return joinNamespace(MINECRAFT_NAMESPACE, id)
}

func (DisplayBuilder) BuildText(text, color string) ColoredText {
	return ColoredText{
		Text:  text,
		Color: color,
	}
}
