package utils

type Vector2 struct {
	X, Y int
}

func (v *Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
