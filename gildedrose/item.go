package gildedrose

import "strings"

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

func (i *Item) IsConjuredItem() bool {
	return strings.HasPrefix(i.Name, "Conjured")
}
