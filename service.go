package excercise4

type Converter struct {
}

func (c Converter) ConvertToC(fah float64) (float64, error) {
	return (fah - 32) * 5 / 9, nil
}

func (c Converter) ConvertToF(celsius float64) (float64, error) {
	return (celsius * 9 / 5) + 32, nil
}
