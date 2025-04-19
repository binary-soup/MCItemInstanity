package builders

import "fmt"

const (
	MINECRAFT_NAMESPACE = "minecraft"
	PACK_NAMESPACE      = "item_insanity"
)

func joinNamespace(namespace, id string) string {
	return fmt.Sprintf("%s:%s", namespace, id)
}
