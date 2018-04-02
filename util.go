package tendon

import (
	"math"
	"math/cmplx"
)

const Around = 4.0

func Powi(angle float64) complex128 {
	return cmplx.Pow(1i, complex(angle, 0))
}

func AbsSq(c complex128) float64 {
	return math.Pow(real(c), 2) + math.Pow(imag(c), 2)
}

func DisplayScale(scrw, scrh float64) float64 {
	w, h, err := DeviceSize()
	s := 1.0
	if err != nil {
		s, sh := w/scrw, h/scrh
		if s < sh {
			s = sh
		}
	}
	return s
}
