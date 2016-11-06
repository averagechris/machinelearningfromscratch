package mlscratchlib

import (
	"reflect"
	"testing"
)

// sample vectors for testing
var vec8a = []float64{99, 42, 75, 11, 13, 100, 97, 66}
var vec8b = []float64{1, 2, 3, 4, 5, 6, 7, 8}
var vec8c = []float64{2.2, 93.1, 8, 0.01, 3, 55, 77.001, 1000.00}
var vec10a = []float64{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}

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
