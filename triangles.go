package triangles

// New creates a new triangle with sides a, b and c
func New(a, b, c int) *Triangle {
	return &Triangle{a, b, c}
}

// Valid checks if the sides form a valid triangle
func (t *Triangle) Valid() bool {
	if t.a <= t.c-t.b || t.a <= t.b-t.c || t.b <= t.a-t.c {
		return false
	}
	return true
}

// Type returns the type of the triangle
func (t *Triangle) Type() TriangleType {
	if !t.Valid() {
		return INVALID
	}
	if t.a == t.b && t.b == t.c {
		return EQUILATERAL
	}
	if t.a == t.b || t.a == t.c || t.b == t.c {
		return ISOSCELES
	}
	return SCALENE
}

// Triangle represents a triangle
type Triangle struct {
	a int
	b int
	c int
}

// TriangleType is the type of the triangle.
type TriangleType int

func (t TriangleType) String() string {
	names := []string{"EQUILATERAL", "ISOSCELES", "SCALENE", "INVALID"}
	return names[t]
}

// Constants for triangle types
const (
	EQUILATERAL TriangleType = iota
	ISOSCELES
	SCALENE
	INVALID
)
