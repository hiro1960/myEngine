// simexecのプロトタイプ　まず1ファイルで作り、その後packageに分ける
package main

// GOPATH環境変数の設定は必要（パッケージを読めなくなる）
import (
	_ "core"
	"fmt"
	_ "model"
	_ "test" // 試験用パッケージj
	"env"
	"simframe"
)

// modelのクラスはmodelパッケージ内に全て定義してある

// 実際のリスト (これはsimframeパッケージで定義した方がいいか？) 
// sinframeに移動 @20210123
// var objList []model.SimObject

func main() {
	// var obj model.SimObject

	// var ObjList []model.SimObject	// testパッケージ内で定義している

	// obj = model.NewVehicle(101, "Tank", 124.0)
	// ObjList = append(ObjList, obj)
	// fmt.Printf(" ID = %d\n", obj.GetId())
	// fmt.Printf(" Name = %s\n", obj.GetName())
	// fmt.Printf(" Weight = %f\n", obj.GetWeight())

	// fmt.Printf(" objList length = %d\n", len(ObjList))

	// obj = model.NewShip(102, "Aegis", 201.0)
	// ObjList = append(ObjList, obj)
	// fmt.Printf(" ID = %d\n", obj.GetId())
	// fmt.Printf(" Name = %s\n", obj.GetName())
	// fmt.Printf(" Mass = %f\n", obj.GetMass())

	// fmt.Printf(" objList length = %d\n", len(ObjList))

	// for _, s := range ObjList {
	// 	fmt.Printf(" ID = %d\n", s.GetId())
	// 	fmt.Printf(" Name = %s\n", s.GetName())
	// }

	// 試験実施
	// TestMain()

	env := env.MyEnv.GetInstance()
	env.SetUp()

	fmt.Printf(" at main, radius = %f \n", env.Radius)

	simframe.Initialize()


}
