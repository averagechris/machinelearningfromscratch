package mlscratchlib

import (
	"errors"
	"math"
)

// AddVector returns a new vector whose elements are the sum of each
// element of the given vectors a and b
func AddVector(a []float64, b []float64) (vector []float64, err error) {
	if len(a) != len(b) {
		return nil, err
	}
	for i := range a {
		vector = append(vector, a[i]+b[i])
	}
	return vector, nil
}

// SubtractVector returns a new vector whose elements are the difference
// between vector a and vector b
func SubtractVector(a []float64, b []float64) (vector []float64, err error) {
	if len(a) != len(b) {
		return nil, err
	}
	for i := range a {
		vector = append(vector, a[i]-b[i])
	}
	return vector, nil
}

// SumVectors takes a slice of vectors and returns a new vector
// the elements of which are the componentwise sum of the slice of
// vectors. example: vec1 = {1, 2, 3} vec2 = {4, 5, 6} vec3 = {7, 8, 9}
// return value = vector{12, 15, 18}
func SumVectors(vectors []([]float64)) (vector []float64, err error) {
	vector = vectors[0]
	for _, v := range vectors[1:] {
		vector, err = AddVector(vector, v)
		if err != nil {
			return nil, err
		}
	}
	return vector, nil
}

// ScalarMultiply takes an int and a vector and returns a new
// vector whose elements are the product of the int and each
// element of the given vector
func ScalarMultiply(num float64, v []float64) (vector []float64) {
	for _, element := range v {
		vector = append(vector, element*float64(num))
	}
	return vector
}

// MeanVector accepts a slice of vectors, sums each element of each
// vector and returns the mean vector
func MeanVector(vectors []([]float64)) (vector []float64, err error) {
	sumvec, err := SumVectors(vectors)
	if err != nil {
		return nil, err
	}
	return ScalarMultiply(1/float64(len(vectors)), sumvec), nil
}

// DotProduct accepts two vectors and returns a number which is the
// the sum of the product of each element in a and b
func DotProduct(a []float64, b []float64) (product float64, err error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must be the same length")
	}
	product = 0
	for i := range a {
		product += a[i] * b[i]
	}
	return product, nil
}

// SumofSquares accepts a vector, squares each element
// and returns the sum of the squared elements
func SumofSquares(vector []float64) (sum float64, err error) {
	sum, err = DotProduct(vector, vector)
	if err != nil {
		return 0, err
	}
	return sum, nil
}

// Magnitude accepts a vector, squares each element, returns the
// square root of the sum of the squared elements which is used
// to measure the length of a vector
func Magnitude(v []float64) (float64, error) {
	sumos, err := SumofSquares(v)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(sumos), nil
}

// SquaredDistance accepts two vectors and returns a float64
// that is the result of subtracting each element of the given
// vectors componentwise, squaring each resulting element
// and summing the resulting vector
// example: (a1 - b1)**2 + (a2 - b2)**2 + ... (aN - bN)**2
func SquaredDistance(a []float64, b []float64) (result float64, err error) {

	if len(a) != len(b) {
		return 0, errors.New("vectors must be the same length")
	}

	var subtracted []float64
	subtracted, err = SubtractVector(a, b)

	if err != nil {
		return 0, err
	}

	result, err = SumofSquares(subtracted)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// Distance accepts two vectors and returns a float64
// that is the result of subtracting each element of the given
// vectors componentwise, and taking the magnitude of the
// resulting vector
func Distance(a []float64, b []float64) (float64, error) {
	vector, err := SubtractVector(a, b)

	if err != nil {
		return 0, err
	}

	result, err := Magnitude(vector)
	if err != nil {
		return 0, err
	}

	return result, nil
}