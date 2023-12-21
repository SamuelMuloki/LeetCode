package solutions

func ConvertTemperature(celsius float64) []float64 {
	kelvin := celsius + float64(273.15)
	fahrenheit := celsius*float64(1.80) + float64(32.00)
	return []float64{kelvin, fahrenheit}
}
