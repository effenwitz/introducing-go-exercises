package math

//Average is a function which returns average of numbers passed to it.
func Average(xs []float64) float64 {
	total := float64(0)
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
