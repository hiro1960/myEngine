// simexecのプロトタイプ　まず1ファイルで作り、その後packageに分ける
package main

// GOPATH環境変数の設定は必要（パッケージを読めなくなる）
import (
	_ "core"
	"env"
	"fmt"
	"model"
	"simframe"
	"logwriter"
	_ "test" // 試験用パッケージj
)

// modelのクラスはmodelパッケージ内に全て定義してある

func main() {

	// 試験実施
	// TestMain()

	env := env.MyEnv.GetInstance()
	env.SetUp()

	fmt.Printf(" at main, radius = %f \n", env.Radius)

	// log用ファイルの準備
	log := logwriter.LogWriter.GetInstance()

	erro := log.Open("./output.log")
	if erro != nil {
		panic(erro)
	}

	log.WriteS("1st line\n")

	// シナリオ読み込み
	simframe.Initialize()

	// objDataDBのテスト
	// var obj *simframe.ObjData
	obj, _ := model.GetObjData( 102 )

	fmt.Printf(" 102 vel= %f\n", obj.Vel)

	// errの使い方の確認
	obj, err := model.GetObjData(101)
	if err != nil {
		fmt.Printf(" GetObjData: Error (%s)\n", err)
	}
	fmt.Printf("101 : %v\n", obj)

	// シミュレーション本体
	simframe.RunSim()

	// log用ファイルのクローズ
	log.Close()

}
