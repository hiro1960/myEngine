// global変数的な物や、汎用函数を定義する
package core

import (
	"math"
)

var RadToDeg float64 = 180.0 / math.Pi
var DegToRad float64 = math.Pi / 180.0

// 線形補間関数
func Hokan(x, x1, x2, v1, v2 float64) float64 {
	var r float64
	if math.Abs(x2-x1) > 0.0 {
		r = v1 + ((v2 - v1) * (x - x1) / (x2 - x1))
	} else {
		r = v1
	}

	return r
}

// deg => rad変換
func ToRadian(v float64) float64 {
	r := v * DegToRad
	return r
}

// rad => deg変換
func ToDegree(v float64) float64 {
	r := v * RadToDeg
	return r
}

// アジマスの差分
func AzimuthDiff(origin, value float64) float64 {
	// [-pi pi]の範囲で角度の差を返す
	// 入力値は　[-pi 2pi]を想定
	if value < 0.0 {
		value += (2.0 * math.Pi)
	}
	diff := value - origin
	if diff > math.Pi {
		diff -= (2.0 * math.Pi)
	} else if diff < -math.Pi {
		diff += (2.0 * math.Pi)
	}

	return diff
}

// アジマスの正規化
func AzimuthAdjust(angle float64) float64 {
	// [-pi pi] -> [0 2pi]への角度の正規化
	var r float64
	if angle < 0.0 {
		r = angle + (2.0 * math.Pi)
	} else {
		r = angle
	}

	return r
}
