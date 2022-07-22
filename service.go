package excercise4

type Converter struct {
}

func (c Converter) ConvertToC(fah float64) float64 {
	return (fah - 32) * 5 / 9
}

func (c Converter) ConvertToF(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}
