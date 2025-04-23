package common

import (
	"os"
	"path/filepath"

	"github.com/binary-soup/go-command/util"
)

const (
	ROOT_FILE = "root.json"
	ALL_FILE  = "all.json"
)

type TreeVisitor interface {
	ParseDirectory(path, dir string) error
	ParseRoot(path, dir, file string) error
	ParseCollect(path, dir, file string) error
	ParseAll(path, dir, file string) error
}

type TreeParser struct {
	Visitor TreeVisitor
}

func (p TreeParser) Parse(path, dir string) error {
	entires, err := os.ReadDir(filepath.Join(path, dir))
	if err != nil {
		return util.ChainError(err, "error reading directory")
	}

	dirs := []string{}
	files := []string{}

	for _, entry := range entires {
		if entry.IsDir() {
			dirs = append(dirs, filepath.Join(dir, entry.Name()))
		} else {
			files = append(files, entry.Name())
		}
	}

	// breath-first traversal

	err = p.parseFiles(path, dir, files)
	if err != nil {
		return err
	}

	err = p.parseSubDirectories(path, dirs)
	if err != nil {
		return err
	}

	return nil
}

func (p TreeParser) parseFiles(path, dir string, files []string) error {
	for _, file := range files {
		var err error

		switch file {
		case ROOT_FILE:
			err = p.Visitor.ParseRoot(path, dir, file)
		case ALL_FILE:
			err = p.Visitor.ParseAll(path, dir, file)
		default: //collect
			err = p.Visitor.ParseCollect(path, dir, file)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (p TreeParser) parseSubDirectories(path string, dirs []string) error {
	for _, dir := range dirs {
		err := p.Visitor.ParseDirectory(path, dir)
		if err != nil {
			return err
		}

		err = p.Parse(path, dir)
		if err != nil {
			return err
		}
	}

	return nil
}
