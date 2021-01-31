// simexecのプロトタイプ　まず1ファイルで作り、その後packageに分ける
package main

// GOPATH環境変数の設定は必要（パッケージを読めなくなる）
import (
	_ "core"
	"env"
	"fmt"
	_ "model"
	"simframe"
	_ "test" // 試験用パッケージj
)

// modelのクラスはmodelパッケージ内に全て定義してある

// 実際のリスト (これはsimframeパッケージで定義した方がいいか？)
// sinframeに移動 @20210123
// var objList []model.SimObject

func main() {

	// 試験実施
	// TestMain()

	env := env.MyEnv.GetInstance()
	env.SetUp()

	fmt.Printf(" at main, radius = %f \n", env.Radius)

	// シナリオ読み込み
	simframe.Initialize()

	// objDataDBのテスト
	// var obj *simframe.ObjData
	obj, _ := simframe.GetObjData( 102 )

	fmt.Printf(" 102 vel= %f\n", obj.Vel)

	// errの使い方の確認
	pos, err := obj.GetPos(101)
	if err != nil {
		fmt.Printf(" GetPos: Error (%s)\n", err)
	}
	fmt.Printf("102 : %f\n", pos)

	// シミュレーション本体
	// simframe.RunSim()

}
