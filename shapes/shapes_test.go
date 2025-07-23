package shapes

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("test rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}

		got := rectangle.Perimeter()
		want := 40.0

		assertCorrectMessage(t, got, want)
	})

	t.Run("test circle perimeter", func(t *testing.T) {
		circle := Circle{10.0}

		got := circle.Perimeter()
		want := 20 * math.Pi

		assertCorrectMessage(t, got, want)
	})

}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("test rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0

		checkArea(t, &rectangle, want)
	})

	t.Run("test circle area", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, &circle, want)
	})

}

func assertCorrectMessage(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
