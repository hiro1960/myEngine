// simframeの本体
package simframe

import (
	_ "core"
	"env"
	"fmt"
	"logwriter"
	"model"
)

func RunSim() {

	fmt.Printf("-- RunSim --\n")

	fmt.Printf(" Count = %d, DeltaTime = %f\n", TimeSet.Count, TimeSet.DeltaTime)

	log := logwriter.LogWriter.GetInstance()
	log.WriteS(fmt.Sprintf("Count,%d,DeltaTime,%f\n", TimeSet.Count, TimeSet.DeltaTime))

	// 計算の本体loop
	var count int32

	for count = 0; count < TimeSet.Count; count++ {

		// 各オブジェクトの状態を確認
		log.WriteS(fmt.Sprintf("time,%d,%f\n", count, float64(count)*TimeSet.DeltaTime))
		for _, v := range ObjList {

			v.Update()

		}

		// ObkDataDBの各オブジェクトのデータを更新したものにする
		model.UpdateObjData()

	}

	// singletonのテスト
	env := env.MyEnv.GetInstance()
	fmt.Printf(" at RunSim(), radius = %f \n", env.Radius)

}
