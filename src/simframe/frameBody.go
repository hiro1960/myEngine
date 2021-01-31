// simframeの本体
package simframe

import (
	. "core"
	"fmt"
	"errors"
	_ "model"
)

// シミュレーションに登場する全てのオブジェクトの位置、速度等データを管理する
// 現在時間のデータと、今回の時間で更新したデータを保持
// サイクルが進む時に、更新したデータが現在時間のデータとなる

// 各オブジェクトのデータの宣言
type ObjData struct {
	ID         int32
	Name       string
	Pos        Point // 現在位置
	vel        Point // 現在速度
	updatedPos Point // 更新後の位置
	updatedVel Point // 更新後の速度
}

// オブジェクトのデータベース本体
var ObjDataDB []ObjData

// IDでobjDataのitemを返す関数
func GetObjData(id int32) (ObjData, error) {
	for _, v := range ObjDataDB {
		if v.ID == id {
			return v, nil
		}
	}

	// return nil  // そもそもここに来ることを想定していない
	var b ObjData	// ダミーの戻り値
	return b, errors.New("Failure")
}

// objDataのitemから現在のPosを返す関数
func (b *ObjData) GetPos( id int32) Point {
	// errorを使う方法をまだ思いつかない
	obj, _ := GetObjData( id )
	return obj.Pos
}

// objDataのitemから現在のVelを返す関数

// objdatqのitemに更新後のPosをセットする関数

// objDataのitemに更新後のVelをセットする関数

// objDataの更新後データをサイクルが進んだ時に現在データにする関数

func RunSim() {

	fmt.Printf("-- RunSim --\n")

	// こちらでもObjListがスコープに入っているかの確認
	// 各オブジェクトの初期値
	fmt.Printf("-- init Object ---\n")
	for i, v := range ObjList {
		var pos Point = v.GetPos()
		var vel Point = v.GetVel()
		fmt.Printf("%d, %d, %s, %f, %f\n", i, v.GetId(), v.GetName(), pos, vel)
	}

	// 計算の本体loop
	// 実験なので、回数を絞る
	var countDown int32 = 10

	for countDown > 0 {

		// 各オブジェクトの状態を確認
		fmt.Printf("count = %d\n", countDown)
		for _, v := range ObjList {
			v.Update()
		}

		countDown--
	}

}
