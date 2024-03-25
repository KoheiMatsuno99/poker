package entity

import (
	"errors"
	"fmt"

	valueobject "github.com/KoheiMatsuno99/poker/domain/valueobject"
)

type Player struct {
	name       string
	money      int
	ante       int //　ゲームの参加料
	chips      int // 掛け金
	cards      []*valueobject.Card
}

func NewPlayer(name string, money int) *Player {
	return &Player{
		name:       name,
		money:      money,
	}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Money() int {
	return p.money
}

func (p *Player) Ante() int {
	return p.ante
}

func (p *Player) Chips() int {
	return p.chips
}

func (p *Player) Cards() []*valueobject.Card {
	return p.cards
}

func (p *Player) Bet(chips int) error {
	if p.money < chips {
		return errors.New("not enough money")
	}
	p.chips = chips
	p.money -= chips
	return nil
}

const numberOfCards = 5

func (p *Player) sortCards() error {
	if len(p.cards) != numberOfCards {
		return fmt.Errorf("number of cards is not %d", numberOfCards)
	}
	for i := 0; i < len(p.cards); i++ {
		for j := i + 1; j < len(p.cards); j++ {
			if valueobject.ValueRankMap()[p.cards[i].Value()] > valueobject.ValueRankMap()[p.cards[j].Value()] {
				p.cards[i], p.cards[j] = p.cards[j], p.cards[i]
			} else if valueobject.ValueRankMap()[p.cards[i].Value()] == valueobject.ValueRankMap()[p.cards[j].Value()] {
				if valueobject.SuitRankMap()[p.cards[i].Suit()] > valueobject.SuitRankMap()[p.cards[j].Suit()] {
					p.cards[i], p.cards[j] = p.cards[j], p.cards[i]
				}
			}
		}
	}
	// A, 2, 3, 4, 5のストレートの場合、Aを1として扱う
	if p.cards[0].Value() == "2" &&
		p.cards[1].Value() == "3" &&
		p.cards[2].Value() == "4" &&
		p.cards[3].Value() == "5" &&
		p.cards[4].Value() == "A" {
		p.cards = append(p.cards[len(p.cards)-1:], p.cards[:len(p.cards)-1]...)
	}
	return nil
}

var handRankMap = map[string]int{
	"ハイカード": 0,
	"ワンペア":   1,
	"ツーペア":   2,
	"スリーカード": 3,
	"ストレート":  4,
	"フラッシュ":   5,
	"フルハウス":   6,
	"フォーカード":  7,
	"ストレートフラッシュ": 8,
	"ロイヤルストレートフラッシュ": 9,
}

func HandRankMap() map[string]int {
	return handRankMap
}

func (p *Player) JudgeHandsScore() (int, error) {
	err := p.sortCards()
	if err != nil {
		return -1, err
	}
	if p.isRoyalStraightFlush() {
		return handRankMap["ロイヤルストレートフラッシュ"], nil
	}
	if p.isStraightFlush() {
		return handRankMap["ストレートフラッシュ"], nil
	}
	if p.isFourCard() {
		return handRankMap["フォーカード"], nil
	}
	if p.isFullHouse() {
		return handRankMap["フルハウス"], nil
	}
	if p.isFlush() {
		return handRankMap["フラッシュ"], nil
	}
	if p.isStraight() {
		return handRankMap["ストレート"], nil
	}
	if p.isThreeCard() {
		return handRankMap["スリーカード"], nil
	}
	if p.isTwoPair() {
		return handRankMap["ツーペア"], nil
	}
	if p.isOnePair() {
		return handRankMap["ワンペア"], nil
	}
	return handRankMap["ハイカード"], nil
}

func (p *Player) isRoyalStraightFlush() bool {
	return p.isStraightFlush() &&
		p.cards[0].Value() == "10" &&
		p.cards[1].Value() == "J" &&
		p.cards[2].Value() == "Q" &&
		p.cards[3].Value() == "K" &&
		p.cards[4].Value() == "A"
}

func (p *Player) isStraightFlush() bool {
	return p.isStraight() && p.isFlush()
}

func (p *Player) isFourCard() bool {
	if p.cards[0].Value() == p.cards[1].Value() &&
		p.cards[1].Value() == p.cards[2].Value() &&
		p.cards[2].Value() == p.cards[3].Value() {
		return true
	}
	if p.cards[1].Value() == p.cards[2].Value() &&
		p.cards[2].Value() == p.cards[3].Value() &&
		p.cards[3].Value() == p.cards[4].Value() {
		return true
	}
	return false
}

func (p *Player) isFullHouse() bool {
	if p.cards[0].Value() == p.cards[1].Value() &&
		p.cards[1].Value() == p.cards[2].Value() &&
		p.cards[3].Value() == p.cards[4].Value() {
		return true
	}
	if p.cards[0].Value() == p.cards[1].Value() &&
		p.cards[2].Value() == p.cards[3].Value() &&
		p.cards[3].Value() == p.cards[4].Value() {
		return true
	}
	return false
}

func (p *Player) isFlush() bool {
	return p.cards[0].Suit() == p.cards[1].Suit() &&
		p.cards[1].Suit() == p.cards[2].Suit() &&
		p.cards[2].Suit() == p.cards[3].Suit() &&
		p.cards[3].Suit() == p.cards[4].Suit()
}

func (p *Player) isStraight() bool {
	// valueobject.ValueRankMap()["A"] = 14なので、A, 2, 3, 4, 5のストレートの場合、Aを1として扱う
	if p.cards[0].Value() == "A" && p.cards[1].Value() == "2" && p.cards[2].Value() == "3" && p.cards[3].Value() == "4" && p.cards[4].Value() == "5" {
		return true
	}
	if valueobject.ValueRankMap()[p.cards[0].Value()] == valueobject.ValueRankMap()[p.cards[1].Value()]-1 &&
		valueobject.ValueRankMap()[p.cards[1].Value()] == valueobject.ValueRankMap()[p.cards[2].Value()]-1 &&
		valueobject.ValueRankMap()[p.cards[2].Value()] == valueobject.ValueRankMap()[p.cards[3].Value()]-1 &&
		valueobject.ValueRankMap()[p.cards[3].Value()] == valueobject.ValueRankMap()[p.cards[4].Value()]-1 {
		return true
	}
	return false
}

func (p *Player) isThreeCard() bool {
	if p.cards[0].Value() == p.cards[1].Value() &&
		p.cards[1].Value() == p.cards[2].Value() {
		return true
	}
	if p.cards[1].Value() == p.cards[2].Value() &&
		p.cards[2].Value() == p.cards[3].Value() {
		return true
	}
	if p.cards[2].Value() == p.cards[3].Value() &&
		p.cards[3].Value() == p.cards[4].Value() {
		return true
	}
	return false
}

func (p *Player) isTwoPair() bool {
	if p.cards[0].Value() == p.cards[1].Value() &&
		p.cards[2].Value() == p.cards[3].Value() {
		return true
	}
	if p.cards[0].Value() == p.cards[1].Value() &&
		p.cards[3].Value() == p.cards[4].Value() {
		return true
	}
	if p.cards[1].Value() == p.cards[2].Value() &&
		p.cards[3].Value() == p.cards[4].Value() {
		return true
	}
	return false
}

func (p *Player) isOnePair() bool {
	if p.cards[0].Value() == p.cards[1].Value() {
		return true
	}
	if p.cards[1].Value() == p.cards[2].Value() {
		return true
	}
	if p.cards[2].Value() == p.cards[3].Value() {
		return true
	}
	if p.cards[3].Value() == p.cards[4].Value() {
		return true
	}
	return false
}
