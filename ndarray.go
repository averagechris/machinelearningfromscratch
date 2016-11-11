package mlscratchlib

import (
	"errors"
	"math"
	"sort"
)

// Vector is a
type Vector []float64

// Matrix is a
type Matrix []Vector

// NDmatrix is a
type NDmatrix struct {
	M []Matrix
}

// ===================
// Functions to create
// ===================

// MakeNDmatrix does this
func MakeNDmatrix(dimensions, mcols, mrows int, mstructure func(int, int) float64) (m NDmatrix) {
	for i := 0; i < dimensions; i++ {
		m.M = append(m.M, MakeMatrix(mcols, mrows, mstructure))
	}
	return m
}

// MakeMatrix does this
func MakeMatrix(cols int, rows int, structure func(int, int) float64) (m Matrix) {
	for i := 0; i < rows; i++ {
		v := Vector{}
		for j := 0; j < cols; j++ {
			v = append(v, structure(i, j))
		}
		m = append(m, v)
	}
	return m
}

// ==========================
// Matrix Structure Functions
// ==========================

// Zeroed is a structure function for creating zeroed matrixes with
func Zeroed(x, y int) float64 {
	return 0
}

// Diagonal is a structure function for creating matrixes that accepts
// two integers and returns a single float - 1 if the two integers
// match or a 0 if the integers do not match.
func Diagonal(x, y int) float64 {
	if x == y {
		return 1
	}
	return 0
}

// =========
// Functions
// =========

// AddVectors accepts two vectors a, b and returns a new vector whose elements
// are the sum of each element of vector a and b
func AddVectors(a, b Vector) (vector Vector, err error) {
	if len(a) != len(b) {
		return nil, errors.New("Vectors must have the same number of elements")
	}
	for i := range a {
		vector = append(vector, a[i]+b[i])
	}
	return vector, nil
}

// SubtractVectors accepts two vectors a, b and returns a new vector whose elements
// are the difference between each element of vector a and b
func SubtractVectors(a, b Vector) (vector Vector, err error) {
	if len(a) != len(b) {
		return nil, errors.New("Vectors must have the same number of elements")
	}
	for i := range a {
		vector = append(vector, a[i]-b[i])
	}
	return vector, nil
}

// ScalarMultiplication returns a new vector whose elements are the product
// of the given scalar value and each element of the given vector
func ScalarMultiplication(scalar float64, v Vector) (vector Vector) {
	for _, element := range v {
		vector = append(vector, element*float64(scalar))
	}
	return vector
}

// DotProduct returns the sum of the componentwise product of the given vectors a and b
func DotProduct(a, b Vector) (product float64, err error) {
	if len(a) < 1 {
		return 0, errors.New("vectors cannot be nil")
	} else if len(a) != len(b) {
		return 0, errors.New("vectors must be the same length")
	}
	product = 0
	for i := range a {
		product += a[i] * b[i]
	}
	return product, nil
}

// SquaredDistance returns the sum of squares of a vector that is the
// result of subtracting the given vector a from the given vector b
func SquaredDistance(a, b Vector) (float64, error) {
	subtracted, err := SubtractVectors(a, b)
	if err != nil {
		return 0, err
	}

	result, err := subtracted.SumOfSquares()
	if err != nil {
		return 0, err
	}

	return result, nil
}

// Distance returns a float that represents the distance between two vectors
// by returning the magnitude of the difference between the given vectors a and b
func Distance(a, b Vector) (float64, error) {
	vector, err := SubtractVectors(a, b)
	if err != nil {
		return 0, err
	}

	distance, err := vector.Magnitude()
	if err != nil {
		return 0, err
	}

	return distance, nil
}

// DeMeanVector returns a new vector whose elements are the difference between
// the mean element value and the element value of the given vector v
func DeMeanVector(v Vector) (vector Vector) {
	mean := v.Mean()
	for _, element := range v {
		vector = append(vector, element-mean)
	}
	return vector
}

// Covariance returns a float64 that represents how the values of the elements of the
// given vectors a and b change in tandem
// https://en.wikipedia.org/wiki/Covariance
func Covariance(a, b Vector) (float64, error) {
	deviations, err := DotProduct(DeMeanVector(a), DeMeanVector(b))
	if err != nil {
		return 0, err
	}
	return (deviations / float64(len(a)-1)), nil
}

// Correlation returns a float64 that represents the Covariance divided by the
// standard deviation of both of the given vectors a and b
// Correlation will always return a float between -1 and 1
// 0 = no correlation ie no linear relationship
// 1 = perfect positive correlation
// -1 = perfect negative correlation
// https://en.wikipedia.org/wiki/Correlation_and_dependence
func Correlation(a, b Vector) (float64, error) {

	aStdDev, err := a.StandardDeviation()
	if err != nil {
		return 0, err
	}

	bStdDev, err := b.StandardDeviation()
	if err != nil {
		return 0, err
	}

	if aStdDev > 0 && bStdDev > 0 {
		covar, err := Covariance(a, b)
		if err != nil {
			return 0, err
		}
		return (covar / aStdDev / bStdDev), nil
	}

	// one of the vectors has a std deviation of 0
	return 0, nil
}

// ==============
// Vector Methods
// ==============

// Sort arranges the elements of the vector in ascending order
func (v Vector) Sort() {
	sort.Float64s(v)
}

// Rsort arranges the elements of the vector in descending order
func (v Vector) Rsort() {
	sort.Sort(sort.Reverse(sort.Float64Slice(v)))
}

// SumValues returns the sum of each element of a vector
func (v Vector) SumValues() float64 {
	var sum float64
	for _, element := range v {
		sum += element
	}
	return sum
}

// Mean returns the mean value of the elements of a vector
func (v Vector) Mean() float64 {
	if len(v) < 1 {
		return 0
	}
	return v.SumValues() / float64(len(v))
}

// DeMean subtracts the mean value of the vector from each element of the vector
func (v Vector) DeMean() {
	mean := v.Mean()
	for i, element := range v {
		v[i] = element - mean
	}
}

// Median returns the median value of the elements of a vector
func (v Vector) Median() float64 {
	length := len(v)
	if length < 1 {
		return 0
	}
	sort.Float64s(v)

	if length%2 != 0 {
		// if there are an odd number of elements, return the one in the middle index
		return v[length/2]
	}

	// else return the average of the two middle-most elements
	high := length / 2
	low := high - 1
	return (v[high] + v[low]) / 2
}

// Quantile accepts a float64 and returns the element that represents the percent
// value of the range of elements in the vector.
// ex. v = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}; v.Quantile(0.50) returns 6
func (v Vector) Quantile(percentile float64) (quantile float64, err error) {

	if len(v) < 1 {
		return 0, errors.New("Vector cannot be nil; len > 1")
	}

	if percentile < 0 || percentile > 1 {
		return 0, errors.New("Arg must be a decimal between 0 and 1")
	}

	var tmp Vector
	copy(tmp, v)
	tmp.Sort()
	index := int(percentile * float64(len(v)))

	return tmp[index], nil
}

// Mode returns a new Vector that contains the most commone elements
func (v Vector) Mode() (mode Vector, err error) {

	if len(v) < 1 {
		return nil, errors.New("Vector cannot be nil; len < 1")
	}

	// make a map
	// keys = unique elements
	// values = number of times an element occures in the vector
	occurences := make(map[float64]int)
	for _, element := range v {
		if occurences[element] == 0 {
			occurences[element] = 1 // add the element to occurrences, initialize to 1
			continue
		}
		occurences[element]++ // increment the occurrence of the element
	}

	// iterate over the map to find the elements that occured the most
	var mostOccurring int
	for element, occurrence := range occurences {
		if occurrence == mostOccurring {
			mode = append(mode, element)
		} else if occurrence > mostOccurring {
			mode = []float64{} // reset mode to remove previous maxes
			mode = append(mode, element)
			mostOccurring = occurrence
		}
	}

	if mostOccurring <= 1 {
		// no element occured more than once in the vector
		return nil, nil
	}
	return mode, nil
}

// Range returns the difference between the highest and the lowest values
// of the elements in the vector
func (v Vector) Range() float64 {
	length := len(v)
	if length < 1 {
		return 0
	}
	tmp := Vector(make([]float64, length))
	copy(tmp, v)
	tmp.Sort()
	return tmp[length-1] - tmp[0]
}

// InterQuartileRange returns a float64 that represents the innerquartile range
// https://en.wikipedia.org/wiki/Interquartile_range
func (v Vector) InterQuartileRange() (float64, error) {
	upper, err := v.Quantile(0.75)
	if err != nil {
		return 0, err
	}
	lower, _ := v.Quantile(0.25)
	if err != nil {
		return 0, err
	}

	return upper - lower, nil
}

// ScalarMultiply multiplies each element of the vector by the given scalar value
// It's the inplace version of the function ScalarMultiplication
func (v Vector) ScalarMultiply(scalar float64) {
	for i, e := range v {
		v[i] = float64(e) * float64(scalar)
	}
}

// SumOfSquares returns the sum of the squared elements
// https://en.wikipedia.org/wiki/Sum_of_squares
func (v Vector) SumOfSquares() (sumofsquares float64, err error) {
	sumofsquares, err = DotProduct(v, v)
	if err != nil {
		return 0, err
	}
	return sumofsquares, nil
}

// Magnitude returns the magnitude of the vector
// https://en.wikipedia.org/wiki/Magnitude_(mathematics)
func (v Vector) Magnitude() (float64, error) {
	s, err := v.SumOfSquares()
	if err != nil {
		return 0, err
	}
	return math.Sqrt(s), nil
}

// Variance returns a float64 that represents how much the value of the elements deviate
// from the mean value of the elements in the vector
func (v Vector) Variance() (float64, error) {
	length := len(v)
	if length < 2 {
		return 0, nil
	}
	deviations := DeMeanVector(v)
	sumos, err := deviations.SumOfSquares()
	if err != nil {
		return 0, err
	}
	return sumos / float64(length-1), nil
}

// StandardDeviation returns the squareroot of the variance of the element values
// https://en.wikipedia.org/wiki/Standard_deviation
func (v Vector) StandardDeviation() (float64, error) {
	variance, err := v.Variance()
	if err != nil {
		return 0, err
	}
	return math.Sqrt(variance), nil
}

// ==============
// Matrix Methods
// ==============

// SumVectors returns a single vector that is the sum of each element
// of the vectors in the matrix componentwise
func (m Matrix) SumVectors() (vector Vector, err error) {
	if len(m) < 1 {
		return nil, errors.New("The matrix cannot be nil")
	}
	// add the elements of each vector to the first vector in the matrix
	vector = m[0]
	for _, v := range m[1:] {
		vector, err = AddVectors(vector, v)
		if err != nil {
			return nil, err
		}
	}
	return vector, nil
}

// MeanVector returns a vector whose elements are the componentwise mean value
// of all of the vectors in the matrix
func (m Matrix) MeanVector() (vector Vector, err error) {
	length := len(m)
	if length < 1 {
		return nil, errors.New("The matrix cannot be nil")
	}
	vector, err = m.SumVectors()
	if err != nil {
		return nil, err
	}
	vector.ScalarMultiply(float64(1 / length))
	return vector, nil
}

// Shape returns two integers that represent the number of
// rows and columns in a matrix
func (m Matrix) Shape() (rows int, columns int) {
	length := len(m)
	if length < 1 {
		return 0, 0
	}
	return length, len(m[0])
}

// GetRow safely returns the vector by index or nil of a vector does not exist at that index
func (m Matrix) GetRow(rowNumber int) Vector {
	if rowNumber > len(m) {
		return nil
	}
	return m[rowNumber]
}

// GetColumn safely returns a vector the elements of which are the value of the nth index
// of each row according to the given columnNumber
func (m Matrix) GetColumn(columnNumber int) (column Vector) {
	if columnNumber > len(m) {
		return nil
	}
	for _, row := range m {
		column = append(column, row[columnNumber])
	}

	return column
}
