package tendon

func Bezier3d(a, b, c, d complex128, time float64) complex128 {
	t := complex(time, 0)
	T := complex(1-time, 0)
	return T*T*T*a + 3*T*T*t*b + 3*T*t*t*c + t*t*t*d
}
