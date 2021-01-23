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

	// シミュレーション本体
	simframe.RunSim()

}
