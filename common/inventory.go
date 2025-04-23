package common

import (
	"item_insanity/cmds/build/data"
	"path/filepath"
)

type ItemSet map[string]struct{}

type InventoryVisitor struct {
	Set ItemSet
	Ids []string
}

func NewInventoryVisitor() InventoryVisitor {
	return InventoryVisitor{
		Set: ItemSet{},
		Ids: []string{},
	}
}

func (v *InventoryVisitor) ParseCollect(path, dir, file string) error {
	collect, err := data.LoadCollectJSON(filepath.Join(path, dir, file))
	if err != nil {
		return err
	}

	for _, item := range collect.Items {
		_, ok := v.Set[item]
		if ok {
			continue
		}

		v.Set[item] = struct{}{}
		v.Ids = append(v.Ids, item)
	}

	return nil
}

func (v InventoryVisitor) ParseDirectory(_, _ string) error {
	return nil
}

func (v InventoryVisitor) ParseRoot(_, _, _ string) error {
	return nil
}

func (v InventoryVisitor) ParseAll(_, _, _ string) error {
	return nil
}
