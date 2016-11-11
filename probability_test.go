package mlscratchlib

import "testing"

func TestUniformProbabilityDistribution(t *testing.T) {
	var expected, result float64
	result = UniformProbabilityDistribution(.0444)
	expected = 1

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = UniformProbabilityDistribution(2)
	expected = 0

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = UniformProbabilityDistribution(-2) // test the negative value case
	expected = 0

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}

func TestUniformCumulativeDensity(t *testing.T) {
	var expected, result float64

	result = UniformCumulativeDensity(1)
	expected = 1

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = UniformCumulativeDensity(0)
	expected = 0

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = UniformCumulativeDensity(0.5)
	expected = 0.5

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}

func TestNormalProbabilityDistribution(t *testing.T) {
	var expected, result float64

	expected = NormalProbabilityDistribution(1, 0, 1)
	result = 0.24197072451914337

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}

func TestNormalCDF(t *testing.T) {
	var expected, result float64

	result = NormalCDF(-10, 0, 1) // n = -10, mean = 0, sigma = 1
	expected = 0.0

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = NormalCDF(9, 0, 1) // n = 9, mean = 0, sigma = 1
	expected = 1

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = NormalCDF(6, 0, 2) // n = 6, mean = 0, sigma = 2
	expected = 0.9986501019683699

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = NormalCDF(6, -1, 2) // n = 6, mean = 0, sigma = 2
	expected = 0.9997673709209645

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	result = NormalCDF(0, 0, 1)
	expected = 0.5

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}

func TestInverseNormalCDF(t *testing.T) {
	var expected, result float64

	result = InverseNormalCDF(.1, 0, 1, 0.00001)
	expected = -1.2815570831298828

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}

	list := []float64{-9.999990463256836, -9.999990463256836, -9.999990463256836,
		-9.999990463256836, -9.999990463256836, -9.999990463256836, -9.999990463256836,
		-9.999990463256836, -9.999990463256836, -9.999990463256836, -8.75, 8.75,
		9.999990463256836, 9.999990463256836, 9.999990463256836, 9.999990463256836,
		9.999990463256836, 9.999990463256836, 9.999990463256836, 9.999990463256836}

	for i := -10; i < 10; i++ {
		result = InverseNormalCDF(float64(i), 0, 1, 0.00001)
		expected = list[i+10]
		if result != expected {
			t.Errorf("\nExpected: %f\nGot: %f", expected, result)
		}
	}

	result = InverseNormalCDF(9904000002, 2, 2, 0)
	expected = 22

	if result != expected {
		t.Errorf("\nExpected: %f\nGot: %f", expected, result)
	}
}
