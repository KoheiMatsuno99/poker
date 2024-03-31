package entity

import (
	"reflect"
	"testing"

	valueobject "github.com/KoheiMatsuno99/poker/domain/valueobject"
)

func TestCompareMainPart(t *testing.T) {
	type args struct {
		players []*Player
		hand    string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Player
		wantErr bool
	}{
		{
			name: "ワンペアの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "4"),
							valueobject.NewCard("diamond", "8"),
							valueobject.NewCard("spade", "J"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "3"),
							valueobject.NewCard("heart", "3"),
							valueobject.NewCard("diamond", "4"),
							valueobject.NewCard("diamond", "5"),
							valueobject.NewCard("club", "6"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("diamond", "3"),
							valueobject.NewCard("spade", "3"),
							valueobject.NewCard("heart", "4"),
							valueobject.NewCard("club", "10"),
							valueobject.NewCard("diamond", "K"),
						},
					},
				},
				hand: "ワンペア",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "3"),
						valueobject.NewCard("heart", "3"),
						valueobject.NewCard("diamond", "4"),
						valueobject.NewCard("diamond", "5"),
						valueobject.NewCard("club", "6"),
					},
				},
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("diamond", "3"),
						valueobject.NewCard("spade", "3"),
						valueobject.NewCard("heart", "4"),
						valueobject.NewCard("club", "10"),
						valueobject.NewCard("diamond", "K"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ツーペアの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "5"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("spade", "J"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "3"),
							valueobject.NewCard("heart", "3"),
							valueobject.NewCard("diamond", "5"),
							valueobject.NewCard("heart", "5"),
							valueobject.NewCard("club", "6"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("diamond", "3"),
							valueobject.NewCard("spade", "3"),
							valueobject.NewCard("heart", "4"),
							valueobject.NewCard("club", "4"),
							valueobject.NewCard("diamond", "K"),
						},
					},
				},
				hand: "ツーペア",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("heart", "2"),
						valueobject.NewCard("spade", "2"),
						valueobject.NewCard("club", "5"),
						valueobject.NewCard("spade", "5"),
						valueobject.NewCard("spade", "J"),
					},
				},
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "3"),
						valueobject.NewCard("heart", "3"),
						valueobject.NewCard("diamond", "5"),
						valueobject.NewCard("heart", "5"),
						valueobject.NewCard("club", "6"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "スリーカードの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("spade", "J"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("heart", "5"),
							valueobject.NewCard("club", "6"),
							valueobject.NewCard("club", "A"),
							valueobject.NewCard("diamond", "A"),
							valueobject.NewCard("heart", "A"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "4"),
							valueobject.NewCard("diamond", "Q"),
							valueobject.NewCard("diamond", "K"),
							valueobject.NewCard("spade", "K"),
							valueobject.NewCard("heart", "K"),
						},
					},
				},
				hand: "スリーカード",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("heart", "5"),
						valueobject.NewCard("club", "6"),
						valueobject.NewCard("club", "A"),
						valueobject.NewCard("diamond", "A"),
						valueobject.NewCard("heart", "A"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "フルハウスの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("diamond", "8"),
							valueobject.NewCard("spade", "8"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "A"),
							valueobject.NewCard("heart", "A"),
							valueobject.NewCard("diamond", "A"),
							valueobject.NewCard("heart", "3"),
							valueobject.NewCard("club", "3"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "10"),
							valueobject.NewCard("diamond", "10"),
							valueobject.NewCard("diamond", "K"),
							valueobject.NewCard("spade", "K"),
							valueobject.NewCard("heart", "K"),
						},
					},
				},
				hand: "フルハウス",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "3"),
						valueobject.NewCard("heart", "3"),
						valueobject.NewCard("club", "A"),
						valueobject.NewCard("diamond", "A"),
						valueobject.NewCard("heart", "A"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "フォーカードの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "4"),
							valueobject.NewCard("heart", "9"),
							valueobject.NewCard("club", "9"),
							valueobject.NewCard("diamond", "9"),
							valueobject.NewCard("spade", "9"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "K"),
							valueobject.NewCard("heart", "K"),
							valueobject.NewCard("diamond", "K"),
							valueobject.NewCard("heart", "K"),
							valueobject.NewCard("club", "8"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "A"),
							valueobject.NewCard("diamond", "A"),
							valueobject.NewCard("heart", "A"),
							valueobject.NewCard("spade", "A"),
							valueobject.NewCard("heart", "3"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("diamond", "2"),
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("diamond", "J"),
						},
					},
				},
				hand: "フォーカード",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("heart", "3"),
						valueobject.NewCard("club", "A"),
						valueobject.NewCard("diamond", "A"),
						valueobject.NewCard("heart", "A"),
						valueobject.NewCard("spade", "A"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ハイカードの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "3"),
							valueobject.NewCard("club", "7"),
							valueobject.NewCard("diamond", "J"),
							valueobject.NewCard("spade", "K"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "3"),
							valueobject.NewCard("heart", "6"),
							valueobject.NewCard("diamond", "8"),
							valueobject.NewCard("heart", "10"),
							valueobject.NewCard("club", "K"),
						},
					},
				},
				hand: "ハイカード",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("spade", "2"),
						valueobject.NewCard("heart", "3"),
						valueobject.NewCard("club", "7"),
						valueobject.NewCard("diamond", "J"),
						valueobject.NewCard("spade", "K"),
					},
				},
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "3"),
						valueobject.NewCard("heart", "6"),
						valueobject.NewCard("diamond", "8"),
						valueobject.NewCard("heart", "10"),
						valueobject.NewCard("club", "K"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ストレートの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "3"),
							valueobject.NewCard("club", "4"),
							valueobject.NewCard("diamond", "5"),
							valueobject.NewCard("spade", "6"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "7"),
							valueobject.NewCard("heart", "8"),
							valueobject.NewCard("diamond", "9"),
							valueobject.NewCard("heart", "10"),
							valueobject.NewCard("club", "J"),
						},
					},
				},
				hand: "ストレート",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "7"),
						valueobject.NewCard("heart", "8"),
						valueobject.NewCard("diamond", "9"),
						valueobject.NewCard("heart", "10"),
						valueobject.NewCard("club", "J"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "フラッシュの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("spade", "7"),
							valueobject.NewCard("spade", "9"),
							valueobject.NewCard("spade", "10"),
							valueobject.NewCard("spade", "Q"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "3"),
							valueobject.NewCard("club", "6"),
							valueobject.NewCard("club", "8"),
							valueobject.NewCard("club", "J"),
							valueobject.NewCard("club", "K"),
						},
					},
				},
				hand: "フラッシュ",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "3"),
						valueobject.NewCard("club", "6"),
						valueobject.NewCard("club", "8"),
						valueobject.NewCard("club", "J"),
						valueobject.NewCard("club", "K"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ストレートフラッシュの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("spade", "3"),
							valueobject.NewCard("spade", "4"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("spade", "6"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("club", "3"),
							valueobject.NewCard("club", "4"),
							valueobject.NewCard("club", "5"),
							valueobject.NewCard("club", "6"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("diamond", "A"),
							valueobject.NewCard("diamond", "2"),
							valueobject.NewCard("diamond", "3"),
							valueobject.NewCard("diamond", "4"),
							valueobject.NewCard("diamond", "5"),
						},
					},
				},
				hand: "ストレートフラッシュ",
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("spade", "2"),
						valueobject.NewCard("spade", "3"),
						valueobject.NewCard("spade", "4"),
						valueobject.NewCard("spade", "5"),
						valueobject.NewCard("spade", "6"),
					},
				},
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "2"),
						valueobject.NewCard("club", "3"),
						valueobject.NewCard("club", "4"),
						valueobject.NewCard("club", "5"),
						valueobject.NewCard("club", "6"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ロイヤルストレートフラッシュの場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "10"),
							valueobject.NewCard("spade", "J"),
							valueobject.NewCard("spade", "Q"),
							valueobject.NewCard("spade", "K"),
							valueobject.NewCard("spade", "A"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "10"),
							valueobject.NewCard("club", "J"),
							valueobject.NewCard("club", "Q"),
							valueobject.NewCard("club", "K"),
							valueobject.NewCard("club", "A"),
						},
					},
				},
				hand: "ロイヤルストレートフラッシュ",
			},
			wantErr: true,
		},
		{
			name: "プレイヤーの役が一致しない場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("spade", "3"),
							valueobject.NewCard("spade", "4"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("spade", "6"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "10"),
							valueobject.NewCard("club", "J"),
							valueobject.NewCard("club", "Q"),
							valueobject.NewCard("club", "K"),
							valueobject.NewCard("club", "A"),
						},
					},
				},
				hand: "ストレートフラッシュ",
			},
			wantErr: true,
		},
		{
			name: "プレイヤーの役が一致しない場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("spade", "3"),
							valueobject.NewCard("spade", "4"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("spade", "6"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "10"),
							valueobject.NewCard("club", "J"),
							valueobject.NewCard("club", "Q"),
							valueobject.NewCard("club", "K"),
							valueobject.NewCard("club", "A"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("diamond", "7"),
							valueobject.NewCard("club", "8"),
							valueobject.NewCard("spade", "9"),
							valueobject.NewCard("heart", "10"),
							valueobject.NewCard("diamond", "J"),
						},
					},
				},
				hand: "ストレート",
			},
			wantErr: true,
		},
		{
			name: "スリーカードで勝者が複数いる場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("spade", "J"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("diamond", "2"),
							valueobject.NewCard("heart", "5"),
							valueobject.NewCard("club", "6"),
						},
					},
				},
				hand: "スリーカード",
			},
			wantErr: true,
		},
		{
			name: "フルハウスで勝者が複数いる場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("diamond", "5"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("diamond", "2"),
							valueobject.NewCard("heart", "5"),
							valueobject.NewCard("club", "5"),
						},
					},
				},
				hand: "フルハウス",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareMainPart(tt.args.players, tt.args.hand)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareMainPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareMainPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareSubMainPart(t *testing.T) {
	type args struct {
		players []*Player
	}
	tests := []struct {
		name    string
		args    args
		want    []*Player
		wantErr bool
	}{
		{
			name: "ツーペアの場合/勝者が1人",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "5"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("spade", "J"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "3"),
							valueobject.NewCard("heart", "3"),
							valueobject.NewCard("diamond", "5"),
							valueobject.NewCard("heart", "5"),
							valueobject.NewCard("club", "6"),
						},
					},
				},
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "3"),
						valueobject.NewCard("heart", "3"),
						valueobject.NewCard("diamond", "5"),
						valueobject.NewCard("heart", "5"),
						valueobject.NewCard("club", "6"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ツーペアの場合/勝者が複数",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "3"),
							valueobject.NewCard("club", "3"),
							valueobject.NewCard("club", "5"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("spade", "J"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "3"),
							valueobject.NewCard("heart", "3"),
							valueobject.NewCard("diamond", "5"),
							valueobject.NewCard("heart", "5"),
							valueobject.NewCard("club", "6"),
						},
					},
				},
			},
			want: []*Player{
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("club", "3"),
						valueobject.NewCard("spade", "3"),
						valueobject.NewCard("club", "5"),
						valueobject.NewCard("spade", "5"),
						valueobject.NewCard("spade", "J"),
					},
				},
				{
					cards: []*valueobject.Card{
						valueobject.NewCard("heart", "3"),
						valueobject.NewCard("spade", "3"),
						valueobject.NewCard("diamond", "5"),
						valueobject.NewCard("heart", "5"),
						valueobject.NewCard("club", "6"),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ツーペア以外の場合",
			args: args{
				players: []*Player{
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("spade", "2"),
							valueobject.NewCard("heart", "2"),
							valueobject.NewCard("club", "2"),
							valueobject.NewCard("spade", "5"),
							valueobject.NewCard("heart", "5"),
						},
					},
					{
						cards: []*valueobject.Card{
							valueobject.NewCard("club", "3"),
							valueobject.NewCard("heart", "3"),
							valueobject.NewCard("diamond", "6"),
							valueobject.NewCard("heart", "6"),
							valueobject.NewCard("club", "6"),
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompareSubMainPart(tt.args.players)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompareSubMainPart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareSubMainPart() = %v, want %v", got, tt.want)
			}
		})
	}
}
