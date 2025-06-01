package somechatgptexercises

// Crea una estructura Rectangle con Width y Height.
//  Añade métodos para calcular el área y el perímetro usando punteros.

type Rectangle struct {
	Width  float64
	Height float64
}

func NewRectangle(w, h float64) *Rectangle {

	rectangle := Rectangle{
		Width:  w,
		Height: h,
	}
	return &rectangle
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return r.Height*2 + r.Width*2
}
