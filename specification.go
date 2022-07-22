package excercise4

import (
	"excercise4/assert"
	"testing"
)

type TempConverter interface {
	ConvertToF(celsius float64) (float64, error)
	ConvertToC(fah float64) (float64, error)
}

func ConverterSpecification(t *testing.T, magicBox TempConverter) {
	t.Run("can convert celsius to fa", func(t *testing.T) {
		t.Run("0", func(t *testing.T) {
			celsius := 0.0

			fah, err := magicBox.ConvertToF(celsius)
			assert.NoError(t, err)

			assert.Equal(t, fah, 32)

		})
		t.Run("15", func(t *testing.T) {
			celsius := 15.0

			fah, err := magicBox.ConvertToF(celsius)
			assert.NoError(t, err)

			assert.Equal(t, fah, 59)
		})
	})
	t.Run("can convert fa to celsius", func(t *testing.T) {
		t.Skip()
		fah := 32.0

		celsius, err := magicBox.ConvertToC(fah)
		assert.NoError(t, err)

		if celsius != 0 {
			t.Errorf("%v is not 0", celsius)
		}
	})
}
