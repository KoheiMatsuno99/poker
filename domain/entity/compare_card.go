package entity

import (
	"fmt"

	valueobject "github.com/KoheiMatsuno99/poker/domain/valueobject"
)

// ハンドの主要部を比較する
func CompareMainPart(players []*Player, hand string) ([]*Player, error) {
	for _, player := range players {
		playerHand, err := player.JudgeHands()
		if err != nil {
			return nil, err
		}
		if playerHand != hand {
			return nil, fmt.Errorf("player hand is not %s", hand)
		}
	}
	switch hand {
	case "ワンペア":
		return DetermineStrongestCardForOnePair(players)
	case "ツーペア":
		return DetermineStrongestCardForTwoPair(players, "main")
	case "スリーカード":
		return DetermineStrongestCardPlayerForThreeOfAKind(players)
	case "フルハウス":
		return DetermineStrongestPlayerForFullHouse(players)
	case "フォーカード":
		return DetermineStrongestCardPlayerForFourOfAKind(players)
	case "ハイカード", "ストレート", "フラッシュ", "ストレートフラッシュ":
		return DetermineStrongestCardPlayerForSpecificHands(players)
	default:
		// ロイヤルストレートフラッシュ同士は必ず引き分けになるので比較不要
		return nil, fmt.Errorf("hand is not valid")
	}
}

// ハンドの準主要部を比較する
func CompareSubMainPart(players []*Player) ([]*Player, error) {
	for _, player := range players {
		playerHand, err := player.JudgeHands()
		if err != nil {
			return nil, err
		}
		if playerHand != "ツーペア" {
			return nil, fmt.Errorf("player hand is not two pair")
		}
	}
	return DetermineStrongestCardForTwoPair(players, "sub")
}

// キッカーを比較する
func CompareKicker(players []*Player, hand string) ([]*Player, error) {
	for _, player := range players {
		playerHand, err := player.JudgeHands()
		if err != nil {
			return nil, err
		}
		if playerHand != hand {
			return nil, fmt.Errorf("player hand is not %s", hand)
		}
	}
	switch hand {
	case "ワンペア":
		return DetermineStrongestCardForOnePairKicker(players)
	case "ツーペア":
		return DetermineStrongestCardForTwoPairKicker(players)
	default:
		return nil, fmt.Errorf("%s does not have kicker", hand)
	}
}

func DetermineStrongestCardForOnePair(players []*Player) ([]*Player, error) {
	winnerCandidate := []*Player{}
	maxCardRank := 0
	for _, player := range players {
		onePair, err := player.SeparateOnePairAndOtherCards()
		if err != nil {
			return nil, err
		}
		if valueobject.ValueRankMap()[onePair[0][len(onePair)-1].Value()] > maxCardRank {
			winnerCandidate = []*Player{player}
			maxCardRank = valueobject.ValueRankMap()[onePair[0][len(onePair)-1].Value()]
		} else if valueobject.ValueRankMap()[onePair[0][len(onePair)-1].Value()] == maxCardRank {
			winnerCandidate = append(winnerCandidate, player)
		}
	}
	return winnerCandidate, nil
}

func DetermineStrongestCardForTwoPair(players []*Player, part string) ([]*Player, error) {
	winnerCandidate := []*Player{}
	maxCardRank := 0
	for _, player := range players {
		twoPairs, err := player.SeparateTwoPairAndOtherCards()
		if err != nil {
			return nil, err
		}
		twoPairsLength := len(twoPairs)
		firstPairLength := len(twoPairs[0])
		switch part {
		case "main":
			if valueobject.ValueRankMap()[twoPairs[0][firstPairLength-1].Value()] > maxCardRank {
				winnerCandidate = []*Player{player}
				maxCardRank = valueobject.ValueRankMap()[twoPairs[0][firstPairLength-1].Value()]
			} else if valueobject.ValueRankMap()[twoPairs[0][firstPairLength-1].Value()] == maxCardRank {
				winnerCandidate = append(winnerCandidate, player)
			}
		case "sub":
			if valueobject.ValueRankMap()[twoPairs[twoPairsLength-2][0].Value()] > maxCardRank {
				winnerCandidate = []*Player{player}
				maxCardRank = valueobject.ValueRankMap()[twoPairs[twoPairsLength-2][0].Value()]
				fmt.Println(maxCardRank)
			} else if valueobject.ValueRankMap()[twoPairs[twoPairsLength-2][0].Value()] == maxCardRank {
				winnerCandidate = append(winnerCandidate, player)
			}
		default:
			return nil, fmt.Errorf("part is not valid")
		}
	}
	return winnerCandidate, nil
}

func DetermineStrongestCardPlayerForThreeOfAKind(players []*Player) ([]*Player, error) {
	winnerCandidate := []*Player{}
	maxCardRank := 0
	for _, player := range players {
		threeCards, err := player.SeparateThreeOfAKindsAndOtherCards()
		if err != nil {
			return nil, err
		}
		if valueobject.ValueRankMap()[threeCards[0][len(threeCards)-1].Value()] > maxCardRank {
			winnerCandidate = []*Player{player}
			maxCardRank = valueobject.ValueRankMap()[threeCards[0][len(threeCards)-1].Value()]
		} else if valueobject.ValueRankMap()[threeCards[0][len(threeCards)-1].Value()] == maxCardRank {
			// 52枚で遊ぶ場合、同じランクのスリーカードが複数存在することはない
			return nil, fmt.Errorf("three of a kind is duplicated")
		}
	}
	return winnerCandidate, nil
}

func DetermineStrongestPlayerForFullHouse(players []*Player) ([]*Player, error) {
	winnerCandidate := []*Player{}
	maxCardRank := 0
	for _, player := range players {
		fullHouse, err := player.SeparateThreeOfAKindAndOnePair()
		if err != nil {
			return nil, err
		}
		if valueobject.ValueRankMap()[fullHouse[0][len(fullHouse)-1].Value()] > maxCardRank {
			winnerCandidate = []*Player{player}
			maxCardRank = valueobject.ValueRankMap()[fullHouse[0][len(fullHouse)-1].Value()]
		} else if valueobject.ValueRankMap()[fullHouse[0][len(fullHouse)-1].Value()] == maxCardRank {
			// 52枚で遊ぶ場合、同じランクのフルハウスが複数存在することはない
			return nil, fmt.Errorf("full house is duplicated")
		}
	}
	return winnerCandidate, nil

}

func DetermineStrongestCardPlayerForFourOfAKind(players []*Player) ([]*Player, error) {
	winnerCandidate := []*Player{}
	maxCardRank := 0
	for _, player := range players {
		fourCards, err := player.SeparateFourOfAKindAndOtherCard()
		if err != nil {
			return nil, err
		}
		if valueobject.ValueRankMap()[fourCards[0][len(fourCards)-1].Value()] > maxCardRank {
			winnerCandidate = []*Player{player}
			maxCardRank = valueobject.ValueRankMap()[fourCards[0][len(fourCards)-1].Value()]
		} else if valueobject.ValueRankMap()[fourCards[0][len(fourCards)-1].Value()] == maxCardRank {
			// 52枚で遊ぶ場合、同じランクのフォーカードが複数存在することはない
			return nil, fmt.Errorf("four of a kind is duplicated")
		}
	}
	return winnerCandidate, nil
}

func DetermineStrongestCardPlayerForSpecificHands(players []*Player) ([]*Player, error) {
	winnerCandidate := []*Player{}
	maxCardRank := 0
	for _, player := range players {
		if valueobject.ValueRankMap()[player.Cards()[len(player.Cards())-1].Value()] > maxCardRank {
			winnerCandidate = []*Player{player}
			maxCardRank = valueobject.ValueRankMap()[player.Cards()[len(player.Cards())-1].Value()]
		} else if valueobject.ValueRankMap()[player.Cards()[len(player.Cards())-1].Value()] == maxCardRank {
			winnerCandidate = append(winnerCandidate, player)
		}
	}
	return winnerCandidate, nil
}

func DetermineStrongestCardForOnePairKicker(players []*Player) ([]*Player, error) {
	winnerCandidate := []*Player{}
	maxCardRank := 0
	kickerLength := 3
	for i := kickerLength - 1; i >= 0; i-- {
		for _, player := range players {
			onePair, err := player.SeparateOnePairAndOtherCards()
			if err != nil {
				return nil, err
			}
			if valueobject.ValueRankMap()[onePair[len(onePair)-1][i].Value()] > maxCardRank {
				winnerCandidate = []*Player{player}
				maxCardRank = valueobject.ValueRankMap()[onePair[len(onePair)-1][i].Value()]
			} else if valueobject.ValueRankMap()[onePair[len(onePair)-1][i].Value()] == maxCardRank {
				winnerCandidate = append(winnerCandidate, player)
			}
		}
		if len(winnerCandidate) == 1 {
			return winnerCandidate, nil
		}
	}
	return winnerCandidate, nil
}

func DetermineStrongestCardForTwoPairKicker(players []*Player) ([]*Player, error) {
	winnerCandidate := []*Player{}
	maxCardRank := 0
	kickerLength := 1
	for i := kickerLength - 1; i >= 0; i-- {
		for _, player := range players {
			twoPairs, err := player.SeparateTwoPairAndOtherCards()
			if err != nil {
				return nil, err
			}
			if valueobject.ValueRankMap()[twoPairs[len(twoPairs)-1][i].Value()] > maxCardRank {
				winnerCandidate = []*Player{player}
				maxCardRank = valueobject.ValueRankMap()[twoPairs[len(twoPairs)-1][i].Value()]
			} else if valueobject.ValueRankMap()[twoPairs[len(twoPairs)-1][i].Value()] == maxCardRank {
				winnerCandidate = append(winnerCandidate, player)
			}
		}
		if len(winnerCandidate) == 1 {
			return winnerCandidate, nil
		}
	}
	return winnerCandidate, nil
}
