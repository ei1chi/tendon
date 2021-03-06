package tendon

func Bezier3d(a, b, c, d complex128, time float64) complex128 {
	t := complex(time, 0)
	T := complex(1-time, 0)
	return T*T*T*a + 3*T*T*t*b + 3*T*t*t*c + t*t*t*d
}

func EaseHH(begin, end, time float64) float64 {
	a := complex(0, 0)
	b := complex(begin, 0)
	c := complex(end, 1)
	d := 1 + 1i
	r := Bezier3d(a, b, c, d, time)
	return imag(r)
}

func EaseVV(begin, end, time float64) float64 {
	a := complex(0, 0)
	b := complex(0, begin)
	c := complex(1, end)
	d := 1 + 1i
	r := Bezier3d(a, b, c, d, time)
	return imag(r)
}

func EaseVH(begin, end, time float64) float64 {
	a := complex(0, 0)
	b := complex(0, begin)
	c := complex(end, 1)
	d := 1 + 1i
	r := Bezier3d(a, b, c, d, time)
	return imag(r)
}

func EaseHV(begin, end, time float64) float64 {
	a := complex(0, 0)
	b := complex(begin, 0)
	c := complex(1, end)
	d := 1 + 1i
	r := Bezier3d(a, b, c, d, time)
	return imag(r)
}
