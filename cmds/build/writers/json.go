package writers

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

func writeJSON[T any](name, path string, obj T) error {
	os.MkdirAll(filepath.Dir(path), 0700)

	file, err := os.Create(path)
	if err != nil {
		return util.ChainErrorF(err, "error creating %s file", name)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(obj); err != nil {
		return util.ChainErrorF(err, "error encoding %s JSON", name)
	}

	style.Create.PrintF("  + %s\n", path)
	return nil
}
