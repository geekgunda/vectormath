package vectormath

// Vector3D represents a 3 dimensional vector
type Vector3D struct {
	X, Y, Z float64
}

// Add3D returns the sum of two 3-dimensional vectors
func Add3D(v1, v2 Vector3D) Vector3D {
	return Vector3D{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

// MultiplyByScalar returns the product of a 3-dimensional vector and a scalar
func MultiplyByScalar3D(v Vector3D, s float64) Vector3D {
	return Vector3D{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

// DotProduct3D handles scalar product of two 3-dimensional vectors
// Ref: https://en.wikipedia.org/wiki/Euclidean_vector#Scalar_multiplication
func DotProduct3D(v1, v2 Vector3D) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// CrossProduct3D handles vector product of two 3-dimensional vectors
// Ref: https://en.wikipedia.org/wiki/Euclidean_vector#Cross_product
func CrossProduct3D(v1, v2 Vector3D) Vector3D {
	return Vector3D{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}
