package excercise4

import "testing"

func TestConvert(t *testing.T) {

	t.Run("c to f", func(t *testing.T) {
		ConverterSpecification(t, Converter{})
	})
}
