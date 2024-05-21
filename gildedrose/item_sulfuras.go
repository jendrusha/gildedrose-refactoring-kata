package gildedrose

import "log/slog"

type SulfurasDecorator struct {
	UpdatableItem
}

func NewSulfurasDecorator(item UpdatableItem) *SulfurasDecorator {
	return &SulfurasDecorator{item}
}

func (d *SulfurasDecorator) Update() {
	slog.Info("Legendary item, no changes")
}

func (d *SulfurasDecorator) GetItem() *Item {
	return d.UpdatableItem.GetItem()
}
