package gildedrose

type BaseItem struct {
	*Item
}

func NewBaseItem(item *Item) *BaseItem {
	return &BaseItem{item}
}

func (i *BaseItem) Update() {
	if i.Name != legendaryItemName {
		i.SellIn--
	}
	if i.notIn(legendaryItemName, agedBrieItemName, backstagePassItemName) && i.Quality > 0 {
		i.Quality--
		if i.SellIn < 0 {
			i.Quality--
		}
	}
}

func (i *BaseItem) GetItem() *Item {
	return i.Item
}

func (i *BaseItem) notIn(names ...string) bool {
	for _, name := range names {
		if i.Name == name {
			return false
		}
	}
	return true
}
