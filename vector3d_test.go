package vectormath

import "testing"

func TestAdd3D(t *testing.T) {
	var cases = []struct {
		name       string
		x1, y1, z1 float64
		x2, y2, z2 float64
		r1, r2, r3 float64
	}{
		{"addition-1", 1, 0, 0, 2, 1, 1, 3, 1, 1},
		{"addition-2", -1, 0, 0, 1, 0, 0, 0, 0, 0},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v1 := Vector3D{X: tt.x1, Y: tt.y1, Z: tt.z1}
			v2 := Vector3D{X: tt.x2, Y: tt.y2, Z: tt.z2}
			v3 := Add3D(v1, v2)
			if v3.X != tt.r1 || v3.Y != tt.r2 || v3.Z != tt.r3 {
				t.Errorf("Addition failed for testcase: %#v", tt)
			}
		})
	}
}

func TestMultiplyByScalar3D(t *testing.T) {
	var cases = []struct {
		name       string
		x1, y1, z1 float64
		s          float64
		r1, r2, r3 float64
	}{
		{"product-1", 1, 0, 0, 2, 2, 0, 0},
		{"product-2", 1, 0, 0, -1, -1, 0, 0},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v1 := Vector3D{X: tt.x1, Y: tt.y1, Z: tt.z1}
			v2 := MultiplyByScalar3D(v1, tt.s)
			if v2.X != tt.r1 || v2.Y != tt.r2 || v2.Z != tt.r3 {
				t.Errorf("Scalar multiplication failed for testcase: %#v", tt)
			}
		})
	}
}

func TestDotProduct3D(t *testing.T) {
	var cases = []struct {
		name       string
		x1, y1, z1 float64
		x2, y2, z2 float64
		exp        float64
	}{
		{"dot-product-1", 1, 0, 0, 1, 0, 0, 1},
		{"dot-product-2", 1, 1, 1, -1, 1, 1, 1},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v1 := Vector3D{X: tt.x1, Y: tt.y1, Z: tt.z1}
			v2 := Vector3D{X: tt.x2, Y: tt.y2, Z: tt.z2}
			act := DotProduct3D(v1, v2)
			if act != tt.exp {
				t.Errorf("Dot product failed for testcase: %#v", tt)
			}
		})
	}
}

func TestCrossProduct3D(t *testing.T) {
	var cases = []struct {
		name       string
		x1, y1, z1 float64
		x2, y2, z2 float64
		r1, r2, r3 float64
	}{
		{"cross-product-1", 1, 0, 0, 0, 1, 0, 0, 0, 1},
		{"cross-product-2", 0, 1, 2, 0, 3, 4, -2, 0, 0},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v1 := Vector3D{X: tt.x1, Y: tt.y1, Z: tt.z1}
			v2 := Vector3D{X: tt.x2, Y: tt.y2, Z: tt.z2}
			v3 := CrossProduct3D(v1, v2)
			if v3.X != tt.r1 || v3.Y != tt.r2 || v3.Z != tt.r3 {
				t.Errorf("Cross product failed for testcase: %#v", tt)
			}
		})
	}
}
