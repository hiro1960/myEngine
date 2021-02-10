// simframeの本体
package simframe

import (
	_ "core"
	"fmt"
	"model"
	"env"
	"logwriter"
)

func RunSim() {

	fmt.Printf("-- RunSim --\n")

	fmt.Printf(" Count = %d, DeltaTime = %f\n", TimeSet.Count, TimeSet.DeltaTime)

	// こちらでもObjListがスコープに入っているかの確認
	// 各オブジェクトの初期値
	fmt.Printf("-- init Object ---\n")
	for i, v := range ObjList {
		fmt.Printf("%d, %d, %s, %f, %f\n", i, v.GetId(), v.GetName(), v.GetPos(), v.GetVel())
	}

	// 計算の本体loop
	// 実験なので、回数を絞る
	var countDown int32 = TimeSet.Count

	for countDown > 0 {

		// 各オブジェクトの状態を確認
		fmt.Printf("count = %d\n", countDown)
		for _, v := range ObjList {
			v.Update()

			// 更新した値の確認
			objDB, _ := model.GetObjData(v.GetId())

			fmt.Printf(" now value: %d, %s, %f, %f\n", objDB.ID, objDB.Name, objDB.Pos, objDB.Vel)
		}

		// ObkDataDBの各オブジェクトのデータを更新したものにする
		model.UpdateObjData()

		countDown--
	}

	// singletonのテスト
	env := env.MyEnv.GetInstance()
	fmt.Printf(" at RunSim(), radius = %f \n", env.Radius)

	log := logwriter.LogWriter.GetInstance()
	log.WriteS("2nd line\n")

}
