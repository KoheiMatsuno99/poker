package valueobject

type Card struct {
	uuid  string
	suit  string
	value string
}

func NewCard(suit string, value string) *Card {
	return &Card{
		suit:  suit,
		value: value,
	}
}

func (c *Card) Suit() string {
	return c.suit
}

func (c *Card) Value() string {
	return c.value
}

func (c *Card) UUID() string {
	return c.uuid
}

var valueRankMap = map[string]int{
	"A":  14,
	"K":  13,
	"Q":  12,
	"J":  11,
	"10": 10,
	"9":  9,
	"8":  8,
	"7":  7,
	"6":  6,
	"5":  5,
	"4":  4,
	"3":  3,
	"2":  2,
}

var suitRankMap = map[string]int{
	"spade":   4,
	"heart":   3,
	"diamond": 2,
	"club":    1,
}

func ValueRankMap() map[string]int {
	return valueRankMap
}

func SuitRankMap() map[string]int {
	return suitRankMap
}

func Suits() []string {
	keys := []string{}
	for key := range suitRankMap {
		keys = append(keys, key)
	}
	return keys
}

func Values() []string {
	keys := []string{}
	for key := range valueRankMap {
		keys = append(keys, key)
	}
	return keys
}
