package extract_cmd

import (
	"fmt"
	"os"
)

type ItemTableWriter struct {
	File *os.File
}

func (w ItemTableWriter) WriteHeader(header string) {
	fmt.Fprintf(w.File, "## %s\n", header)
}

func (w ItemTableWriter) WriteTable(items ItemMap, ids []string) {
	w.WriteTableHeader()

	for _, id := range ids {
		w.WriteItem(items[id])
	}

	fmt.Fprintln(w.File)
}

func (w ItemTableWriter) WriteTableHeader() {
	fmt.Fprint(w.File, "| ID | Type | Name |\n|---|:---:|---|\n")
}

func (w ItemTableWriter) WriteItem(item Item) {
	fmt.Fprintf(w.File, "| %s | [%s] | \"%s\" |\n", item.ID, item.Type, item.Name)
}
