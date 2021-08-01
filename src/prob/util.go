package prob

func NormalizedFloat64(v, min, max float64) float64 {
	return (v - min) / (max - min)
}
