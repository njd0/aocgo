package vector

import "fmt"

type Vector struct {
	X, Y int
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vector) Sub(v2 Vector) Vector {
	return Vector{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func (v1 Vector) Equal(v2 Vector) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v Vector) String() string {
	return fmt.Sprintf("Vector(X: %.2f, Y: %.2f)", v.X, v.Y)
}