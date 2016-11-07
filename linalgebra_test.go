package mlscratchlib

import (
	"reflect"
	"sort"
	"testing"
)

// sample vectors for testing
var vec8a = []float64{99, 42, 75, 11, 13, 100, 97, 66}
var vec8b = []float64{1, 2, 3, 4, 5, 6, 7, 8}
var vec8c = []float64{2.2, 93.1, 8, 0.01, 3, 55, 77.001, 1000.00}
var vec10a = []float64{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}

// sample matrixes for testing
var matrix3a = [][]float64{vec8a, vec8b, vec8c}

func TestSumValues(t *testing.T) {
	var expected, result float64
	result = SumValues(vec8b)
	expected = 36

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}

func TestVectorMean(t *testing.T) {
	var expected, result float64
	result = VectorMean(vec8a)
	expected = 62.875

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}

func TestRangeVector(t *testing.T) {
	var expected, result float64
	result = RangeVector(vec8b)
	expected = 7

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}

func TestVectorMedian(t *testing.T) {
	var expected, result float64
	result = VectorMedian([]float64{})
	expected = 0

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = VectorMedian(vec8b) // {1, 2, 3, 4, 5, 6, 7, 8}
	expected = 4.5               // average of 4 & 5 the middle two elements

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = VectorMedian([]float64{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 101})
	expected = 50

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}

func TestAddVector(t *testing.T) {
	var err error
	var expected []float64

	// test expected logic
	result, _ := AddVector(vec8a, vec8b)
	expected = []float64{100, 44, 78, 15, 18, 106, 104, 74}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of %v, got result: %v", expected, result)
	}

	result, err = AddVector(vec10a, vec8a)
	expected = nil

	// test the length mismatch logic
	if result != nil {
		t.Errorf("vector length does not match.\nExpected nil, got %v.\nAddVector err value: %v", result, err)
	}
}

func TestSubtractVector(t *testing.T) {
	var err error
	var expected []float64

	// test expected logic
	result, _ := SubtractVector(vec8a, vec8b)
	expected = []float64{98, 40, 72, 7, 8, 94, 90, 58}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of %v, got result: %v", expected, result)
	}

	// test the length mismatch logic
	result, err = SubtractVector(vec8b, vec10a)

	if result != nil {
		t.Errorf("vector length does not match.\nExpected nil, got %v.\nSumbtractVector err value: %v", result, err)
	}
}

func TestSumVectors(t *testing.T) {
	var err error
	var expected []float64
	var result []float64

	// test expected logic
	vectors := [][]float64{vec8a, vec8b, vec8c}
	result, err = SumVectors(vectors)
	expected = []float64{102.2, 137.1, 86, 15.01, 21, 161, 181.001, 1074}
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestScalarMultiply(t *testing.T) {
	var expected []float64
	var result []float64

	result = ScalarMultiply(5.0, vec8a)
	expected = []float64{495, 210, 375, 55, 65, 500, 485, 330}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestMeanVector(t *testing.T) {
	var err error
	var expected, result []float64

	vectors := [][]float64{vec8a, vec8b, vec8c}
	result, err = MeanVector(vectors)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// this is what is returned from summing the vectors[]float64{102.2, 137.1, 86, 15.01, 21, 161, 181.001, 1074}
	// multiply each element by 0.3333333333 which is 1/len(vectors) since vectors is a slice of 3 vectors
	expected = []float64{34.06666666666666, 45.699999999999996, 28.666666666666664, 5.003333333333333, 7, 53.666666666666664, 60.333666666666666, 358}

	if !reflect.DeepEqual(result, expected) {
		res, _ := SumVectors(vectors)
		t.Errorf("Expected result of:\n%v\ngot result:\n%v\n\nresult of summing vectors:\n%v", expected, result, res)
	}
}

func TestQuantileVector(t *testing.T) {
	var err error
	var expected, result float64

	result, err = QuantileVector(vec8b, .50)
	expected = 5

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestModeVector(t *testing.T) {
	var err error
	var expected, result []float64

	result, err = ModeVector(vec8b)
	expected = nil // because they all have the same number of occurrences

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = ModeVector([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 9, 5, 5, 5, 5, 5, 5, 5, 5, 9, 9, 7, 1, 2, 1})
	expected = []float64{5} // because 5 occurs the most times in the list

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = ModeVector([]float64{1, 1, 1, 2, 4, 5, 5, 5, 6, 7, 7, 7, 8, 9, 9, 0, 0})
	sort.Float64s(result)         // elements are appended to a slice from a map, so can come out in dif orders every time
	expected = []float64{1, 5, 7} // because each of these values occur three times

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestDotProduct(t *testing.T) {
	var err error
	var expected, result float64

	result, err = DotProduct(vec8a, vec8b)
	expected = 2324

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestSumofSquares(t *testing.T) {
	var err error
	var expected, result float64

	result, err = SumofSquares(vec8b)
	expected = 204

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestMagnitude(t *testing.T) {
	var err error
	var expected, result float64

	result, err = Magnitude(vec8c)
	expected = 1008.8109853193511

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestSquaredDistance(t *testing.T) {
	var err error
	var expected, result float64

	result, err = SquaredDistance(vec8a, vec8b)
	expected = 36801

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestDistance(t *testing.T) {
	var err error
	var expected, result float64

	result, err = Distance(vec8a, vec8b)
	expected = 191.835867344978

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestDeMeanVector(t *testing.T) {
	var expected, result []float64

	result = DeMeanVector(vec8a)
	expected = []float64{36.125, -20.875, 12.125, -51.875, -49.875, 37.125, 34.125, 3.125}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestVarianceVector(t *testing.T) {
	var expected, result float64

	result = VarianceVector(vec8a) // 99, 42, 75, 11, 13, 100, 97, 66
	expected = 1374.125

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result = VarianceVector([]float64{11111111.111111111})
	expected = 0

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestCovariance(t *testing.T) {
	var err error
	var expected, result float64

	result, err = Covariance(vec8b, vec8c)
	expected = 503.43535714285724

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = Covariance([]float64{}, vec8a) // test for nil slice
	expected = 0

	if err == nil {
		t.Errorf("Function accepted a nil slice BAD BAD")
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = Covariance(vec8a, vec10a) // test for slices with different lengths
	expected = 0

	if err == nil {
		t.Errorf("Slices must match, accepted slices of mismatched lenghths")
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestCorrelation(t *testing.T) {
	var err error
	var expected, result float64

	result, err = Correlation(vec8c, vec8b)
	expected = 0.5983028629867763

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = Correlation([]float64{}, vec8a) // test for nil slice
	expected = 0

	if err == nil {
		t.Errorf("Function accepted a nil slice BAD BAD")
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = Covariance(vec8a, vec10a) // test for slices with different lengths
	expected = 0

	if err == nil {
		t.Errorf("Slices must match, accepted slices of mismatched lenghths")
	}

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = Correlation(vec8a, vec8a) // test that matching vectors returns perfect pos correlation
	expected = 1

	if result != expected {
		t.Logf("Matching vectors should have perfect positive correlation ie. 1")
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = Correlation([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []float64{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10})
	expected = -1

	// this test returns 0.999repeating
	if int(result) != int(expected) && expected > -0.9999999999 && expected < -0.9999999998 {
		t.Logf("Opposite vectors should have perfect negative correlation ie. -1")
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestStandardDeviationVector(t *testing.T) {
	var expected, result float64

	result = StandardDeviationVector(vec8a)
	expected = 37.069192060254025

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestInterQuartileRangeVector(t *testing.T) {
	var expected, result float64

	result = InterQuartileRangeVector(vec8a)
	expected = 57

	if result != expected {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

}

func TestShape(t *testing.T) {
	var columns, rows, eColumns, eRows int

	rows, columns = Shape(matrix3a)
	eRows, eColumns = 3, 8

	if columns != eColumns || rows != eRows {
		t.Errorf("Expected Columns: %d, Rows: %d\nGot Columns: %d, Rows: %d", eColumns, eRows, columns, rows)
	}
}

func TestGetRow(t *testing.T) {
	var err error
	var expected, result []float64

	result, err = GetRow(matrix3a, 0)
	expected = vec8a

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = GetRow(matrix3a, len(matrix3a)+1)
	expected = nil

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestGetColumn(t *testing.T) {
	var err error
	var expected, result []float64

	t.Logf("%v", matrix3a)

	matrix := [][]float64{{99, 42, 75, 11, 13, 100, 97, 66}, {1, 2, 3, 4, 5, 6, 7, 8}, {2.2, 93.1, 8, 0.01, 3, 55, 77.001, 1000.00}}

	result, err = GetColumn(matrix, 1)
	expected = []float64{42, 2, 93.1}

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result, err = GetColumn(matrix3a, len(matrix3a)+1)
	expected = nil

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestCreateMatrix(t *testing.T) {
	var expected, result [][]float64

	result = CreateMatrix(5, 3, IsDiagonal)
	expected = [][]float64{{1, 0, 0, 0, 0}, {0, 1, 0, 0, 0}, {0, 0, 1, 0, 0}}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}

func TestIsDiagonal(t *testing.T) {
	var expected, result float64

	result = IsDiagonal(0, 1)
	expected = 0

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}

	result = IsDiagonal(0, 0)
	expected = 1

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result of:\n%v\ngot result:\n%v", expected, result)
	}
}
