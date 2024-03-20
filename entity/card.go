package entity

import "errors"

type Card struct {
	uuid  string
	suit  string
	value string
}

func NewCard(suit string, value string) *Card {
	return &Card{
		suit: suit,
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

func (c *Card) ChangeUUID(uuid string) error {
	if c.uuid != "" {
		return errors.New("uuid already set")
	}
	c.uuid = uuid
	return nil
}