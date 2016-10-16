package cui

type Size struct {
	Width  int
	Height int
}

type Point struct {
	X int
	Y int
}

func PointZero() Point {
	return Point{
		X: 0,
		Y: 0,
	}
}
