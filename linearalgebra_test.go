package mlscratchlib

import (
	"reflect"
	"testing"
)

func TestMakeNDmatrix(t *testing.T) {
	type args struct {
		dimensions int
		mcols      int
		mrows      int
		mstructure func(int, int) float64
	}
	tests := []struct {
		name  string
		args  args
		wantM NDmatrix
	}{
		{
		// name: "standard", args: {3, 10, 10, Zeroed}, wantM: NDmatrix{M{{}, {}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := MakeNDmatrix(tt.args.dimensions, tt.args.mcols, tt.args.mrows, tt.args.mstructure); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("TEST: %v, MakeNDmatrix() = %v, want %v", tt.name, gotM, tt.wantM)
			}
		})
	}
}

func TestMakeMatrix(t *testing.T) {
	type args struct {
		cols      int
		rows      int
		structure func(int, int) float64
	}
	tests := []struct {
		name  string
		args  args
		wantM Matrix
	}{
		{
			name: "small zeroed square", args: args{2, 2, Zeroed}, wantM: Matrix{Vector{0, 0}, Vector{0, 0}},
		}, {
			name: "nil matrix", args: args{0, 0, Zeroed}, wantM: nil,
		}, {
			name: "diagonal structure matrix", args: args{5, 5, Diagonal}, wantM: Matrix{Vector{1, 0, 0, 0, 0}, Vector{0, 1, 0, 0, 0}, Vector{0, 0, 1, 0, 0}, Vector{0, 0, 0, 1, 0}, Vector{0, 0, 0, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := MakeMatrix(tt.args.cols, tt.args.rows, tt.args.structure); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("TEST: %v, MakeMatrix() = %v \n want: %v", tt.name, gotM, tt.wantM)
			}
		})
	}
}

func TestZeroed(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "positive", args: args{99000, 1}, want: 0},
		{name: "negative", args: args{-11, 8}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zeroed(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("TEST: %v, Zeroed() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDiagonal(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "x > y", args: args{999, 1}, want: 0},
		{name: "x < y", args: args{5, 11}, want: 0},
		{name: "x equals y", args: args{5, 5}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diagonal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("TEST: %v, Diagonal() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestAddVectors(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}
	tests := []struct {
		name       string
		args       args
		wantVector Vector
		wantErr    bool
	}{
		{
			name: "same length", args: args{Vector{1, 2, 3, 4, 5}, Vector{6, 7, 8, 9, 10}}, wantVector: Vector{7, 9, 11, 13, 15}, wantErr: false,
		}, {
			name: "different lengths", args: args{Vector{1}, Vector{1, 2, 3, 4, 5}}, wantVector: nil, wantErr: true,
		}, {
			name: "negative values", args: args{Vector{-5, -5, -5, -5}, Vector{99, 97, 95, 93}}, wantVector: Vector{94, 92, 90, 88}, wantErr: false,
		}, {
			name: "zeroed values", args: args{Vector{0, 0, 0, 0}, Vector{5, 4, 3, 2}}, wantVector: Vector{5, 4, 3, 2}, wantErr: false,
		}, {
			name: "nil vector", args: args{Vector{}, Vector{1, 2, 3}}, wantVector: nil, wantErr: true,
		}, {
			name: "nil vector two", args: args{nil, Vector{1, 2, 3}}, wantVector: nil, wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVector, err := AddVectors(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TEST: %v, AddVectors() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotVector, tt.wantVector) {
				t.Errorf("TEST: %v, AddVectors() = %v, want %v", tt.name, gotVector, tt.wantVector)
			}
		})
	}
}

func TestSubtractVectors(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}
	tests := []struct {
		name       string
		args       args
		wantVector Vector
		wantErr    bool
	}{
		{
			name: "same length", args: args{Vector{1, 2, 3}, Vector{10, 10, 10}}, wantVector: Vector{-9, -8, -7}, wantErr: false,
		}, {
			name: "different lengths", args: args{Vector{2, 3}, Vector{10, 10, 10}}, wantVector: nil, wantErr: true,
		}, {
			name: "large values minus small values", args: args{Vector{100, 200, 300}, Vector{10, 20, 30}}, wantVector: Vector{90, 180, 270}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVector, err := SubtractVectors(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TEST: %v, SubtractVectors() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotVector, tt.wantVector) {
				t.Errorf("TEST: %v, SubtractVectors() = %v, want %v", tt.name, gotVector, tt.wantVector)
			}
		})
	}
}

func TestScalarMultiplication(t *testing.T) {
	type args struct {
		scalar float64
		v      Vector
	}
	tests := []struct {
		name       string
		args       args
		wantVector Vector
	}{
		{
			name: "pos scalar pos values", args: args{2, Vector{2, 4, 6, 8, 10}}, wantVector: Vector{4, 8, 12, 16, 20},
		}, {
			name: "neg scalar pos values", args: args{-1, Vector{2, 4, 6, 8, 10}}, wantVector: Vector{-2, -4, -6, -8, -10},
		}, {
			name: "pos scalar neg values", args: args{2, Vector{-1.1, -2.2, -3.3, -4.4}}, wantVector: Vector{-2.2, -4.4, -6.6, -8.8},
		}, {
			name: "neg scalar neg values", args: args{-1, Vector{11, 12, 13}}, wantVector: Vector{-11, -12, -13},
		}, {
			name: "zero scalar", args: args{0.0, Vector{44, 99.999999999, 11, 42, 23.2323232323}}, wantVector: Vector{0, 0, 0, 0, 0},
		}, {
			name: "nil vector", args: args{3.14, nil}, wantVector: nil,
		}, {
			name: "nil vector two", args: args{3.14, Vector{}}, wantVector: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVector := ScalarMultiplication(tt.args.scalar, tt.args.v); !reflect.DeepEqual(gotVector, tt.wantVector) {
				t.Errorf("TEST: %v, ScalarMultiplication() = %v, want %v", tt.name, gotVector, tt.wantVector)
			}
		})
	}
}

func TestDotProduct(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}
	tests := []struct {
		name        string
		args        args
		wantProduct float64
		wantErr     bool
	}{
		{
			name: "pos values", args: args{Vector{1, 2, 3, 4}, Vector{2.2, 2.2, 2.2, 2.2}}, wantProduct: 22, wantErr: false,
		}, {
			name: "neg values", args: args{Vector{1, 2, 3, 4}, Vector{-2.2, -2.2, -2.2, -2.2}}, wantProduct: -22, wantErr: false,
		}, {
			name: "odd number of neg values", args: args{Vector{1, 2, 3, 4}, Vector{2.2, -2.2, -2.2, -2.2}}, wantProduct: -17.6, wantErr: false,
		}, {
			name: "nil vector", args: args{nil, Vector{2.2, 2.2, 2.2, 2.2}}, wantProduct: 0, wantErr: true,
		}, {
			name: "different length vectors", args: args{Vector{2.2, 2.2, 2.2, 2.2}, nil}, wantProduct: 0, wantErr: true,
		}, {
			name: "different length vectors two", args: args{Vector{2.2, 2.2, 2.2, 2.2}, Vector{100000.00000000111212}}, wantProduct: 0, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProduct, err := DotProduct(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TEST: %v, DotProduct() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if gotProduct != tt.wantProduct {
				t.Errorf("TEST: %v, DotProduct() = %v, want %v", tt.name, gotProduct, tt.wantProduct)
			}
		})
	}
}

func TestSquaredDistance(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "large values, small values", args: args{Vector{11, 22, 33, 44}, Vector{1, 2, 3, 4}}, want: 3000, wantErr: false,
		}, {

			name: "small values, large values", args: args{Vector{1, 2, 3, 4}, Vector{11, 22, 33, 44}}, want: 3000, wantErr: false,
		}, {
			name: "even number of neg values", args: args{Vector{-1, -2, -3, -4}, Vector{11, 22, 33, 44}}, want: 4320, wantErr: false,
		}, {
			name: "odd number of neg values", args: args{Vector{1, -2, -3, -4}, Vector{11, 22, 33, 44}}, want: 4276, wantErr: false,
		}, {
			name: "different lengths", args: args{nil, Vector{11, 22, 33, 44}}, want: 0, wantErr: true,
		}, {
			name: "different lengths two", args: args{Vector{}, Vector{11, 22, 33, 44}}, want: 0, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SquaredDistance(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TEST: %v, SquaredDistance() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("TEST: %v, SquaredDistance() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "pos, pos values", args: args{Vector{100, 101, 102, 103, 104, 105}, Vector{0, 1, 2, 3, 4, 5}}, want: 244.94897427831782, wantErr: false,
		}, {
			name: "even number of neg, pos values", args: args{Vector{-100, -101, -102, -103, -104, -105}, Vector{0, 1, 2, 3, 4, 5}}, want: 257.332469774026, wantErr: false,
		}, {
			name: "odd number of neg, pos values", args: args{Vector{-100, -101, -102, -103, -104, 105}, Vector{0, 1, 2, 3, 4, 5}}, want: 253.21927256826245, wantErr: false,
		}, {
			name: "pos, even number of neg values", args: args{Vector{100, 101, 102, 103, 104, 105}, Vector{-0, -1, -2, -3, -4, -5}}, want: 257.332469774026, wantErr: false,
		}, {
			name: "pos, odd number of neg values", args: args{Vector{100, 101, 102, 103, 104, 105}, Vector{-0, -1, -2, 3, -4, -5}}, want: 254.9195951667898, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distance(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TEST: %v, Distance() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("TEST: %v, Distance() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestDeMeanVector(t *testing.T) {
	type args struct {
		v Vector
	}
	tests := []struct {
		name       string
		args       args
		wantVector Vector
	}{
		{
			name: "pos values", args: args{Vector{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}}, wantVector: Vector{-45, -35, -25, -15, -5, 5, 15, 25, 35, 45},
		}, {
			name: "neg values", args: args{Vector{-10, -20, -30, -40, -50, -60, -70, -80, -90, -100}}, wantVector: Vector{45, 35, 25, 15, 5, -5, -15, -25, -35, -45},
		}, {
			name: "mixed pos/neg values", args: args{Vector{-10, 20, -30, 40, -50, 60, -70, 80, -90, 100}}, wantVector: Vector{-15, 15, -35, 35, -55, 55, -75, 75, -95, 95},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVector := DeMeanVector(tt.args.v); !reflect.DeepEqual(gotVector, tt.wantVector) {
				t.Errorf("DeMeanVector() = %v, want %v", gotVector, tt.wantVector)
			}
		})
	}
}

func TestCovariance(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "pos vals with perf covar", args: args{Vector{1, 2, 3}, Vector{4, 5, 6}}, want: 1, wantErr: false,
		}, {
			name: "pos vals with no covar", args: args{Vector{3, 99, 3}, Vector{4, 5, 6}}, want: 0, wantErr: false,
		}, {
			name: "pos vals with neg covar", args: args{Vector{9, 10, 100}, Vector{5, 4, 3}}, want: -45.5, wantErr: false,
		}, {
			name: "neg vals", args: args{Vector{-1, -2, -3}, Vector{1, 2, 3}}, want: -1, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Covariance(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TEST: %v, Covariance() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("TEST: %v, Covariance() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestCorrelation(t *testing.T) {
	type args struct {
		a Vector
		b Vector
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "pos values perf pos correlation", args: args{Vector{1, 2, 3, 4, 5}, Vector{6, 7, 8, 9, 10}}, want: 0.9999999999999999, wantErr: false,
		}, {
			name: "pos values perf neg correlation", args: args{Vector{1, 2, 3, 4, 5}, Vector{10, 9, 8, 7, 6}}, want: -0.9999999999999999, wantErr: false,
		}, {
			name: "neg values perf pos correlation", args: args{Vector{1, 2, 3, 4, 5}, Vector{10, 9, 8, 7, 6}}, want: -0.9999999999999999, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Correlation(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Correlation() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Correlation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Sort(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.Sort()
		})
	}
}

func TestVector_Rsort(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.Rsort()
		})
	}
}

func TestVector_SumValues(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.SumValues(); got != tt.want {
				t.Errorf("Vector.SumValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Mean(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Mean(); got != tt.want {
				t.Errorf("Vector.Mean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_DeMean(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.DeMean()
		})
	}
}

func TestVector_Median(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Median(); got != tt.want {
				t.Errorf("Vector.Median() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Quantile(t *testing.T) {
	type args struct {
		percentile float64
	}
	tests := []struct {
		name         string
		v            Vector
		args         args
		wantQuantile float64
		wantErr      bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuantile, err := tt.v.Quantile(tt.args.percentile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.Quantile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotQuantile != tt.wantQuantile {
				t.Errorf("Vector.Quantile() = %v, want %v", gotQuantile, tt.wantQuantile)
			}
		})
	}
}

func TestVector_Mode(t *testing.T) {
	tests := []struct {
		name     string
		v        Vector
		wantMode Vector
		wantErr  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMode, err := tt.v.Mode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.Mode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotMode, tt.wantMode) {
				t.Errorf("Vector.Mode() = %v, want %v", gotMode, tt.wantMode)
			}
		})
	}
}

func TestVector_Range(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Range(); got != tt.want {
				t.Errorf("Vector.Range() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_InterQuartileRange(t *testing.T) {
	tests := []struct {
		name    string
		v       Vector
		want    float64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.InterQuartileRange()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.InterQuartileRange() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Vector.InterQuartileRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_ScalarMultiply(t *testing.T) {
	type args struct {
		scalar float64
	}
	tests := []struct {
		name string
		v    Vector
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.ScalarMultiply(tt.args.scalar)
		})
	}
}

func TestVector_SumOfSquares(t *testing.T) {
	tests := []struct {
		name             string
		v                Vector
		wantSumofsquares float64
		wantErr          bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSumofsquares, err := tt.v.SumOfSquares()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.SumOfSquares() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotSumofsquares != tt.wantSumofsquares {
				t.Errorf("Vector.SumOfSquares() = %v, want %v", gotSumofsquares, tt.wantSumofsquares)
			}
		})
	}
}

func TestVector_Magnitude(t *testing.T) {
	tests := []struct {
		name    string
		v       Vector
		want    float64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Magnitude()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.Magnitude() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Vector.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Variance(t *testing.T) {
	tests := []struct {
		name    string
		v       Vector
		want    float64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.Variance()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.Variance() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Vector.Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_StandardDeviation(t *testing.T) {
	tests := []struct {
		name    string
		v       Vector
		want    float64
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.StandardDeviation()
			if (err != nil) != tt.wantErr {
				t.Errorf("Vector.StandardDeviation() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Vector.StandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_SumVectors(t *testing.T) {
	tests := []struct {
		name       string
		m          Matrix
		wantVector Vector
		wantErr    bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVector, err := tt.m.SumVectors()
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.SumVectors() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotVector, tt.wantVector) {
				t.Errorf("Matrix.SumVectors() = %v, want %v", gotVector, tt.wantVector)
			}
		})
	}
}

func TestMatrix_MeanVector(t *testing.T) {
	tests := []struct {
		name       string
		m          Matrix
		wantVector Vector
		wantErr    bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVector, err := tt.m.MeanVector()
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.MeanVector() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotVector, tt.wantVector) {
				t.Errorf("Matrix.MeanVector() = %v, want %v", gotVector, tt.wantVector)
			}
		})
	}
}

func TestMatrix_Shape(t *testing.T) {
	tests := []struct {
		name        string
		m           Matrix
		wantRows    int
		wantColumns int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRows, gotColumns := tt.m.Shape()
			if gotRows != tt.wantRows {
				t.Errorf("Matrix.Shape() gotRows = %v, want %v", gotRows, tt.wantRows)
			}
			if gotColumns != tt.wantColumns {
				t.Errorf("Matrix.Shape() gotColumns = %v, want %v", gotColumns, tt.wantColumns)
			}
		})
	}
}

func TestMatrix_GetRow(t *testing.T) {
	type args struct {
		rowNumber int
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want Vector
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.GetRow(tt.args.rowNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.GetRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_GetColumn(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		name       string
		m          Matrix
		args       args
		wantColumn Vector
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotColumn := tt.m.GetColumn(tt.args.columnNumber); !reflect.DeepEqual(gotColumn, tt.wantColumn) {
				t.Errorf("Matrix.GetColumn() = %v, want %v", gotColumn, tt.wantColumn)
			}
		})
	}
}
