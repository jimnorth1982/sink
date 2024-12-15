package types

type ItemType struct {
	Id   int8   `json:"id"`
	Name string `json:"name"`
}

type Rarity struct {
	Id   int8   `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	Id            int32    `json:"id"`
	Name          string   `json:"name"`
	ItemType      ItemType `json:"item_type"`
	Rarity        Rarity   `json:"rarity"`
	RequiredLevel int      `json:"required_level"`
}