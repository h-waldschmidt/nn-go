package neuralnet

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func Mse(yTrue, yPred mat.VecDense) float64 {
	if yTrue.Len() != yPred.Len() {
		panic("vectors need to have the same dimesions")
	}

	var sum float64
	for i := 0; i < yTrue.Len(); i++ {
		sum += math.Pow(yTrue.AtVec(i)-yPred.AtVec(i), 2)
	}

	return sum / float64(yTrue.Len())
}

func MseDerivative(yTrue, yPred mat.VecDense) mat.VecDense {
	if yTrue.Len() != yPred.Len() {
		panic("vectors need to have the same dimesions")
	}

	yTrue.SubVec(&yPred, &yTrue)
	yTrue.ScaleVec(2, &yTrue)
	return yTrue
}

func Mae(yTrue, yPred mat.VecDense) float64 {
	if yTrue.Len() != yPred.Len() {
		panic("vectors need to have the same dimesions")
	}

	var sum float64
	for i := 0; i < yTrue.Len(); i++ {
		sum += math.Abs(yTrue.AtVec(i) - yPred.AtVec(i))
	}

	return sum / float64(yTrue.Len())
}

// TODO: add derivative of Mae
