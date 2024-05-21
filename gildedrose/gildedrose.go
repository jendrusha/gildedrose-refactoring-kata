package gildedrose

import (
	"runtime"
	"sync"
)

const (
	legendaryItemName     = "Sulfuras, Hand of Ragnaros"
	agedBrieItemName      = "Aged Brie"
	backstagePassItemName = "Backstage passes to a TAFKAL80ETC concert"
)

type UpdatableItem interface {
	Update()
	GetItem() *Item
}

func UpdateQuality(items []*Item) {
	updatableItems := createUpdatableItems(items)
	tasks := make(chan UpdatableItem, len(updatableItems))

	var wg sync.WaitGroup

	startWorkers(&wg, tasks)
	assignTasks(updatableItems, tasks)

	wg.Wait()

	updateItems(items, updatableItems)
}

func createUpdatableItems(items []*Item) []UpdatableItem {
	updatableItems := make([]UpdatableItem, len(items))
	for i, item := range items {
		updatableItems[i] = newUpdatableItem(item)
	}
	return updatableItems
}

func startWorkers(wg *sync.WaitGroup, tasks chan UpdatableItem) {
	numWorkers := runtime.NumCPU()

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ui := range tasks {
				ui.Update()
			}
		}()
	}
}

func assignTasks(updatableItems []UpdatableItem, tasks chan UpdatableItem) {
	for _, updatableItem := range updatableItems {
		tasks <- updatableItem
	}
	close(tasks)
}

func updateItems(items []*Item, updatableItems []UpdatableItem) {
	for i, updatableItem := range updatableItems {
		items[i].SellIn = updatableItem.GetItem().SellIn
		items[i].Quality = updatableItem.GetItem().Quality
	}
}

func newUpdatableItem(item *Item) UpdatableItem {
	baseItem := NewBaseItem(item)

	switch item.Name {
	case agedBrieItemName:
		return NewAgedBrieDecorator(baseItem)
	case legendaryItemName:
		return NewSulfurasDecorator(baseItem)
	case backstagePassItemName:
		return NewBackstagePassDecorator(baseItem)
	default:
		if item.IsConjuredItem() {
			return NewConjuredDecorator(baseItem)
		}
		return baseItem
	}
}
