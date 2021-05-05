// simexecのプロトタイプ　まず1ファイルで作り、その後packageに分ける
package main

// GOPATH環境変数の設定は必要（パッケージを読めなくなる）
import (
	. "core"
	"env"
	"fmt"
	"logwriter"
	_ "model"
	"simframe"
	_ "test" // 試験用パッケージj
)

// modelのクラスはmodelパッケージ内に全て定義してある

func main() {

	// 試験実施
	// TestMain()

	env := env.MyEnv.GetInstance()
	env.SetUp()
	// この後、Pointが持つ地球半径を初期設定する

	fmt.Printf(" at main, radius = %f \n", env.Radius)

	// pos := core.NewPoint([3]float64{})
	// pos.SetRadius(env.Radius)
	// TestPoint()
	// TestCSV()
	// TestCSV2()
	// TestMat()
	p := new(Quartenion)
	p.Initialize(0.0, 0.0, 0.0)

	// log用ファイルの準備
	log := logwriter.LogWriter.GetInstance()

	erro := log.Open("./output.log")
	if erro != nil {
		panic(erro)
	}

	// log.WriteS("1st line\n")
	// log.Outf.WriteString(fmt.Sprintf("%d,%f\n", 101, 1.23))	// こちらの方がlogインスタンス内に複数のファイル操作ができるようになる（少し汚いが）

	// シナリオ読み込み
	simframe.Initialize()

	// objDataDBのテスト
	// obj, _ := model.GetObjData(102)
	// fmt.Printf(" 102 vel= %f\n", obj.Vel)

	// errの使い方の確認　（テスト用）
	// obj, err := model.GetObjData(101)
	// if err != nil {
	// 	fmt.Printf(" GetObjData: Error (%s)\n", err)
	// }
	// fmt.Printf("101 : %v\n", obj)

	// シミュレーション本体
	simframe.RunSim()

	// log用ファイルのクローズ
	log.Close()

}
