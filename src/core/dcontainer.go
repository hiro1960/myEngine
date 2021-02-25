// 2変数　線形補間型
package core

import (
	_ "encoding/csv"
	_ "fmt"
	"strconv"
)

// Dcontainer型
type Dcontainer struct {
	index []float64
	value []float64
}

// Dcontiner型　コンストラクタ　（[][]stringを設定値としてもらう）
func NewDcontainer(record [][]string) *Dcontainer {
	p := new(Dcontainer)
	for _,v := range record {
		for i, s := range v {
			if i ==0 {
				// index
				d, _ := strconv.ParseFloat(s, 64)
				p.index = append(p.index, d)
			}
			if i == 1 {
				// value
				d, _ := strconv.ParseFloat(s, 64)
				p.value = append(p.value, d)
			}
		}
	}

	// 値の確認用
	// fmt.Printf("in NewDcontainer:\n")
	// for i, _ := range p.index {
	// 	fmt.Printf("%d %f %f\n", i, p.index[i], p.value[i])
	// }
	// fmt.Printf("lenght = %d\n", len(p.index))

	return p
}

// 補間結果を返す
func (p *Dcontainer)GetValue(v float64) float64 {
	// 下限のチェック
	if v <= p.index[0] {
		// 補間テーブルの下限を超えている
		return p.value[0]
	}
	// 上限のチェック
	if v >= p.index[len(p.index) - 1] {
		// 補間テーブルの上限を超えている
		return p.value[len(p.index) - 1]
	}

	// indexのサーチ
	var l int
	for i, _ := range p.index {
		if v < p.index[i] {
			// 補間対象が見つかった
			l = i	// 変数iは、for-loop内しか使えないので別の変数に代入する
			break;
		}
	}

	r := Hokan(v, p.index[l - 1], p.index[l], p.value[l - 1], p.value[l])

	return r
}