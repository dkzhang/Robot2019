package server

import "fmt"

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
