package main

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

func TestMakeNormalizedParaMatrix(t *testing.T) {

	para := []float64{1.0, 2.0, 3.0, 4.0}

	matrix := MakeNormalizedParaMatrix(para)

	for _, r := range matrix {
		t.Logf("%v", r)
	}
}

func TestHorizontalTwoWayBlooming(t *testing.T) {
	para := []float64{1.0, 2.0, 3.0, 4.0}
	gain := []float64{1.0, 1.0, 1.0, 1.0}

	theArray := make([]float64, 16)
	for i := 0; i < len(theArray); i++ {
		if i < 8 {
			theArray[i] = float64(i)
		} else {
			theArray[i] = float64(i - 8)
		}
	}

	t.Logf("%v", theArray)
	newArray := HorizontalTwoWayBlooming(theArray, para, gain, 16, 1, 8)
	t.Logf("%v", newArray)
}
