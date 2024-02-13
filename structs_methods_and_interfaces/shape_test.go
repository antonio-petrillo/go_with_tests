package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64){
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func (t *testing.T) {
		r := Rectangle{10.0, 10.0}
		want := 40.0
		checkPerimeter(t, r, want)
	})

	t.Run("circles", func (t *testing.T) {
		c := Circle{10.0}
		want := 62.83185307179586
		checkPerimeter(t, c, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64){
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func (t *testing.T) {
		r := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, r, want)
	})

	t.Run("circles", func (t *testing.T) {
		c := Circle{10}
		want := 314.1592653589793
		checkArea(t, c, want)
	})

	t.Run("triangles", func (t *testing.T) {
		tr := Triangle{12.0, 6.0}
		want := 36.0
		checkArea(t, tr, want)
	})
}
