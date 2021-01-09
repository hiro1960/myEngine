// modelパッケージ
// (ディレクトリ名とパッケージ名は同じ必要があるが、ファイル名は関係ない)
package model

import (
	"core"
)

// 本インターフェースを利用する全てのstructの持つメソッドをリストアップすること
// 公開するのは、このインターフェース名だけにする
type SimObject interface {
	GetId() int32
	GetName() string
	SetPos(pos core.Point)
	GetPos() core.Point
	GetWeight() float64 // どのstructが使うメソッドかコメントした方がいいか
	GetMass() float64
}

// 基底となるデータ構造
// class名自体は隠ぺいするため、小文字とする
type baseObject struct {
	ID   int32
	Name string
	Pos  core.Point
}

// IDを返す
func (b baseObject) GetId() int32 {
	return b.ID
}

// 名前を返す
func (b baseObject) GetName() string {
	return b.Name
}

// 位置をセット
// （引数はポインタでもいいが、とりあえず値渡しにしておく）
func (b *baseObject) SetPos(p core.Point) {
	for i, _ := range b.Pos.Value {
		b.Pos.Value[i] = p.Value[i]
	}
}

// 位置を返す(返す結果は外部に影響されたくないので、値渡しにしたい)
func (b *baseObject) GetPos() core.Point {
	return b.Pos
}
