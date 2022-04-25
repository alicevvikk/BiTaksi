package utils 

import (
	"testing"
	"math"
)

var tests = []struct {
	c1     []float64
	c2     []float64
	dist float64
}{
	{
		[]float64{22.55,  43.12},  // Rio de Janeiro, Brazil
		[]float64{13.45,  100.28}, // Bangkok, Thailand
		6094.544408786774,
	},
	{
		[]float64{20.10,57.30}, // Port Louis, Mauritius
		[]float64{0.57, 100.21}, // Padang, Indonesia
		5145.525771394785,
	},
	{
		[]float64{51.45, 1.15},  // Oxford, United Kingdom
		[]float64{41.54, 12.27}, // Vatican, City Vatican City
		1389.1793118293067,
	},
	{
		[]float64{22.34, 17.05}, // Windhoek, Namibia
		[]float64{51.56, 4.29},  // Rotterdam, Netherlands
		3429.89310043882,
	},
	{
		[]float64{63.24, 56.59}, // Esperanza, Argentina
		[]float64{8.50,  13.14},  // Luanda, Angola
		6996.18595539861,
	},
	{
		[]float64{90.00, 0.00}, // North/South Poles
		[]float64{48.51, 2.21}, // Paris,  France
		4613.477506482742,
	},
	{
		[]float64{45.04, 7.42},  // Turin, Italy
		[]float64{3.09, 101.42}, // Kuala Lumpur, Malaysia
		10078.111954385415,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range tests {
		km := CalculateDistance(input.c1, input.c2)

		if math.Abs(input.dist - km) > 0.1 {
			t.Errorf("fail: want %v %v -> %v got %v",
				input.c1,
				input.c2,
				input.dist,
				km,
			)
		}
	}
}
