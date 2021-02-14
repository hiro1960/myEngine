// shipの実装本体
// (ディレクトリ名とパッケージ名は同じ必要があるが、ファイル名は関係ない)
package shipModel

import (
	_ "fmt"
	. "core"
)

type vel struct {
	Vel float64
}

// shipとのデータ交換用変数(現状、未使用)
var VV vel

func UpdateVel(vel *Point) {
	// fmt.Printf("in updateVel(): Vel = %f\n", vel)

	for i, _ := range vel.Value {
		vel.Value[i] = vel.Value[i] - 0.1
	}
}