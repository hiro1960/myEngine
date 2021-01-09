// shipの実装本体
// (ディレクトリ名とパッケージ名は同じ必要があるが、ファイル名は関係ない)
package model

import (
	. "core"
)

// class名自体は隠ぺいするため、小文字とする
// baseObjectは同じmodelパッケージ内で定義してあるので、そのまま使える
type ship struct {
	*baseObject
	Mass float64 // 容積
}

// modelの他の実装のためのメソッドだが、定義は必要
func (v ship) GetWeight() float64 {
	return 0
}

func (v ship) GetMass() float64 {
	return v.Mass
}

// 隠ぺいしたクラスを生成するためのメソッド
func NewShip(id int32, name string, mass float64) *ship {
	var pos Point
	return &ship{&baseObject{id, name, pos}, mass}
}
