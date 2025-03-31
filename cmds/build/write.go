package build_cmd

import (
	"encoding/json"
	"os"

	"github.com/binary-soup/go-command/util"
)

func writeJSON[T any](name, path string, obj T) error {
	file, err := os.Create(path)
	if err != nil {
		return util.ChainErrorF(err, "error creating %s file", name)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(obj); err != nil {
		return util.ChainErrorF(err, "error encoding %s JSON", name)
	}

	return nil
}
