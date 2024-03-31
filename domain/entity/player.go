package entity

import (
	"errors"
	"fmt"

	valueobject "github.com/KoheiMatsuno99/poker/domain/valueobject"
)

type Player struct {
	name  string
	money int
	chips int // 掛け金
	cards []*valueobject.Card
	isActive bool
}

func NewPlayer(name string, money int) *Player {
	return &Player{
		name:  name,
		money: money,
	}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Money() int {
	return p.money
}

func (p *Player) Chips() int {
	return p.chips
}

func (p *Player) Cards() []*valueobject.Card {
	return p.cards
}

func (p *Player) IsActive() bool {
	return p.isActive
}

func (p *Player) DrawCard(card *valueobject.Card) {
	p.cards = append(p.cards, card)
}

func (p *Player) Bet(chips int) error {
	if p.money < chips {
		return errors.New("not enough money")
	}
	p.chips = chips
	p.money -= chips
	return nil
}

func (p *Player) Win(chips int) {
	p.money += chips
}

const numberOfCards = 5

func (p *Player) SortCards() error {
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
	"ハイカード":          0,
	"ワンペア":           1,
	"ツーペア":           2,
	"スリーカード":         3,
	"ストレート":          4,
	"フラッシュ":          5,
	"フルハウス":          6,
	"フォーカード":         7,
	"ストレートフラッシュ":     8,
	"ロイヤルストレートフラッシュ": 9,
}

func HandRankMap() map[string]int {
	return handRankMap
}

func (p *Player) JudgeHands() (string, error) {
	err := p.SortCards()
	if err != nil {
		return "", err
	}
	if p.isRoyalStraightFlush() {
		return "ロイヤルストレートフラッシュ", nil
	}
	if p.isStraightFlush() {
		return "ストレートフラッシュ", nil
	}
	if p.isFourCard() {
		return "フォーカード", nil
	}
	if p.isFullHouse() {
		return "フルハウス", nil
	}
	if p.isFlush() {
		return "フラッシュ", nil
	}
	if p.isStraight() {
		return "ストレート", nil
	}
	if p.isThreeCard() {
		return "スリーカード", nil
	}
	if p.isTwoPair() {
		return "ツーペア", nil
	}
	if p.isOnePair() {
		return "ワンペア", nil
	}
	return "ハイカード", nil
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

// ワンペアの配列、それ以外の配列の順番で返す
func (p *Player) SeparateOnePairAndOtherCards() ([][]*valueobject.Card, error) {
	hands, err := p.JudgeHands()
	if err != nil {
		return nil, err
	}
	if hands != "ワンペア" {
		return nil, fmt.Errorf("not one pair")
	}
	if p.cards[0].Value() == p.cards[1].Value() {
		return [][]*valueobject.Card{{p.cards[0], p.cards[1]}, {p.cards[2], p.cards[3], p.cards[4]}}, nil
	}
	if p.cards[1].Value() == p.cards[2].Value() {
		return [][]*valueobject.Card{{p.cards[1], p.cards[2]}, {p.cards[0], p.cards[3], p.cards[4]}}, nil
	}
	if p.cards[2].Value() == p.cards[3].Value() {
		return [][]*valueobject.Card{{p.cards[2], p.cards[3]}, {p.cards[0], p.cards[1], p.cards[4]}}, nil
	}
	if p.cards[3].Value() == p.cards[4].Value() {
		return [][]*valueobject.Card{{p.cards[3], p.cards[4]}, {p.cards[0], p.cards[1], p.cards[2]}}, nil
	}
	return nil, fmt.Errorf("not one pair")
}

// ツーペアの強い方の配列、ツーペアの弱い方の配列、それ以外の配列の順番で返す
func (p *Player) SeparateTwoPairAndOtherCards() ([][]*valueobject.Card, error) {
	hands, err := p.JudgeHands()
	if err != nil {
		return nil, err
	}
	if hands != "ツーペア" {
		return nil, fmt.Errorf("not two pair")
	}
	if p.cards[0].Value() == p.cards[1].Value() && p.cards[2].Value() == p.cards[3].Value() {
		return [][]*valueobject.Card{{p.cards[2], p.cards[3]}, {p.cards[0], p.cards[1]}, {p.cards[4]}}, nil
	}
	if p.cards[0].Value() == p.cards[1].Value() && p.cards[3].Value() == p.cards[4].Value() {
		return [][]*valueobject.Card{{p.cards[3], p.cards[4]}, {p.cards[0], p.cards[1]}, {p.cards[2]}}, nil
	}
	if p.cards[1].Value() == p.cards[2].Value() && p.cards[3].Value() == p.cards[4].Value() {
		return [][]*valueobject.Card{{p.cards[3], p.cards[4]}, {p.cards[1], p.cards[2]}, {p.cards[0]}}, nil
	}
	return nil, fmt.Errorf("not two pair")
}

// スリーカードの配列、それ以外の配列の順番で返す
func (p *Player) SeparateThreeOfAKindsAndOtherCards() ([][]*valueobject.Card, error) {
	hands, err := p.JudgeHands()
	if err != nil {
		return nil, err
	}
	if hands != "スリーカード" {
		return nil, fmt.Errorf("not three of a kind")
	}
	if p.cards[0].Value() == p.cards[1].Value() && p.cards[1].Value() == p.cards[2].Value() {
		return [][]*valueobject.Card{{p.cards[0], p.cards[1], p.cards[2]}, {p.cards[3], p.cards[4]}}, nil
	}
	if p.cards[1].Value() == p.cards[2].Value() && p.cards[2].Value() == p.cards[3].Value() {
		return [][]*valueobject.Card{{p.cards[1], p.cards[2], p.cards[3]}, {p.cards[0], p.cards[4]}}, nil
	}
	if p.cards[2].Value() == p.cards[3].Value() && p.cards[3].Value() == p.cards[4].Value() {
		return [][]*valueobject.Card{{p.cards[2], p.cards[3], p.cards[4]}, {p.cards[0], p.cards[1]}}, nil
	}
	return nil, fmt.Errorf("not three of a kind")
}

// 3枚組の配列、2枚組の配列の順番で返す
func (p *Player) SeparateThreeOfAKindAndOnePair() ([][]*valueobject.Card, error) {
	hands, err := p.JudgeHands()
	if err != nil {
		return nil, err
	}
	if hands != "フルハウス" {
		return nil, fmt.Errorf("not full house")
	}
	if p.cards[0].Value() == p.cards[1].Value() && p.cards[1].Value() == p.cards[2].Value() && p.cards[3].Value() == p.cards[4].Value() {
		return [][]*valueobject.Card{{p.cards[0], p.cards[1], p.cards[2]}, {p.cards[3], p.cards[4]}}, nil
	}
	if p.cards[0].Value() == p.cards[1].Value() && p.cards[2].Value() == p.cards[3].Value() && p.cards[3].Value() == p.cards[4].Value() {
		return [][]*valueobject.Card{{p.cards[2], p.cards[3], p.cards[4]}, {p.cards[0], p.cards[1]}}, nil
	}
	return nil, fmt.Errorf("not full house")
}

func (p *Player) SeparateFourOfAKindAndOtherCard() ([][]*valueobject.Card, error) {
	hands, err := p.JudgeHands()
	if err != nil {
		return nil, err
	}
	if hands != "フォーカード" {
		return nil, fmt.Errorf("not four of a kind")
	}
	if p.cards[0].Value() == p.cards[1].Value() &&
		p.cards[1].Value() == p.cards[2].Value() &&
		p.cards[2].Value() == p.cards[3].Value() {
		return [][]*valueobject.Card{{p.cards[0], p.cards[1], p.cards[2], p.cards[3]}, {p.cards[4]}}, nil
	}
	if p.cards[1].Value() == p.cards[2].Value() &&
		p.cards[2].Value() == p.cards[3].Value() &&
		p.cards[3].Value() == p.cards[4].Value() {
		return [][]*valueobject.Card{{p.cards[1], p.cards[2], p.cards[3], p.cards[4]}, {p.cards[0]}}, nil
	}
	return nil, fmt.Errorf("not four of a kind")
}
