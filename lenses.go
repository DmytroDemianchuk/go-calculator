package main

var sensors = []string{
	// Small format
	"Fullframe (35mm)",
	"APS-H (Canon)",
	"APS-C (Generic)",
	"APS-C (Canon)",
	"Micro Four Thirds",
	"1\"",
	// Medium format
	"6x17",
	"6x12",
	"6x9",
	"6x7",
	"6x6",
	"645 (6x4.5)",
	// Large format
	"8x10",
	"5x7",
	"4x5",
}

var sensorToCoC = map[string]float64{
	// Small format
	"Fullframe (35mm)":  0.029e-3,
	"APS-H (Canon)":     0.023e-3,
	"APS-C (Generic)":   0.019e-3,
	"APS-C (Canon)":     0.018e-3,
	"Micro Four Thirds": 0.015e-3,
	"1\"":               0.011e-3,

	// Medium format
	"6x17":        0.12e-3,
	"6x12":        0.083e-3,
	"6x9":         0.067e-3,
	"6x7":         0.059e-3,
	"6x6":         0.053e-3,
	"645 (6x4.5)": 0.047e-3,

	// Large format
	"8x10": 0.22e-3,
	"5x7":  0.15e-3,
	"4x5":  0.11e-3,
}

func depthOfField(focallength, distance, aperture, circle float64) float64 {
	return 2 * distance * distance * aperture * circle / (focallength * focallength)
}
