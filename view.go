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

/*
Render view on the screen.

The first argument `point` is the parent view's point.
*/
func (view *View) Render(point Point) {
	// TODO: Render `view`.
}
