package cui

/*
View is a type to define rendering of application's models.

This type stores the attributes required for rendering,
such as the size and the position of ui component.
*/
type View struct {
	Size      Size
	Point     Point
	Renderers []Renderer
}

/*
Render view on the screen.

The first argument `point` is the parent view's point.
*/
func (view *View) Render(point Point) {
	// TODO: Render `view`.

	if view.Renderers == nil {
		return
	}

	viewPoint := Point{X: view.Point.X + point.X, Y: view.Point.Y + point.Y}
	for _, r := range view.Renderers {
		r.Render(viewPoint)
	}
}
