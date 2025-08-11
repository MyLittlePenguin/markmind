package algebra

import "math"

type Vec2D struct{ X, Y float64 }
type Vec3D struct{ X, Y, Z float64 }

type Matrix2D [2]Vec2D
type Matrix3D [3]Vec3D

func NewVec2D(x, y float64) Vec2D { return Vec2D{x, y} }
func NewZeroVec2D() Vec2D         { return Vec2D{0, 0} }

func NewVec3D(x, y, z float64) Vec3D { return Vec3D{x, y, z} }
func NewZeroVec3D() Vec3D            { return Vec3D{0, 0, 0} }

func NewMatrix2DFromRows(rv1, rv2 [2]float64) Matrix2D {
	return Matrix2D{
		Vec2D{rv1[0], rv2[0]},
		Vec2D{rv1[1], rv2[1]},
	}
}

func NewMatrix3DFromRows(rv1, rv2, rv3 [3]float64) Matrix3D {
	return Matrix3D{
		Vec3D{rv1[0], rv2[0], rv3[0]},
		Vec3D{rv1[1], rv2[1], rv3[1]},
		Vec3D{rv1[2], rv2[2], rv3[2]},
	}
}

func NewMatrix2D(v1, v2 Vec2D) Matrix2D     { return Matrix2D{v1, v2} }
func NewMatrix3D(v1, v2, v3 Vec3D) Matrix3D { return Matrix3D{v1, v2, v3} }

func (self Vec2D) Compare(other Vec2D) int {
	diff := self.Distance() - other.Distance()
	if diff < 0 {
		return -1
	} else if diff > 0 {
		return 1
	} else {
		return 0
	}
}

func (self Vec3D) Compare(other Vec3D) int {
	diff := self.Distance() - other.Distance()
	if diff < 0 {
		return -1
	} else if diff > 0 {
		return 1
	} else {
		return 0
	}
}

func (self Vec2D) Equals(other Vec2D) bool {
	return self.X == other.X && self.Y == other.Y
}

func (self Vec3D) Equals(other Vec3D) bool {
	return self.X == other.X && self.Y == other.Y && self.Z == other.Z
}

func (self Matrix2D) Equals(other Matrix2D) bool {
	return self[0].Equals(other[0]) && self[1].Equals(other[1])
}

func (self Matrix3D) Equals(other Matrix3D) bool {
	return self[0].Equals(other[0]) && self[1].Equals(other[1]) && self[2].Equals(other[2])
}

func (self Vec2D) Add(other Vec2D) Vec2D {
	return Vec2D{
		self.X + other.X,
		self.Y + other.Y,
	}
}

func (self Vec3D) Add(other Vec3D) Vec3D {
	return Vec3D{
		self.X + other.X,
		self.Y + other.Y,
		self.Z + other.Z,
	}
}

func (self Vec2D) Subtract(other Vec2D) Vec2D {
	return Vec2D{
		self.X - other.X,
		self.Y - other.Y,
	}
}

func (self Vec3D) Subtract(other Vec3D) Vec3D {
	return Vec3D{
		self.X - other.X,
		self.Y - other.Y,
		self.Z - other.Z,
	}
}

func (self Vec2D) Distance() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y)
}

func (self Vec3D) Distance() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y + self.Z*self.Z)
}

func (self Matrix2D) Multiply(other Vec2D) Vec2D {
	return Vec2D{
		self[0].X*other.X + self[1].X*other.Y,
		self[0].Y*other.X + self[1].Y*other.Y,
	}
}

func (self Matrix3D) Multiply(other Vec3D) Vec3D {
	return Vec3D{
		self[0].X*other.X + self[1].X*other.Y + self[2].X*other.Z,
		self[0].Y*other.X + self[1].Y*other.Y + self[2].Y*other.Z,
		self[0].Z*other.X + self[1].Z*other.Y + self[2].Z*other.Z,
	}
}

func (self Matrix2D) MultiplyMatrix(other Matrix2D) Matrix2D {
	return Matrix2D{
		Vec2D{
			self[0].X*other[0].X + self[1].X*other[0].Y,
			self[0].Y*other[0].X + self[1].Y*other[0].Y,
		},
		Vec2D{
			self[0].X*other[1].X + self[1].X*other[1].Y,
			self[0].Y*other[1].X + self[1].Y*other[1].Y,
		},
	}
}

func (self Matrix3D) MultiplyMatrix(other Matrix3D) Matrix3D {
	return Matrix3D{
		Vec3D{
			self[0].X*other[0].X + self[1].X*other[0].Y + self[2].X*other[0].Z,
			self[0].Y*other[0].X + self[1].Y*other[0].Y + self[2].Y*other[0].Z,
			self[0].Z*other[0].X + self[1].Z*other[0].Y + self[2].Z*other[0].Z,
		},
		Vec3D{
			self[0].X*other[1].X + self[1].X*other[1].Y + self[2].X*other[1].Z,
			self[0].Y*other[1].X + self[1].Y*other[1].Y + self[2].Y*other[1].Z,
			self[0].Z*other[1].X + self[1].Z*other[1].Y + self[2].Z*other[1].Z,
		},
		Vec3D{
			self[0].X*other[2].X + self[1].X*other[2].Y + self[2].X*other[2].Z,
			self[0].Y*other[2].X + self[1].Y*other[2].Y + self[2].Y*other[2].Z,
			self[0].Z*other[2].X + self[1].Z*other[2].Y + self[2].Z*other[2].Z,
		},
	}
}

func Identity2D() Matrix2D {
	return NewMatrix2DFromRows(
		[2]float64{1, 0},
		[2]float64{0, 1},
	)
}

func Identity3D() Matrix3D {
	return NewMatrix3DFromRows(
		[3]float64{1, 0, 0},
		[3]float64{0, 1, 0},
		[3]float64{0, 0, 1},
	)
}

func (self Vec2D) To3D() Vec3D {
	return Vec3D{
		self.X,
		self.Y,
		1,
	}
}

func (self Matrix2D) To3D() Matrix3D {
	m := Matrix3D{
		self[0].To3D(), self[1].To3D(), Vec3D{0, 0, 1},
	}
	m[0].Z = 0
	m[1].Z = 0
	return m
}
