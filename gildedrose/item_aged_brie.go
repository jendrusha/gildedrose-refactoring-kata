package gildedrose

type AgedBrieDecorator struct {
	UpdatableItem
}

func NewAgedBrieDecorator(item UpdatableItem) *AgedBrieDecorator {
	return &AgedBrieDecorator{item}
}

func (d *AgedBrieDecorator) Update() {
	d.UpdatableItem.Update()

	item := d.GetItem()

	increaseQuality := func() {
		if item.Quality < 50 {
			item.Quality++
		}
	}

	increaseQuality()

	if item.SellIn < 0 {
		increaseQuality()
	}
}

func (d *AgedBrieDecorator) GetItem() *Item {
	return d.UpdatableItem.GetItem()
}
