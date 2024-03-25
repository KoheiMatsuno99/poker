package entity

import (
	"reflect"
	"testing"

	valueobject "github.com/KoheiMatsuno99/poker/domain/valueobject"
)

func TestPlayer_sortCards(t *testing.T) {
	type fields struct {
		cards []*valueobject.Card
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*valueobject.Card
		wantErr bool
	}{
		{
			name: "ロイヤルストレートフラッシュ",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "J"),
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "K"),
					valueobject.NewCard("spade", "Q"),
					valueobject.NewCard("spade", "10"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "10"),
				valueobject.NewCard("spade", "J"),
				valueobject.NewCard("spade", "Q"),
				valueobject.NewCard("spade", "K"),
				valueobject.NewCard("spade", "A"),
			},
			wantErr: false,
		},
		{
			name: "ストレートフラッシュ/A,2,3,4,5",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("spade", "5"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "A"),
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("spade", "3"),
				valueobject.NewCard("spade", "4"),
				valueobject.NewCard("spade", "5"),
			},
			wantErr: false,
		},
		{
			name: "ストレートフラッシュ/Aを含まない",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "6"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("spade", "5"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("spade", "3"),
				valueobject.NewCard("spade", "4"),
				valueobject.NewCard("spade", "5"),
				valueobject.NewCard("spade", "6"),
			},
			wantErr: false,
		},
		{
			name: "フォーカード/A,A,A,A,2",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("diamond", "A"),
					valueobject.NewCard("club", "A"),
					valueobject.NewCard("heart", "A"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("club", "A"),
				valueobject.NewCard("diamond", "A"),
				valueobject.NewCard("heart", "A"),
				valueobject.NewCard("spade", "A"),
			},
			wantErr: false,
		},
		{
			name: "フォーカード/A,2,2,2,2",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("club", "2"),
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("diamond", "2"),
					valueobject.NewCard("heart", "2"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("club", "2"),
				valueobject.NewCard("diamond", "2"),
				valueobject.NewCard("heart", "2"),
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("spade", "A"),
			},
			wantErr: false,
		},
		{
			name: "フルハウス/A,A,A,2,2",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("diamond", "A"),
					valueobject.NewCard("club", "A"),
					valueobject.NewCard("heart", "A"),
					valueobject.NewCard("diamond", "2"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("diamond", "2"),
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("club", "A"),
				valueobject.NewCard("diamond", "A"),
				valueobject.NewCard("heart", "A"),
			},
			wantErr: false,
		},
		{
			name: "フルハウス/A,A,2,2,2",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("heart", "2"),
					valueobject.NewCard("diamond", "A"),
					valueobject.NewCard("diamond", "2"),
					valueobject.NewCard("club", "A"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("diamond", "2"),
				valueobject.NewCard("heart", "2"),
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("club", "A"),
				valueobject.NewCard("diamond", "A"),
			},
			wantErr: false,
		},
		{
			name: "ストレート/A,2,3,4,5",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("club", "5"),
					valueobject.NewCard("heart", "4"),
					valueobject.NewCard("diamond", "3"),
					valueobject.NewCard("spade", "A"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "A"),
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("diamond", "3"),
				valueobject.NewCard("heart", "4"),
				valueobject.NewCard("club", "5"),
			},
			wantErr: false,
		},
		{
			name: "ストレート/Aを含まない",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("club", "5"),
					valueobject.NewCard("heart", "4"),
					valueobject.NewCard("diamond", "3"),
					valueobject.NewCard("spade", "6"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("diamond", "3"),
				valueobject.NewCard("heart", "4"),
				valueobject.NewCard("club", "5"),
				valueobject.NewCard("spade", "6"),
			},
			wantErr: false,
		},
		{
			name: "スリーカード/A,A,A,2,3",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("diamond", "A"),
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("club", "A"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("spade", "3"),
				valueobject.NewCard("club", "A"),
				valueobject.NewCard("diamond", "A"),
				valueobject.NewCard("spade", "A"),
			},
			wantErr: false,
		},
		{
			name: "スリーカード/A,2,2,2,3",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("heart", "2"),
					valueobject.NewCard("club", "2"),
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("diamond", "2"),
					valueobject.NewCard("diamond", "3"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("club", "2"),
				valueobject.NewCard("diamond", "2"),
				valueobject.NewCard("heart", "2"),
				valueobject.NewCard("diamond", "3"),
				valueobject.NewCard("spade", "A"),
			},
			wantErr: false,
		},
		{
			name: "ツーペア/A,A,2,2,3",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("club", "A"),
					valueobject.NewCard("heart", "A"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("club", "2"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("club", "2"),
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("spade", "3"),
				valueobject.NewCard("club", "A"),
				valueobject.NewCard("heart", "A"),
			},
			wantErr: false,
		},
		{
			name: "ツーペア/A,2,2,3,3",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("club", "2"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("club", "A"),
					valueobject.NewCard("diamond", "3"),
					valueobject.NewCard("heart", "2"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("club", "2"),
				valueobject.NewCard("heart", "2"),
				valueobject.NewCard("diamond", "3"),
				valueobject.NewCard("spade", "3"),
				valueobject.NewCard("club", "A"),
			},
			wantErr: false,
		},
		{
			name: "ワンペア/A,A,2,3,4",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("heart", "2"),
					valueobject.NewCard("club", "A"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("diamond", "A"),
					valueobject.NewCard("spade", "3"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("heart", "2"),
				valueobject.NewCard("spade", "3"),
				valueobject.NewCard("spade", "4"),
				valueobject.NewCard("club", "A"),
				valueobject.NewCard("diamond", "A"),
			},
			wantErr: false,
		},
		{
			name: "ワンペア/A,2,3,4,4",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("club", "4"),
					valueobject.NewCard("heart", "4"),
					valueobject.NewCard("club", "3"),
					valueobject.NewCard("diamond", "A"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("club", "3"),
				valueobject.NewCard("club", "4"),
				valueobject.NewCard("heart", "4"),
				valueobject.NewCard("diamond", "A"),
			},
			wantErr: false,
		},
		{
			name: "ハイカード/A,2,3,Q,K",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("club", "Q"),
					valueobject.NewCard("heart", "K"),
					valueobject.NewCard("diamond", "3"),
					valueobject.NewCard("spade", "A"),
				},
			},
			want: []*valueobject.Card{
				valueobject.NewCard("spade", "2"),
				valueobject.NewCard("diamond", "3"),
				valueobject.NewCard("club", "Q"),
				valueobject.NewCard("heart", "K"),
				valueobject.NewCard("spade", "A"),
			},
			wantErr: false,
		},
		{
			name: "カードが5枚未満",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("club", "Q"),
					valueobject.NewCard("heart", "K"),
					valueobject.NewCard("diamond", "3"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				cards: tt.fields.cards,
			}
			err := p.sortCards()
			if err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}
			if !reflect.DeepEqual(p.cards, tt.want) {
				t.Errorf("got %v, want %v", p.cards, tt.want)
			}
		})
	}
}

func TestPlayer_JudgeHandsScore(t *testing.T) {
	type fields struct {
		cards []*valueobject.Card
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "ロイヤルストレートフラッシュ",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "J"),
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "K"),
					valueobject.NewCard("spade", "Q"),
					valueobject.NewCard("spade", "10"),
				},
			},
			want:    HandRankMap()["ロイヤルストレートフラッシュ"],
			wantErr: false,
		},
		{
			name: "ストレートフラッシュ/2,3,4,5,6",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "6"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("spade", "5"),
				},
			},
			want:    HandRankMap()["ストレートフラッシュ"],
			wantErr: false,
		},
		{
			name: "ストレートフラッシュ/A,2,3,4,5",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("spade", "5"),
				},
			},
			want:    HandRankMap()["ストレートフラッシュ"],
			wantErr: false,
		},
		{
			name: "フォーカード/A,A,A,A,2",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("diamond", "A"),
					valueobject.NewCard("club", "A"),
					valueobject.NewCard("heart", "A"),
				},
			},
			want:    HandRankMap()["フォーカード"],
			wantErr: false,
		},
		{
			name: "フルハウス/A,A,A,2,2",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("diamond", "A"),
					valueobject.NewCard("club", "A"),
					valueobject.NewCard("heart", "A"),
					valueobject.NewCard("diamond", "2"),
				},
			},
			want:    HandRankMap()["フルハウス"],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				cards: tt.fields.cards,
			}
			got, err := p.JudgeHandsScore()
			if (err != nil) != tt.wantErr {
				t.Errorf("Player.JudgeHands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Player.JudgeHands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_isRoyalStraightFlush(t *testing.T) {
	type fields struct {
		cards []*valueobject.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ロイヤルストレートフラッシュ",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "10"),
					valueobject.NewCard("spade", "J"),
					valueobject.NewCard("spade", "Q"),
					valueobject.NewCard("spade", "K"),
					valueobject.NewCard("spade", "A"),
				},
			},
			want: true,
		},
		{
			name: "ロイヤルストレートフラッシュでない/10,J,Q,K,Aの組み合わせでない",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "10"),
					valueobject.NewCard("spade", "J"),
					valueobject.NewCard("spade", "Q"),
					valueobject.NewCard("spade", "K"),
					valueobject.NewCard("spade", "2"),
				},
			},
			want: false,
		},
		{
			name: "ロイヤルストレートフラッシュでない/スートが異なる",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "10"),
					valueobject.NewCard("heart", "J"),
					valueobject.NewCard("spade", "Q"),
					valueobject.NewCard("spade", "K"),
					valueobject.NewCard("spade", "A"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				cards: tt.fields.cards,
			}
			if got := p.isRoyalStraightFlush(); got != tt.want {
				t.Errorf("Player.isRoyalStraightFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_isStraightFlush(t *testing.T) {
	type fields struct {
		cards []*valueobject.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ストレートフラッシュ/A,2,3,4,5",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("spade", "5"),
				},
			},
			want: true,
		},
		{
			name: "ストレートフラッシュ/Aを含まない",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("spade", "5"),
					valueobject.NewCard("spade", "6"),
				},
			},
			want: true,
		},
		{
			name: "ストレートフラッシュでない/フラッシュでない",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("heart", "3"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("spade", "5"),
					valueobject.NewCard("spade", "6"),
				},
			},
		},
		{
			name: "ストレートフラッシュでない/ストレートでない",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("spade", "5"),
					valueobject.NewCard("spade", "7"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				cards: tt.fields.cards,
			}
			if got := p.isStraightFlush(); got != tt.want {
				t.Errorf("Player.isStraightFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_isFourCard(t *testing.T) {
	type fields struct {
		cards      []*valueobject.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "フォーカード/A,A,A,A,2",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("club", "A"),
					valueobject.NewCard("diamond", "A"),
					valueobject.NewCard("heart", "A"),
					valueobject.NewCard("spade", "A"),
				},
			},
			want: true,
		},
		{
			name: "フォーカード/A,2,2,2,2",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("club", "2"),
					valueobject.NewCard("diamond", "2"),
					valueobject.NewCard("heart", "2"),
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "A"),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				cards:      tt.fields.cards,
			}
			if got := p.isFourCard(); got != tt.want {
				t.Errorf("Player.isFourCard() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestPlayer_isStraight(t *testing.T) {
	type fields struct {
		cards []*valueobject.Card
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ストレート/A,2,3,4,5",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "A"),
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("diamond", "4"),
					valueobject.NewCard("heart", "5"),
				},
			},
			want: true,
		},
		{
			name: "ストレート/Aを含まない",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "3"),
					valueobject.NewCard("spade", "4"),
					valueobject.NewCard("diamond", "5"),
					valueobject.NewCard("heart", "6"),
				},
			},
			want: true,
		},
		{
			name: "ストレートでない",
			fields: fields{
				cards: []*valueobject.Card{
					valueobject.NewCard("spade", "2"),
					valueobject.NewCard("spade", "J"),
					valueobject.NewCard("spade", "Q"),
					valueobject.NewCard("diamond", "K"),
					valueobject.NewCard("heart", "A"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				cards: tt.fields.cards,
			}
			if got := p.isStraight(); got != tt.want {
				t.Errorf("Player.isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

