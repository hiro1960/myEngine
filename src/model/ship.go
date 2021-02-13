// shipの実装本体
// (ディレクトリ名とパッケージ名は同じ必要があるが、ファイル名は関係ない)
package model

import (
	. "core"
	_ "fmt"
	"model/ship"
)

// class名自体は隠ぺいするため、小文字とする
// ship本体の実装を別ディレクトリに作ろうと思うと、当該パッケージ名は別名にする必要がある（例：shiBody）
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
	var vel Point
	return &ship{&baseObject{id, name, pos, vel}, mass}
}

// ship独自の更新関数を定義する
func (b *ship) Update() {
	// 102のJeepを決め打ちで情報を得る (テスト用)
	// obj, _ := GetObjData(102)
	// fmt.Printf("in ship Jeep data: %d %s %f %f\n", obj.ID, obj.Name, obj.Pos, obj.Vel)

	// 位置に単純に速度ベクトルを加算するだけ
	b.Pos.Add(&b.Vel)

	// 定義は、model/ship以下にあるが、前頭句はパッケージ名を使う
	shipModel.UpdateVel(&b.Vel)

	// model/shipにて定義した変数をパッケージ名で使える。これでパッケージ内の関数とやり取りが行える。
	shipModel.VV.Vel = 0

	// ObjDataの更新（他のオブジェクトからの参照用に更新しておく）
	objDB, err := GetObjData(b.ID)
	if err == nil {
		*objDB.UpdatedPos = b.Pos
		*objDB.UpdatedVel = b.Vel
	}

}
