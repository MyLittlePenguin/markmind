package algebra

import (
	"markmind/internal/core/algebra"
	"testing"
)

func TestV2Addition(t *testing.T) {
	v1 := algebra.NewVec2D(20, 2)
	v2 := algebra.NewVec2D(10, 5)
	v3 := v1.Add(v2)
	if !v3.Equals(algebra.NewVec2D(30, 7)) {
		t.Fatalf("Fuck... addition doesn't work (%v) + (%v) = (%v)", v1, v2, v3)
	}
}

func TestV3Addition(t *testing.T) {
	v1 := algebra.NewVec3D(20, 2, 1)
	v2 := algebra.NewVec3D(10, 5, 1)
	v3 := v1.Add(v2)
	if !v3.Equals(algebra.NewVec3D(30, 7, 2)) {
		t.Fatalf("Fuck... addition doesn't work (%v) + (%v) = (%v)", v1, v2, v3)
	}
}

func TestM2xV2(t *testing.T) {
	v := algebra.NewVec2D(3, 2)
	m := algebra.NewMatrix2DFromRows(
		[2]float64{3, 1},
		[2]float64{2, 4},
	)
	r := m.Multiply(v)
	expectation := algebra.NewVec2D(
		9+2,
		6+8,
	)
	if !r.Equals(expectation) {
		t.Fatalf("OMG... Multiplication is screwed (%v) * (%v) = (%v)", m, v, r)
	}
}

func TestM3xV3(t *testing.T) {
	v := algebra.NewVec3D(3, 2, 1)
	m := algebra.NewMatrix3DFromRows(
		[3]float64{3, 1, 0},
		[3]float64{2, 4, 0},
		[3]float64{0, 0, 1},
	)
	r := m.Multiply(v)
	expectation := algebra.NewVec3D(
		9+2,
		6+8,
		1,
	)
	if !r.Equals(expectation) {
		t.Fatalf("OMG... Multiplication is screwed (%v) * (%v) = (%v)", m, v, r)
	}
}

func TestM2xM2(t *testing.T) {
  m1 := algebra.NewMatrix2DFromRows(
    [2]float64{2, 3},
    [2]float64{4, 5},
  )
  m2 := algebra.NewMatrix2DFromRows(
    [2]float64{1, -1},
    [2]float64{-1, 1},
  )

  r := m1.MultiplyMatrix(m2)
  expectation := algebra.NewMatrix2DFromRows(
    [2]float64{-1, 1},
    [2]float64{-1, 1},
  )
  if !r.Equals(expectation) {
    t.Fatalf("Matrix Multiplication is just wrong (%v) * (%v) = (%v)\n", m1, m2, r)
  }
}

func TestIdentity(t *testing.T) {
  identity := algebra.Identity2D()

  expectation := algebra.NewMatrix2DFromRows(
    [2]float64{1, 0},
    [2]float64{0, 1},
  )
  if !identity.Equals(expectation) {
    t.Fatal("We have an Identity Crisis")
  }

}

func TestNewMatrix2D(t *testing.T)  {
  m := algebra.NewMatrix2D(

  )
}

func TestIdentityMultiplication(t *testing.T) {
  i := algebra.Identity2D()
  v := algebra.NewMatrix2D(2, 3)
  r := i.Multiply(v)
  expectation := algebra.NewMatrix2DFromRows(
    [2]float64{2, 0},
    [2]float64{0, 3},
  )

  if !r.Equals(expectation) {
    t.Fatal("The other kind of identity crisis")
  }
}
