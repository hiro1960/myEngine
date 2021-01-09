// Point型　coreパッケージ
package core

// Point型
type Point struct {
	Value [3]float64
}

// Pointのコンストラクタ
func NewPoint(v [3]float64) *Point {
	// p := new(Point)
	var p *Point = new(Point) // こちらの方が型が明確に分かるので、いい気がする
	for i, _ := range v {
		p.Value[i] = v[i]
	}
	return p
}

// Pointの加算（自分自身に加算する）
// (p自体もポインタにしておかないと変更結果を外に繁栄できない)
func (p *Point) Add(v *Point) {
	for i, _ := range p.Value {
		p.Value[i] = p.Value[i] + v.Value[i]
	}
}

// Pointの減算（自分自身に減算する）
func (p *Point) Sub(v *Point) {
	for i, _ := range p.Value {
		p.Value[i] = p.Value[i] - v.Value[i]
	}
}
