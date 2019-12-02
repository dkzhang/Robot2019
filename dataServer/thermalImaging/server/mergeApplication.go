package main

import "fmt"

func MoveMergeThermalArray(a1, a2, a3, a4 []float64, offset [4]int, w, h int) (r []float64, newWidth, newHeight int, err error) {

	//for example:
	//offset := [4]int{0,6,10,16}

	if len(a1) != w*h || len(a1) != w*h || len(a1) != w*h || len(a1) != w*h {
		return nil, -1, -1,
			fmt.Errorf("illegal array length: %d,%d,%d,%d <=> %d, %d", len(a1), len(a2), len(a3), len(a4), w, h)
	}

	fmt.Printf("a1,a2,a3,a4: %d,%d,%d,%d <=> %d, %d \n", len(a1), len(a2), len(a3), len(a4), w, h)

	//默认最后一个矩阵的偏移量最大，即offset[3]最大
	newWidth = offset[3] + w
	newHeight = h

	r = make([]float64, newWidth*h)
	r1 := make([]float64, newWidth*h)
	r2 := make([]float64, newWidth*h)
	r3 := make([]float64, newWidth*h)
	r4 := make([]float64, newWidth*h)

	factor := make([]float64, newWidth*h)

	for i := 0; i < h; i++ {
		iw := i * w
		iwn := i * newWidth
		copy(r1[iwn+offset[0]:iwn+offset[0]+w], a1[iw:iw+w])
		copy(r2[iwn+offset[1]:iwn+offset[1]+w], a1[iw:iw+w])
		copy(r3[iwn+offset[2]:iwn+offset[2]+w], a1[iw:iw+w])
		copy(r4[iwn+offset[3]:iwn+offset[3]+w], a1[iw:iw+w])

		for m := 0; m < 4; m++ {
			for n := 0; n < w; n++ {
				factor[iwn+offset[m]+n] += 1.0
			}
		}
	}

	for i := 0; i < len(r); i++ {
		r[i] = (r1[i] + r2[i] + r3[i] + r4[i]) / factor[i]
	}
	return
}

func MergeThermalArray(a1, a2, a3, a4 []float64, w, h int) (r []float64, newWidth, newHeight int, err error) {

	if len(a1) != w*h || len(a1) != w*h || len(a1) != w*h || len(a1) != w*h {
		return nil, -1, -1,
			fmt.Errorf("illegal array length: %d,%d,%d,%d <=> %d, %d", len(a1), len(a2), len(a3), len(a4), w, h)
	}

	fmt.Printf("a1,a2,a3,a4: %d,%d,%d,%d <=> %d, %d \n", len(a1), len(a2), len(a3), len(a4), w, h)

	r = make([]float64, w*h*4)
	//横向合并
	newWidth = w * 4
	newHeight = h

	for i := 0; i < h; i++ {
		iw := i * w
		iwn := i * newWidth
		copy(r[iwn+0*w:iwn+1*w], a1[iw:iw+w])
		copy(r[iwn+1*w:iwn+2*w], a2[iw:iw+w])
		copy(r[iwn+2*w:iwn+3*w], a3[iw:iw+w])
		copy(r[iwn+3*w:iwn+4*w], a4[iw:iw+w])
	}

	return
}

func HorizontalTwoWayBlooming(theArray []float64, para, gain []float64, width, height int, bWidth int) []float64 {
	//Parameter check

	//copy
	newArray := make([]float64, len(theArray))
	copy(newArray, theArray)

	//
	bTimes := width/bWidth - 1
	bLen := len(para)
	normalizedParaMatrix := MakeNormalizedParaMatrix(para)

	for row := 0; row < height; row++ {
		for bt := 1; bt <= bTimes; bt++ {

			//left
			for i := bWidth*bt - bLen; i < bWidth*bt; i++ {
				result := normalizedParaMatrix[i-(bWidth*bt-bLen)][bLen] * theArray[i]

				for j := bWidth * bt; j < bWidth*bt+bLen; j++ {
					result += normalizedParaMatrix[i-(bWidth*bt-bLen)][j-(bWidth*bt)] * theArray[j]
				}

				newArray[i] = result * gain[i-(bWidth*bt-bLen)]
			}

			//right
			for i := bWidth*bt + bLen - 1; i >= bWidth*bt; i-- {
				result := normalizedParaMatrix[(bWidth*bt+bLen-1)-i][bLen] * theArray[i]

				for j := bWidth*bt - 1; j >= bWidth*bt-bLen; j-- {
					result += normalizedParaMatrix[(bWidth*bt+bLen-1)-i][(bWidth*bt-1)-j] * theArray[j]
				}
				newArray[i] = result * gain[(bWidth*bt+bLen-1)-i]
			}
		}
	}
	return newArray
}

func MakeNormalizedParaMatrix(para []float64) [][]float64 {
	bLen := len(para)

	normalizedParaMatrix := make([][]float64, bLen)

	for i := 0; i < bLen; i++ {
		array := make([]float64, bLen+1)

		sum := 0.0
		for j := 0; j <= i; j++ {
			array[j] = para[i-j]
			sum += array[j]
		}
		sum += 1

		for j := 0; j <= i; j++ {
			array[j] /= sum
		}
		array[bLen] = 1.0 / sum

		normalizedParaMatrix[i] = array
	}

	return normalizedParaMatrix
}
