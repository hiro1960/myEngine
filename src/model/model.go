// modelパッケージ
// (ディレクトリ名とパッケージ名は同じ必要があるが、ファイル名は関係ない)
package model

import (
	. "core"
	"fmt"
	"errors"
)

// シミュレーションに登場する全てのオブジェクトの位置、速度等データを管理する
// 現在時間のデータと、今回の時間で更新したデータを保持
// サイクルが進む時に、更新したデータが現在時間のデータとなる

// 各オブジェクトのデータの宣言
type ObjData struct {
	ID         int32
	Name       string
	// 構造体はポインタで定義しておかないと、参照時に自動的に新規生成が行われてしまうので注意が必要
	Pos        *Point // 現在位置
	Vel        *Point // 現在速度 頭文字が大文字でないと、外部で参照できない
	UpdatedPos *Point // 更新後の位置
	UpdatedVel *Point // 更新後の速度
}

// シミュレーションに登場する全てのオブジェクトの位置、速度等データを管理する
// 現在時間のデータと、今回の時間で更新したデータを保持
// サイクルが進む時に、更新したデータが現在時間のデータとなる

// オブジェクトのデータベース本体
var ObjDataDB []ObjData

// IDでobjDataのitemを返す関数
func GetObjData(id int32) (*ObjData, error) {
	for _, v := range ObjDataDB {
		if v.ID == id {
			return &v, nil	// 正常に見つかった
		}
	}

	// return nil  // そもそもここに来ることを想定していない
	var b ObjData	// ダミーの戻り値
	return &b, errors.New("Failure")
}

// objDataの更新後データをサイクルが進んだ時に現在データにする関数
func UpdateObjData() {
	for _, v := range ObjDataDB {
		// v.Pos = v.updatedPos
		v.setPos(*v.UpdatedPos)
		v.setVel(*v.UpdatedVel)

		fmt.Printf("in UpdateObjData(): id = %d, pos = %f, vel = %f\n", v.ID, v.Pos, v.Vel)
	}
}

func (v *ObjData) setPos( b Point) {
	for i, _ := range b.Value {
		v.Pos.Value[i] = b.Value[i]
	}
}

func (v *ObjData) setVel( b Point) {
	for i, _ := range b.Value {
		v.Vel.Value[i] = b.Value[i]
	}
}

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

	// 位置に単純に速度ベクトルを加算するだけ
	b.Pos.Add(&b.Vel)

	// ObjDataの更新（他のオブジェクトからの参照用に更新しておく）
	objDB, err := GetObjData( b.ID )
	if err == nil {
		*objDB.UpdatedPos = b.Pos
		*objDB.UpdatedVel = b.Vel
		// fmt.Printf("no error GetObjData()\n")
	}

	fmt.Printf(" in update(): %d, %s, %f, %f\n", b.GetId(), b.GetName(), b.Pos, b.Vel)

}