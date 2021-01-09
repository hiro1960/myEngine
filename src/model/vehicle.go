// vehicleの実装本体
// (ディレクトリ名とパッケージ名は同じ必要があるが、ファイル名は関係ない)
package model

import (
	"core"
)

// class名自体は隠ぺいするため、小文字とする
// baseObjectは同じmodelパッケージ内で定義してあるので、そのまま使える
type vehicle struct {
	*baseObject
	Weight float64 // 重量[kg]
}

func (v vehicle) GetWeight() float64 {
	return v.Weight
}

// modelの他の実装のためのメソッドだが、定義は必要
func (v vehicle) GetMass() float64 {
	return 0
}

// 隠ぺいしたクラスを生成するためのメソッド
func NewVehicle(id int32, name string, weight float64) *vehicle {
	var pos core.Point
	return &vehicle{&baseObject{id, name, pos}, weight}
}
