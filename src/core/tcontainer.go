// 3変数　線形補間型
package core

import (
	"fmt"
	"strconv"
)

// Tcontainer型
type Tcontainer struct {
	index1 []float64
	index2 []float64
	numOfTag1 int32
	numOfTag2 int32
	value  []float64
}

// Tcontainer型　コンストラクタ　（[][]stringを設定値としてもらう）
func NewTcontainer(record [][]string) *Tcontainer {
	p := new(Tcontainer)
	for _, v := range record {
		for i, s := range v {
			if i == 0 {
				// index1
				d, _ := strconv.ParseFloat(s, 64)
				p.index1 = append(p.index1, d)
			}
			if i == 1 {
				// index2
				d, _ := strconv.ParseFloat(s, 64)
				p.index2 = append(p.index2, d)
			}
			if i == 2 {
				// value
				d, _ := strconv.ParseFloat(s, 64)
				p.value = append(p.value, d)
			}
		}

	}

	// 値の確認用
	fmt.Printf("in NewTcontainer:\n")
	fmt.Printf("lenght = %d\n", len(p.index1))
	fmt.Printf("lenght = %d\n", len(p.index2))
	fmt.Printf("lenght = %d\n", len(p.value))
	for i, _ := range p.index1 {
		fmt.Printf("%d %f %f %f\n", i, p.index1[i], p.index2[i], p.value[i])
		// fmt.Printf("%d\n", i)
	}

	// index2の数を数える
	var cnt int
	cnt = 0
	p.numOfTag2 = 0
	t1 := p.index1[0]
	for {
		if t1 != p.index1[cnt] {
			break
		}
		cnt++
		p.numOfTag2++
	}
	fmt.Printf("numOftag2 = %d\n", p.numOfTag2)

	// index1の数を数える
	cnt = 0
	p.numOfTag1 = 1
	t2 := p.index1[0]
	for {
		if t2 != p.index1[cnt] {
			p.numOfTag1++
		}
		t2 = p.index1[cnt]
		cnt++

		if cnt == len(p.index1) {
			// データの終端に達した
			break
		}
	}
	fmt.Printf("numOftag1 = %d\n", p.numOfTag1)

	return p
}

// 補間結果を返す
func (p *Tcontainer) GetValue(v1, v2 float64) float64 {
	var idx1 float64	// index1の入力
	var idx2 float64	// index2の入力

	// 初期値設定
	idx1 = v1
	idx2 = v2

	// index1 下限のチェック
	if v1 <= p.index1[0] {
		// index1の下限を超えている
		idx1 = p.index1[0]
	}
	// index1 上限のチェック
	if v1 >= p.index1[len(p.index1)-1] {
		idx1 = p.index1[len(p.index1)-1]
	}
	// index2 下限のチェック
	if v2 <= p.index2[0] {
		// index2の下限を超えている
		idx2 = p.index2[0]
	}
	// index2 上限のチェック
	if v2 >= p.index2[len(p.index2)-1] {
		idx2 = p.index2[len(p.index2)-1]
	}

	// 変数を使用したことにする
	fmt.Printf("%f %f\n", idx1, idx2)

	return 0.0
}