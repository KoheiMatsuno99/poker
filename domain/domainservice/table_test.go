package domainservice

import (
	"reflect"
	"testing"

	"github.com/KoheiMatsuno99/poker/domain/entity"
	"github.com/KoheiMatsuno99/poker/domain/valueobject"
)

func TestTable_JudgeWinner(t *testing.T) {
	type fields struct {
		players []*entity.Player
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Player
		wantErr bool
	}{
		{
			name: "役のみで勝者が1人に決まる場合",
			fields: fields{
				players: []*entity.Player{
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("club", "3"))
						player.DrawCard(valueobject.NewCard("club", "4"))
						player.DrawCard(valueobject.NewCard("club", "5"))
						player.DrawCard(valueobject.NewCard("club", "6"))
						player.DrawCard(valueobject.NewCard("club", "7"))
						return player
					}(),
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("diamond", "5"))
						player.DrawCard(valueobject.NewCard("club", "5"))
						player.DrawCard(valueobject.NewCard("club", "8"))
						player.DrawCard(valueobject.NewCard("heart", "J"))
						player.DrawCard(valueobject.NewCard("spade", "Q"))
						return player
					}(),
				},
			},
			want: []*entity.Player{
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("club", "3"))
					player.DrawCard(valueobject.NewCard("club", "4"))
					player.DrawCard(valueobject.NewCard("club", "5"))
					player.DrawCard(valueobject.NewCard("club", "6"))
					player.DrawCard(valueobject.NewCard("club", "7"))
					return player
				}(),
			},
			wantErr: false,
		},
		{
			name: "役のみで勝者が決まる場合/ロイヤルストレートフラッシュ同士",
			fields: fields{
				players: []*entity.Player{
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("club", "10"))
						player.DrawCard(valueobject.NewCard("club", "J"))
						player.DrawCard(valueobject.NewCard("club", "Q"))
						player.DrawCard(valueobject.NewCard("club", "K"))
						player.DrawCard(valueobject.NewCard("club", "A"))
						return player
					}(),
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("diamond", "10"))
						player.DrawCard(valueobject.NewCard("diamond", "J"))
						player.DrawCard(valueobject.NewCard("diamond", "Q"))
						player.DrawCard(valueobject.NewCard("diamond", "K"))
						player.DrawCard(valueobject.NewCard("diamond", "A"))
						return player
					}(),
				},
			},
			want: []*entity.Player{
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("club", "10"))
					player.DrawCard(valueobject.NewCard("club", "J"))
					player.DrawCard(valueobject.NewCard("club", "Q"))
					player.DrawCard(valueobject.NewCard("club", "K"))
					player.DrawCard(valueobject.NewCard("club", "A"))
					return player
				}(),
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("diamond", "10"))
					player.DrawCard(valueobject.NewCard("diamond", "J"))
					player.DrawCard(valueobject.NewCard("diamond", "Q"))
					player.DrawCard(valueobject.NewCard("diamond", "K"))
					player.DrawCard(valueobject.NewCard("diamond", "A"))
					return player
				}(),
			},
			wantErr: false,
		},
		{
			name: "主要部で勝者が1人に決まる場合",
			fields: fields{
				players: []*entity.Player{
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("club", "3"))
						player.DrawCard(valueobject.NewCard("heart", "3"))
						player.DrawCard(valueobject.NewCard("spade", "3"))
						player.DrawCard(valueobject.NewCard("club", "5"))
						player.DrawCard(valueobject.NewCard("club", "7"))
						return player
					}(),
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("club", "4"))
						player.DrawCard(valueobject.NewCard("heart", "4"))
						player.DrawCard(valueobject.NewCard("spade", "4"))
						player.DrawCard(valueobject.NewCard("club", "5"))
						player.DrawCard(valueobject.NewCard("club", "7"))
						return player
					}(),
				},
			},
			want: []*entity.Player{
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("club", "4"))
					player.DrawCard(valueobject.NewCard("heart", "4"))
					player.DrawCard(valueobject.NewCard("spade", "4"))
					player.DrawCard(valueobject.NewCard("club", "5"))
					player.DrawCard(valueobject.NewCard("club", "7"))
					return player
				}(),
			},
			wantErr: false,
		},
		{
			name: "準主要部で勝者が1人に決まる場合",
			fields: fields{
				players: []*entity.Player{
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("club", "4"))
						player.DrawCard(valueobject.NewCard("heart", "4"))
						player.DrawCard(valueobject.NewCard("diamond", "9"))
						player.DrawCard(valueobject.NewCard("spade", "9"))
						player.DrawCard(valueobject.NewCard("club", "Q"))
						return player
					}(),
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("club", "2"))
						player.DrawCard(valueobject.NewCard("heart", "2"))
						player.DrawCard(valueobject.NewCard("club", "9"))
						player.DrawCard(valueobject.NewCard("heart", "9"))
						player.DrawCard(valueobject.NewCard("spade", "Q"))
						return player
					}(),
				},
			},
			want: []*entity.Player{
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("club", "4"))
					player.DrawCard(valueobject.NewCard("heart", "4"))
					player.DrawCard(valueobject.NewCard("diamond", "9"))
					player.DrawCard(valueobject.NewCard("spade", "9"))
					player.DrawCard(valueobject.NewCard("club", "Q"))
					return player
				}(),
			},
			wantErr: false,
		},
		{
			name: "準主要部でも勝者が決まらない場合/ツーペア",
			fields: fields{
				players: []*entity.Player{
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("club", "4"))
						player.DrawCard(valueobject.NewCard("heart", "4"))
						player.DrawCard(valueobject.NewCard("diamond", "9"))
						player.DrawCard(valueobject.NewCard("spade", "9"))
						player.DrawCard(valueobject.NewCard("club", "Q"))
						return player
					}(),
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("diamond", "4"))
						player.DrawCard(valueobject.NewCard("spade", "4"))
						player.DrawCard(valueobject.NewCard("club", "9"))
						player.DrawCard(valueobject.NewCard("heart", "9"))
						player.DrawCard(valueobject.NewCard("spade", "Q"))
						return player
					}(),
				},
			},
			want: []*entity.Player{
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("club", "4"))
					player.DrawCard(valueobject.NewCard("heart", "4"))
					player.DrawCard(valueobject.NewCard("diamond", "9"))
					player.DrawCard(valueobject.NewCard("spade", "9"))
					player.DrawCard(valueobject.NewCard("club", "Q"))
					return player
				}(),
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("diamond", "4"))
					player.DrawCard(valueobject.NewCard("spade", "4"))
					player.DrawCard(valueobject.NewCard("club", "9"))
					player.DrawCard(valueobject.NewCard("heart", "9"))
					player.DrawCard(valueobject.NewCard("spade", "Q"))
					return player
				}(),
			},
			wantErr: false,
		},
		{
			name: "準主要部でも勝者が決まらない場合/ワンペア",
			fields: fields{
				players: []*entity.Player{
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("club", "4"))
						player.DrawCard(valueobject.NewCard("heart", "4"))
						player.DrawCard(valueobject.NewCard("diamond", "9"))
						player.DrawCard(valueobject.NewCard("spade", "J"))
						player.DrawCard(valueobject.NewCard("club", "K"))
						return player
					}(),
					func() *entity.Player {
						player := &entity.Player{}
						player.DrawCard(valueobject.NewCard("diamond", "4"))
						player.DrawCard(valueobject.NewCard("spade", "4"))
						player.DrawCard(valueobject.NewCard("club", "9"))
						player.DrawCard(valueobject.NewCard("heart", "J"))
						player.DrawCard(valueobject.NewCard("spade", "K"))
						return player
					}(),
				},
			},
			want: []*entity.Player{
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("club", "4"))
					player.DrawCard(valueobject.NewCard("heart", "4"))
					player.DrawCard(valueobject.NewCard("diamond", "9"))
					player.DrawCard(valueobject.NewCard("spade", "J"))
					player.DrawCard(valueobject.NewCard("club", "K"))
					return player
				}(),
				func() *entity.Player {
					player := &entity.Player{}
					player.DrawCard(valueobject.NewCard("diamond", "4"))
					player.DrawCard(valueobject.NewCard("spade", "4"))
					player.DrawCard(valueobject.NewCard("club", "9"))
					player.DrawCard(valueobject.NewCard("heart", "J"))
					player.DrawCard(valueobject.NewCard("spade", "K"))
					return player
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				players: tt.fields.players,
			}
			got, err := tr.JudgeWinner()
			if (err != nil) != tt.wantErr {
				t.Errorf("Table.JudgeWinner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.JudgeWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
