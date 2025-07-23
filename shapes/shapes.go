package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectagle Rectangle) float64 {
	return 2 * (rectagle.Height + rectagle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
