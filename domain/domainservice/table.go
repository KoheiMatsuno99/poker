package domainservice

import (
	"github.com/KoheiMatsuno99/poker/domain/entity"
)

type Table struct {
	uuid    string
	players []*entity.Player
}

func NewTable(uuid string, players []*entity.Player) *Table {
	return &Table{
		uuid:    uuid,
		players: players,
	}
}

func (t *Table) Uuid() string {
	return t.uuid
}

func (t *Table) Players() []*entity.Player {
	return t.players
}
