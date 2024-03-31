package domainservice

import (
	"fmt"

	"math/rand"

	"github.com/KoheiMatsuno99/poker/domain/entity"
	"github.com/KoheiMatsuno99/poker/domain/valueobject"
)

type Table struct {
	uuid    string
	deck    []*valueobject.Card
	players []*entity.Player
}

func NewTable(uuid string, players []*entity.Player) *Table {
	return &Table{
		uuid:    uuid,
		deck:    initialDeck,
		players: players,
	}
}

func (t *Table) Uuid() string {
	return t.uuid
}

func (t *Table) Players() []*entity.Player {
	return t.players
}

func createDeck() []*valueobject.Card {
	deck := []*valueobject.Card{}
	for _, suit := range valueobject.Suits() {
		for _, value := range valueobject.Values() {
			deck = append(deck, valueobject.NewCard(suit, value))
		}
	}
	return deck
}

func shuffleDeck(deck []*valueobject.Card) []*valueobject.Card {
	shuffledDeck := make([]*valueobject.Card, len(deck))
	copy(shuffledDeck, deck)
	for i := len(shuffledDeck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	}
	return shuffledDeck
}

var initialDeck = shuffleDeck(createDeck())

var ante = 10

func (t *Table) Ante() int {
	return ante
}

func (t *Table) CalculateTotalChips() int {
	totalChips := ante * len(t.players)
	for _, player := range t.players {
		totalChips += player.Chips()
	}
	return totalChips
}

// 勝ったプレイヤーに賞金を配る
func (t *Table) DistributeChips(winners []*entity.Player) {
	totalChips := t.CalculateTotalChips()
	for _, winner := range winners {
		winner.Win(totalChips / len(winners))
	}
}

// テーブル上のプレイヤーにカードを配る
func (t *Table) DealCards() {
	const numberOfCards = 5
	for i := 0; i < numberOfCards; i++ {
		for _, player := range t.players {
			card := t.deck[0]
			player.DrawCard(card)
			t.deck = t.deck[1:]
		}
	}
}

// テーブル上のプレイヤーの役を判定し、勝者を返す
func (t *Table) JudgeWinner() ([]*entity.Player, error) {
	// step1 まず、各プレイヤーの役を判定し、役の強さを比較する
	firstStepWinnerCandidates := []*entity.Player{}
	currentHand := "ハイカード"
	for _, player := range t.players {
		hands, err := player.JudgeHands()
		if err != nil {
			return nil, err
		}
		if entity.HandRankMap()[hands] > entity.HandRankMap()[currentHand] {
			firstStepWinnerCandidates = []*entity.Player{player}
			currentHand = hands
		} else if entity.HandRankMap()[hands] == entity.HandRankMap()[currentHand] {
			firstStepWinnerCandidates = append(firstStepWinnerCandidates, player)
		}
	}
	if len(firstStepWinnerCandidates) == 0 {
		return nil, fmt.Errorf("no winner candidates")
	}
	if len(firstStepWinnerCandidates) == 1 || currentHand == "ロイヤルストレートフラッシュ" {
		return firstStepWinnerCandidates, nil
	}
	/* step2
	同じ役の場合は、ハンドの「主要部」のランクの大小を比較する。
	ワンペアならペアになっているカード、
	ツーペアならペアになっているカードのうち強い方、
	スリーカード・フルハウスなら三枚組になっているカード、
	フォーカードなら四枚組になっているカード、
	ノーペア・ストレート・フラッシュ・ストレートフラッシュなら最も強いカード
	*/
	secondStepWinnerCandidates, err := entity.CompareMainPart(firstStepWinnerCandidates, currentHand)
	if err != nil {
		return nil, err
	}
	if len(secondStepWinnerCandidates) == 1 {
		return secondStepWinnerCandidates, nil
	}
	/* step3
	（ツーペアのみ）主要部が同じなら、準主要部のランクの大小を比較する。
	ペアになっているカードのうち弱い方を比較する。
	*/
	if currentHand == "ツーペア" {
		thirdStepWinnerCandidates, err := entity.CompareSubMainPart(secondStepWinnerCandidates)
		if err != nil {
			return nil, err
		}
		if len(thirdStepWinnerCandidates) == 1 {
			return thirdStepWinnerCandidates, nil
		}
		/* step4
		それが同じなら、キッカー（残ったカード）のうち最も高いランクのカードを比較する。
		以下順に二番目、三番目、四番目に高い札のランクを比べる。
		これらが全て同じ場合には、引き分けとみなされる。
		*/
		fourthStepWinnerCandidates, err := entity.CompareKicker(thirdStepWinnerCandidates, currentHand)
		if err != nil {
			return nil, err
		}
		return fourthStepWinnerCandidates, nil
	}

	fourthStepWinnerCandidates, err := entity.DetermineStrongestCardForOnePair(secondStepWinnerCandidates)
	if err != nil {
		return nil, err
	}

	return fourthStepWinnerCandidates, nil
}
