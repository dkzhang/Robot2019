package server

import "testing"

func TestMergeThermalArray(t *testing.T) {
	var a1, a2, a3, a4 []float64
	w := 4
	h := 4
	a1 = make([]float64, w*h)
	a2 = make([]float64, w*h)
	a3 = make([]float64, w*h)
	a4 = make([]float64, w*h)

	for i := 0; i < w*h; i++ {
		a1[i] = float64(i)
		a2[i] = float64(i)
		a3[i] = float64(i)
		a4[i] = float64(i)
	}
	t.Logf("%v", a1)

	a, nw, nh, err := MergeThermalArray(a1, a2, a3, a4, w, h)
	t.Logf("%v \n nw = %d, nh = %d \n %v", a, nw, nh, err)
}
