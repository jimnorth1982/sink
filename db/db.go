package db

import (
	"errors"
	"sink/types"
)

var rarities = map[int8]types.Rarity{
	1: {Id: 1, Name: "Base"},
	2: {Id: 2, Name: "Magic"},
	3: {Id: 3, Name: "Rare"},
	4: {Id: 4, Name: "Unique"},
}

var item_types = map[int8]types.ItemType{
	1: {Id: 1, Name: "Wand"},
}

var items = map[int32]types.Item{
	1: {
		Id:            1,
		Name:          "Desert Wand",
		Rarity:        rarities[1],
		ItemType:      item_types[1],
		RequiredLevel: 10,
	},
	2: {
		Id:            2,
		Name:          "Cold Ice Wand of Frost",
		Rarity:        rarities[2],
		ItemType:      item_types[1],
		RequiredLevel: 25,
	},
	3: {
		Id:            3,
		Name:          "Casters Desert Wand",
		Rarity:        rarities[3],
		ItemType:      item_types[1],
		RequiredLevel: 40,
	},
	4: {
		Id:            4,
		Name:          "Jimmy's Winter Ice Wand",
		Rarity:        rarities[4],
		ItemType:      item_types[1],
		RequiredLevel: 72,
	},
}

func GetItems() map[int32]types.Item {
	return items
}

func GetItemById(id int32) types.Item {
	return items[id]
}

func AddItem(item types.Item) (added_item types.Item, err error) {
	item, ok := items[item.Id]
	if ok {
		return item, errors.New("item already exists in system")
	}

	items[item.Id] = item
	return item, nil
}