package mlscratchlib

import "math"

// UniformProbabilityDistribution accepts a number float64 and
// returns 1 if the number is between 0 and 1
// returns 0 otherwise
// https://en.wikipedia.org/wiki/Uniform_distribution_(continuous)#Probability_density_function
func UniformProbabilityDistribution(number float64) float64 {
	if number >= 0 && number < 1 {
		return 1
	}
	return 0
}

// UniformCumulativeDensity accepts a float64 and returns a float64
// that represents the probability that a random number is less than
// or equal to some value
func UniformCumulativeDensity(number float64) float64 {
	if number < 0 {
		return 0 // uniform distribution random numbers are never less than 0
	} else if number < 1 {
		return number // ex. Probability of number <= 0.4 is 0.4
	}
	return 1 // uniform distribution random numbers are always less than 1
}

// NormalProbabilityDistribution accepts thre float64 arguments.
// the first of which is the number to return the npdf for
// the other arguments are the mean and standard deviation
// returns a float64
// the return value is intended for use in range type functions to
// plot a set of data on a curve, the mean value shifts the entire
// curve left to right on an x,y grid -- the sigma value affects
// the center of the curve
func NormalProbabilityDistribution(number float64, mean float64, sigma float64) float64 {
	base := (1 / (sigma * math.Sqrt((math.Pi * 2))))
	exponent := -(math.Pow(-(number-mean), 2) / 2 * math.Pow(sigma, 2))
	return base * math.Exp(exponent)
}

// NormalCDF accepts a number, mean, and std deviation as a float
// returns a float64
// useful for graphing the cumulative normal distribution of a
// vector's elements
func NormalCDF(number, mean, sigma float64) (cdf float64) {
	return (1 + math.Erf((number-mean)/math.Sqrt(2)/sigma)) / 2
}

// InverseNormalCDF finds the approximate inverse and returns a float
// that is within the tolerance range of the given probability
// this functions calls itself again if mean and sigma are not
// the standard values 0 and 1 - then re-scales the return value
func InverseNormalCDF(probability, mean, sigma, tolerance float64) (mid float64) {
	if mean != 0 || sigma != 1 {
		return mean + sigma*InverseNormalCDF(probability, 0, 1, tolerance)
	}
	low, hi, mid := -10.0, 10.0, 0
	for hi-low > tolerance {
		mid = (low + hi) / 2
		midp := NormalCDF(mid, 0, 1)

		if midp < probability {
			// didn't find it so search above the mid point
			low = mid
		} else if midp > probability {
			// din't find it so search below the mid point
			hi = mid
		} else {
			// found the probability at the midpoint
			break
		}
	}
	return mid
}
