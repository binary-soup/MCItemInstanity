package extract_cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/style"
	"github.com/binary-soup/go-command/util"
)

const ITEM_TABLE_FILE = "item_table.md"

type ItemTableWriter struct {
	file *os.File
}

func (w *ItemTableWriter) OpenFile(path string) error {
	var err error
	fullPath := filepath.Join(path, ITEM_TABLE_FILE)

	w.file, err = os.Create(fullPath)
	if err != nil {
		return util.ChainError(err, "error creating item table file")
	}

	style.Create.PrintF("+ %s\n", fullPath)

	// write header
	fmt.Fprint(w.file, "| ID | Type | Name |\n|---|:---:|---|\n")

	return nil
}

func (w ItemTableWriter) CloseFile() {
	w.file.Close()
}

func (w ItemTableWriter) WriteItem(item Item) {
	fmt.Fprintf(w.file, "| %s | [%s] | \"%s\" |\n", item.ID, item.Type, item.Name)
}
