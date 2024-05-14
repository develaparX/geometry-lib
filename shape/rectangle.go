package shape

//struct sebagai modul
type Rectangle struct {
	Width  float32
	Length float32
}

// method
func (rectangle *Rectangle) Area() float32 {
	return rectangle.Width * rectangle.Length
}

func (rectangle *Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Width + rectangle.Length)
}

//field dan method harus huruf besar
