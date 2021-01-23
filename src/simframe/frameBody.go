// simframeの本体
package simframe

import (
	. "core"
	"fmt"
	_ "model"
)

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
	var countDoen int32 = 10

	for countDoen > 0 {

		// 各オブジェクトの状態を確認
		fmt.Printf("count = %d\n", countDoen)
		for _, v := range ObjList {
			v.Update()
		}

		countDoen--
	}

}
