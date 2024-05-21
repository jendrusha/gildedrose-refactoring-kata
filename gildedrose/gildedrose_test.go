package gildedrose_test

import (
	"testing"

	"github.com/jendrusha/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality(t *testing.T) {
	tests := []struct {
		name          string
		items         []*gildedrose.Item
		expectedItems []*gildedrose.Item
	}{
		{
			name: "Normal Item",
			items: []*gildedrose.Item{
				{"Normal Item", 10, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Normal Item", 9, 19},
			},
		},
		{
			name: "Normal Item on sell date",
			items: []*gildedrose.Item{
				{"Normal Item", 0, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Normal Item", -1, 8},
			},
		},
		{
			name: "Normal Item after sell date",
			items: []*gildedrose.Item{
				{"Normal Item", -10, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Normal Item", -11, 8},
			},
		},
		{
			name: "Normal Item of zero quality",
			items: []*gildedrose.Item{
				{"Normal Item", 5, 0},
			},
			expectedItems: []*gildedrose.Item{
				{"Normal Item", 4, 0},
			},
		},
		{
			name: "Aged Brie",
			items: []*gildedrose.Item{
				{"Aged Brie", 10, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Aged Brie", 9, 21},
			},
		},
		{
			name: "Aged Brie on sell date",
			items: []*gildedrose.Item{
				{"Aged Brie", 0, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Aged Brie", -1, 12},
			},
		},
		{
			name: "Aged Brie on sell date near max quality",
			items: []*gildedrose.Item{
				{"Aged Brie", 0, 49},
			},
			expectedItems: []*gildedrose.Item{
				{"Aged Brie", -1, 50},
			},
		},
		{
			name: "Aged Brie on sell date with max quality",
			items: []*gildedrose.Item{
				{"Aged Brie", 0, 50},
			},
			expectedItems: []*gildedrose.Item{
				{"Aged Brie", -1, 50},
			},
		},
		{
			name: "Aged Brie after sell date",
			items: []*gildedrose.Item{
				{"Aged Brie", -10, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Aged Brie", -11, 12},
			},
		},
		{
			name: "Aged Brie after sell date with max quality",
			items: []*gildedrose.Item{
				{"Aged Brie", -10, 50},
			},
			expectedItems: []*gildedrose.Item{
				{"Aged Brie", -11, 50},
			},
		},
		{
			name: "Sulfuras",
			items: []*gildedrose.Item{
				{"Sulfuras, Hand of Ragnaros", 10, 80},
			},
			expectedItems: []*gildedrose.Item{
				{"Sulfuras, Hand of Ragnaros", 10, 80},
			},
		},
		{
			name: "Backstage passes",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
			},
		},
		{
			name: "Backstage passes close to sell date",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 10, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 9, 22},
			},
		},
		{
			name: "Backstage passes very close to sell date",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 5, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 4, 23},
			},
		},
		{
			name: "Backstage passes after sell date",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 0, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
			},
		},
		{
			name: "Conjured item",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", 10, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", 9, 18},
			},
		},
		{
			name: "Conjured item after sell date",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", 0, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", -1, 16},
			},
		},
		{
			name: "Backstage passes long before sell date",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 11, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 10, 11},
			},
		},
		{
			name: "Backstage passes long before sell date at max quality",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 11, 50},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 10, 50},
			},
		},
		{
			name: "Backstage passes medium close to sell date (upper bound)",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 10, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 9, 12},
			},
		},
		{
			name: "Backstage passes medium close to sell date (upper bound) at max quality",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 10, 50},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
			},
		},
		{
			name: "Backstage passes medium close to sell date (lower bound)",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 6, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 5, 12},
			},
		},
		{
			name: "Backstage passes medium close to sell date (lower bound) at max quality",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 6, 50},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
			},
		},
		{
			name: "Backstage passes very close to sell date (upper bound)",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 5, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 4, 13},
			},
		},
		{
			name: "Backstage passes very close to sell date (upper bound) at max quality",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
			},
		},
		{
			name: "Backstage passes very close to sell date (lower bound)",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 1, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 0, 13},
			},
		},
		{
			name: "Backstage passes very close to sell date (lower bound) at max quality",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 1, 50},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
			},
		},
		{
			name: "Backstage passes after sell date",
			items: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", 0, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
			},
		},
		{
			name: "Conjured item",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", 10, 20},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", 9, 18},
			},
		},
		{
			name: "Conjured item before sell date",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", 5, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", 4, 8},
			},
		},
		{
			name: "Conjured item before sell date at zero quality",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", 5, 0},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", 4, 0},
			},
		},
		{
			name: "Conjured item on sell date",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", 0, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", -1, 6},
			},
		},
		{
			name: "Conjured item on sell date at zero quality",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", 0, 0},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", -1, 0},
			},
		},
		{
			name: "Conjured item after sell date",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", -10, 10},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", -11, 6},
			},
		},
		{
			name: "Conjured item after sell date at zero quality",
			items: []*gildedrose.Item{
				{"Conjured Mana Cake", -10, 0},
			},
			expectedItems: []*gildedrose.Item{
				{"Conjured Mana Cake", -11, 0},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gildedrose.UpdateQuality(test.items)

			for i, item := range test.items {
				expectedItem := test.expectedItems[i]

				if item.Name != expectedItem.Name {
					t.Errorf("Expected Name to be %s, but got %s", expectedItem.Name, item.Name)
				}
				if item.SellIn != expectedItem.SellIn {
					t.Errorf("Expected SellIn to be %d, but got %d", expectedItem.SellIn, item.SellIn)
				}
				if item.Quality != expectedItem.Quality {
					t.Errorf("Expected Quality to be %d, but got %d", expectedItem.Quality, item.Quality)
				}
			}
		})
	}
}
