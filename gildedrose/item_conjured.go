package gildedrose

type ConjuredDecorator struct {
	UpdatableItem
}

func NewConjuredDecorator(item UpdatableItem) *ConjuredDecorator {
	return &ConjuredDecorator{item}
}

func (d *ConjuredDecorator) Update() {
	d.UpdatableItem.Update()

	item := d.GetItem()

	decreaseQuality := func() {
		if item.Quality > 0 {
			item.Quality--
		}
	}

	decreaseQuality()

	if item.SellIn < 0 {
		decreaseQuality()
	}
}

func (d *ConjuredDecorator) GetItem() *Item {
	return d.UpdatableItem.GetItem()
}
