// modelパッケージ
// (ディレクトリ名とパッケージ名は同じ必要があるが、ファイル名は関係ない)
package model

import (
	. "core"
	"fmt"
)

// 本インターフェースを利用する全てのstructの持つメソッドをリストアップすること
// 公開するのは、このインターフェース名だけにする
type SimObject interface {
	GetId() int32
	GetName() string
	SetPos(pos Point)
	GetPos() Point
	SetVel(vel Point)
	GetVel() Point
	GetWeight() float64 // どのstructが使うメソッドかコメントした方がいいか
	GetMass() float64
	Update()	// オブジェクトの更新処理位
}

// 基底となるデータ構造
// class名自体は隠ぺいするため、小文字とする
type baseObject struct {
	ID   int32
	Name string
	Pos  Point // オブジェクトの位置
	Vel  Point // オブジェクトの速度
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
func (b *baseObject) SetPos(p Point) {
	for i, _ := range b.Pos.Value {
		b.Pos.Value[i] = p.Value[i]
	}
}

// 位置を返す(返す結果は外部に影響されたくないので、値渡しにしたい)
func (b *baseObject) GetPos() Point {
	return b.Pos
}

// 速度をセット
// （引数はポインタでもいいが、とりあえず値渡しにしておく）
func (b *baseObject) SetVel(p Point) {
	for i, _ := range b.Pos.Value {
		b.Vel.Value[i] = p.Value[i]
	}
}

// 位置を返す(返す結果は外部に影響されたくないので、値渡しにしたい)
func (b *baseObject) GetVel() Point {
	return b.Vel
}

// 基本のmodelの更新処理、現状を出力するだけ
func (b *baseObject) Update() {
	fmt.Printf("-- Update: id = %d\n", b.ID)

	var pos Point = b.GetPos()
	var vel Point = b.GetVel()
	fmt.Printf(" %d, %s, %f, %f\n", b.GetId(), b.GetName(), pos, vel)

}