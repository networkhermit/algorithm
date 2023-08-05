package tests

import (
	"testing"

	"muse/algorithms/number_theory/tests"
)

type Category byte

const (
	Even Category = iota
	Odd
)

var sample = []tests.UniqueCategorySample[int64, Category]{
	{N: 0, Category: Even},
	{N: 1, Category: Odd},
	{N: -1, Category: Odd},
	{N: 60, Category: Even},
	{N: 62, Category: Even},
	{N: 42, Category: Even},
	{N: -87, Category: Odd},
	{N: -98, Category: Even},
	{N: -56, Category: Even},
	{N: 41, Category: Odd},
	{N: 14, Category: Even},
	{N: 20, Category: Even},
	{N: -63, Category: Odd},
	{N: -81, Category: Odd},
	{N: -39, Category: Odd},
	{N: 33, Category: Odd},
	{N: -32, Category: Even},
	{N: -50, Category: Even},
	{N: -1471, Category: Odd},
	{N: 4046, Category: Even},
	{N: 5019, Category: Odd},
	{N: 4521, Category: Odd},
	{N: -8744, Category: Even},
	{N: 4787, Category: Odd},
	{N: 4973, Category: Odd},
	{N: -1184, Category: Even},
	{N: -4437, Category: Odd},
	{N: -5871, Category: Odd},
	{N: -4298, Category: Even},
	{N: -2027, Category: Odd},
	{N: -2796, Category: Even},
	{N: -2209, Category: Odd},
	{N: -6094, Category: Even},
	{N: 3257, Category: Odd},
	{N: -4732, Category: Even},
	{N: 7495, Category: Odd},
	{N: -3916, Category: Even},
	{N: 1469, Category: Odd},
	{N: 6164, Category: Even},
	{N: -7545, Category: Odd},
	{N: -7763, Category: Odd},
	{N: 7523, Category: Odd},
	{N: -8076, Category: Even},
	{N: -3778, Category: Even},
	{N: -1648, Category: Even},
	{N: 4220, Category: Even},
	{N: -8551, Category: Odd},
	{N: 9692, Category: Even},
	{N: 5394, Category: Even},
	{N: 2472, Category: Even},
	{N: 4056, Category: Even},
	{N: 5769, Category: Odd},
	{N: -2322, Category: Even},
	{N: 503, Category: Odd},
	{N: -8721, Category: Odd},
	{N: -6344, Category: Even},
	{N: -4335, Category: Odd},
	{N: 1677, Category: Odd},
	{N: -1703, Category: Odd},
	{N: -4086, Category: Even},
	{N: 7076, Category: Even},
	{N: -7165, Category: Odd},
	{N: 7636, Category: Even},
	{N: -8043, Category: Odd},
	{N: -3753, Category: Odd},
	{N: 4007, Category: Odd},
	{N: -261, Category: Odd},
	{N: -6538, Category: Even},
	{N: 9766, Category: Even},
	{N: -7563, Category: Odd},
	{N: -7944, Category: Even},
	{N: 8922, Category: Even},
	{N: -5759, Category: Odd},
	{N: -8791, Category: Odd},
	{N: -2211, Category: Odd},
	{N: 3493, Category: Odd},
	{N: 5573, Category: Odd},
	{N: -2645, Category: Odd},
	{N: -603656, Category: Even},
	{N: 807727, Category: Odd},
	{N: -69847, Category: Odd},
	{N: -843676, Category: Even},
	{N: -830961, Category: Odd},
	{N: -608772, Category: Even},
	{N: 931043, Category: Odd},
	{N: 855512, Category: Even},
	{N: 358482, Category: Even},
	{N: -98919, Category: Odd},
	{N: 215211, Category: Odd},
	{N: -933334, Category: Even},
	{N: -613634, Category: Even},
	{N: -95643, Category: Odd},
	{N: 53934, Category: Even},
	{N: 161818, Category: Even},
	{N: 67041621, Category: Odd},
	{N: 99662694, Category: Even},
	{N: -94392019, Category: Odd},
	{N: -20543495, Category: Odd},
	{N: -27591794, Category: Even},
	{N: -8314396, Category: Even},
	{N: 97455764, Category: Even},
	{N: 59367920, Category: Even},
	{N: 26856309, Category: Odd},
	{N: 64178815, Category: Odd},
	{N: -11480741, Category: Odd},
	{N: 45428276, Category: Even},
	{N: 46193175, Category: Odd},
	{N: -31079636, Category: Even},
	{N: 63115980, Category: Even},
	{N: 42559270, Category: Even},
	{N: -1645871504, Category: Even},
	{N: -1233918598, Category: Even},
	{N: 722012346, Category: Even},
	{N: -1525999934, Category: Even},
	{N: -365543955, Category: Odd},
	{N: 2008798151, Category: Odd},
	{N: -1300713468, Category: Even},
	{N: 1425587979, Category: Odd},
	{N: 1324445673, Category: Odd},
	{N: 2136612365, Category: Odd},
	{N: -995371213, Category: Odd},
	{N: -2048365905, Category: Odd},
	{N: 2096138065, Category: Odd},
	{N: -768738192, Category: Even},
	{N: -846034014, Category: Even},
	{N: 411817058, Category: Even},
	{N: 2147483647, Category: Odd},
	{N: -2147483648, Category: Even},
}

var Derive = func(fn func(int64) bool, c Category) func(t *testing.T) {
	return tests.UniqueCategoryDerive(fn, sample, c)
}