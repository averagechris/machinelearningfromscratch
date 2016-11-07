package mlscratchlib

import (
	"errors"
	"math"
	"sort"
)

// SumValues accepts a []float64 vector and returns the sum of each
// of it's elements as a float64
func SumValues(vector []float64) (sum float64) {
	sum = 0
	for _, element := range vector {
		sum += element
	}
	return sum
}

// VectorMean accepts a []float64 vector and returns a float64 number
// that represents the mean value of the elements in the vector
func VectorMean(vector []float64) (meanValue float64) {
	if len(vector) < 1 {
		return 0
	}
	return SumValues(vector) / float64(len(vector))
}

// VectorMedian accepts a []float64 vector and returns the element
// which is in the middle most index when the vector is sorted from
// low to high if there are an odd number of elements and the average
// of the two middle-most elements if the number of elements is even
func VectorMedian(vector []float64) (medianValue float64) {
	if len(vector) < 1 {
		return 0
	}
	sort.Float64s(vector)

	if len(vector)%2 != 0 {
		// if there are an odd number of elements, return the one in the middle index
		middleIndex := (len(vector) / 2)
		return vector[middleIndex]
	}

	// return the average of the two middle-most elements
	high := len(vector) / 2
	low := high - 1
	average := (vector[high] + vector[low]) / 2

	return average
}

// QuantileVector accepts a vector []float64 and a decimal to represent the percentile
// of the element that you want to return.
func QuantileVector(vector []float64, percentile float64) (quantile float64, err error) {
	if len(vector) < 1 {
		return 0, errors.New("something went wrong vector length is 0")
	}
	if percentile < 0 {
		return 0, errors.New("percentile must be a decimal between 0 and 1")
	} else if percentile > 1 {
		return 0, errors.New("percentile must be a decimal between 0 and 1")
	}
	index := int(percentile * float64(len(vector)))
	sort.Float64s(vector)

	return vector[index], nil
}

// ModeVector accepts a vector []float64 and returns a slice of the most common
// of the most common value or values. Can be a slice of 1 element or more.
func ModeVector(vector []float64) (mode []float64, err error) {
	occurences := make(map[float64]int)
	for _, element := range vector {
		if occurences[element] == 0 {
			occurences[element] = 1 // adds the element to the occurrences map and initializes it to 1 occurrence
			continue
		}
		occurences[element]++ // increments the occurrence of the element
	}

	var max int
	for element, occurrence := range occurences {
		if occurrence == max {
			mode = append(mode, element)
		} else if occurrence > max {
			mode = []float64{} // reset mode to remove previous maxes
			mode = append(mode, element)
			max = occurrence
		}
	}

	if max <= 1 {
		return nil, nil // there is no value more than one occurrence
	}
	return mode, nil
}

// RangeVector accepts a vector and returns a float64 that represents the
// difference between the highest and the lowest values in the vector
func RangeVector(vector []float64) float64 {
	if len(vector) < 1 {
		return 0
	}
	sort.Float64s(vector)
	return vector[len(vector)-1] - vector[0]
}

// AddVector returns a new vector whose elements are the sum of each
// element of the given vectors a and b
func AddVector(a []float64, b []float64) (vector []float64, err error) {
	if len(a) != len(b) {
		return nil, errors.New("the vectors must have the same number of elements")
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
		return nil, errors.New("the vectors must have the same number of elements")
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
	if len(vectors) < 1 {
		return nil, errors.New("something went wrong, vectors has 0 elements")
	}
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

// DeMeanVector accepts a vector, computes the mean and subtracts the mean
// from each element of the vector and returns the vector so that the resulting
// vector has a 0 mean
func DeMeanVector(vector []float64) (result []float64) {
	mean := VectorMean(vector)
	for _, element := range vector {
		result = append(result, element-mean)
	}
	return result
}

// VarianceVector accepts a vector, returns a float that represents the
// amount of variance there is in the data within the vector
func VarianceVector(vector []float64) float64 {
	if len(vector) < 2 {
		return 0
	}
	deviations := DeMeanVector(vector)
	sumos, _ := SumofSquares(deviations)
	divisor := len(vector) - 1
	return sumos / float64(divisor)
}

// Covariance accepts two vectors and returns a float that represents the
// amount of variance the two values of the two vectors have in tandem
func Covariance(a []float64, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must be the same length")
	} else if len(a) < 1 {
		return 0, errors.New("something went wrong vector is 0 length")
	}
	deviations, err := DotProduct(DeMeanVector(a), DeMeanVector(b))
	if err != nil {
		return 0, err
	}
	return (deviations / float64(len(a)-1)), nil
}

// StandardDeviationVector accepts a vector, gets the variance and
// returns the square root of the variance
func StandardDeviationVector(vector []float64) float64 {
	return math.Sqrt(VarianceVector(vector))
}

// Correlation accepts two vectors, gets the standard deviation of both
// vector a and b, and returns the covariance of those numbers if they are
// both greater than zero. Correlation returns a float64 that represents a
// good indicator that the elements of the two vectors have an interesting
// relationship. Correlation will always return a float between -1 and 1
// 0 represents no correlation ie no ***LINEAR*** relationship
// 1 represents perfect positive correlation
// -1 represents perfect negative correlation
func Correlation(a []float64, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must be the same length")
	} else if len(a) < 1 {
		return 0, errors.New("something went wrong, vector length is 0")
	}

	aStdDev := StandardDeviationVector(a)
	bStdDev := StandardDeviationVector(b)

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

// InterQuartileRangeVector accepts a vector, calculates the
// quantile for the 25 and 75 percentiles and returns the difference
// as a float.
func InterQuartileRangeVector(vector []float64) float64 {
	upper, _ := QuantileVector(vector, 0.75)
	lower, _ := QuantileVector(vector, 0.25)
	return upper - lower
}

// Shape accepts a matrix of [][]float and returns
// two ints: the number of colums and the number of rows
func Shape(matrix [][]float64) (rows int, columns int) {
	if len(matrix) < 1 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}

// GetRow accepts a matirx of [][]float and an integer
// then returns the row of the matrix coresponding to
// the integer. These matrixes are 0 indexed so GetRow(1)
// will return the 2nd row of the matrix
func GetRow(matrix [][]float64, number int) (row []float64, err error) {
	if number > len(matrix) {
		return nil, errors.New("Index out of range.")
	}
	return matrix[number], nil
}

// GetColumn accepts a matrix of [][]float and a number n
// returns a vector whose elements are the element of the nth
// index of each row
func GetColumn(matrix [][]float64, number int) (column []float64, err error) {
	if number > len(matrix) {
		return nil, errors.New("Index out of range.")
	}
	for _, r := range matrix {
		column = append(column, r[number])
	}

	return column, nil
}

// CreateMatrix accepts two integers, for number of rows and
// number of columns, it also accepts an "entry function" that
// is used to set the value of a cell within the matrix based on
// the position of the cell.
func CreateMatrix(columns int, rows int, entryFunction func(int, int) float64) (matrix [][]float64) {
	for i := 0; i < rows; i++ {
		matrix = append(matrix, make([]float64, columns))
		for index := range matrix[i] {
			matrix[i][index] = entryFunction(i, index)
		}
	}
	return matrix
}

// IsDiagonal is an entry function for CreateMatrix that accepts
// two integers and returns a single float - 1 if the two integers
// match or a 0 if the integers do not match.
func IsDiagonal(x int, y int) float64 {
	if x == y {
		return 1
	}
	return 0
}
