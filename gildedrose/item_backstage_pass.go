package gildedrose

type BackstagePassDecorator struct {
	UpdatableItem
}

func NewBackstagePassDecorator(item UpdatableItem) *BackstagePassDecorator {
	return &BackstagePassDecorator{item}
}

func (d *BackstagePassDecorator) Update() {
	d.UpdatableItem.Update()

	item := d.GetItem()

	if item.SellIn < 0 {
		item.Quality = 0
		return
	}

	increaseQuality := func(increment int) {
		if item.Quality+increment > 50 {
			item.Quality = 50
		} else {
			item.Quality += increment
		}
	}

	switch {
	case item.SellIn < 5:
		increaseQuality(3)
	case item.SellIn < 10:
		increaseQuality(2)
	default:
		increaseQuality(1)
	}
}

func (d *BackstagePassDecorator) GetItem() *Item {
	return d.UpdatableItem.GetItem()
}
