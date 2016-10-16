package cui

/*
View is a type to define rendering of application's models.

This type stores the attributes required for rendering,
such as the size and the position of ui component.
*/
type View struct {
	Size  Size
	Point Point
}

func (view *View) Render() {
	// TODO: Render `view`.
}

type Size struct {
	Width  int
	Height int
}

type Point struct {
	X int
	Y int
}
