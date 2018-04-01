package tendon

import "log"

type Rect struct {
	Left, Top, Right, Bottom float64
}

func (r Rect) Width() float64 {
	return r.Right - r.Left
}

func (r Rect) Height() float64 {
	return r.Bottom - r.Top
}

func (r Rect) Shift(x, y float64) Rect {
	r.Left += x
	r.Right += x
	r.Top += y
	r.Bottom += y
	return r
}

func (r Rect) Contains(p complex128) bool {
	x, y := real(p), imag(p)
	return r.Left < x && x < r.Right && r.Top < y && y < r.Bottom
}

func (r Rect) AnchorPos(anchor int) (float64, float64) {
	var x, y float64
	switch anchor {
	case 1, 4, 7:
		x = r.Left
	case 2, 5, 8:
		x = r.Left + r.Width()/2
	case 3, 6, 9:
		x = r.Right
	}
	switch anchor {
	case 1, 2, 3:
		y = r.Top
	case 4, 5, 6:
		y = r.Top + r.Height()/2
	case 7, 8, 9:
		y = r.Bottom
	}
	return x, y
}

func (r Rect) SnapInside(anchor int, w, h float64) Rect {
	if anchor < 1 || 9 < anchor {
		log.Fatal("WithSnap error. anchor must be between 1 and 9")
	}
	rect := Rect{}
	switch anchor {
	case 1, 4, 7: // snap left
		rect.Left = r.Left
		rect.Right = r.Right + w
	case 2, 5, 8: // snap center
		center := r.Width() / 2
		rect.Left = center - w/2
		rect.Right = center + w/2
	case 3, 6, 9: // snap right
		rect.Right = r.Right
		rect.Left = r.Right - w
	}
	switch anchor {
	case 1, 2, 3: // snap top
		rect.Top = r.Top
		rect.Bottom = r.Bottom + h
	case 4, 5, 6: // snap center
		center := r.Height() / 2
		rect.Top = center - h/2
		rect.Bottom = center + h/2
	case 7, 8, 9: // snap bottom
		rect.Bottom = r.Bottom
		rect.Top = r.Bottom - h
	}
	return rect
}

func (r Rect) SnapOutside(anchor int, w, h float64) Rect {
	if anchor < 1 || 9 < anchor {
		log.Fatal("WithSnap error. anchor must be between 1 and 9")
	}
	rect := Rect{}
	switch anchor {
	case 1, 4, 7: // snap left
		rect.Left = r.Left - w
		rect.Right = r.Left
	case 2, 8: // snap center
		center := r.Width() / 2
		rect.Left = center - w/2
		rect.Right = center + w/2
	case 3, 6, 9: // snap right
		rect.Right = r.Right + w
		rect.Left = r.Right
	}
	switch anchor {
	case 1, 2, 3: // snap top
		rect.Top = r.Top - h
		rect.Bottom = r.Top
	case 4, 6: // snap center
		center := r.Height() / 2
		rect.Top = center - h/2
		rect.Bottom = center + h/2
	case 7, 8, 9: // snap bottom
		rect.Bottom = r.Bottom + h
		rect.Top = r.Bottom
	}
	return rect
}

func (rect Rect) WithMargin(l, t, r, b float64) Rect {
	return Rect{
		Left:   rect.Left + l,
		Top:    rect.Top + t,
		Right:  rect.Right - r,
		Bottom: rect.Bottom - b,
	}
}

func (r Rect) VSplit(ws ...float64) []Rect {
	base := Rect{
		Top:    r.Top,
		Bottom: r.Bottom,
	}
	x := r.Left
	var rs []Rect
	for _, w := range ws {
		rect := base
		rect.Left = x
		x += w
		rect.Right = x
		rs = append(rs, rect)
	}
	// 残った領域を追加
	if x < r.Right {
		rect := base
		rect.Left = x
		rect.Right = r.Right
		rs = append(rs, rect)
	}
	return rs
}

func (r Rect) HSplit(hs ...float64) []Rect {
	base := Rect{
		Left:  r.Left,
		Right: r.Right,
	}
	y := r.Top
	var rs []Rect
	for _, h := range hs {
		rect := base
		rect.Top = y
		y += h
		rect.Bottom = y
		rs = append(rs, rect)
	}
	// 残った領域を追加
	if y < r.Bottom {
		rect := base
		rect.Top = y
		rect.Bottom = r.Bottom
		rs = append(rs, rect)
	}
	return rs
}

func (r Rect) DivideRows(rows int) []Rect {
	h := r.Height() / float64(rows)
	child := Rect{
		Left:  r.Left,
		Top:   r.Top,
		Right: r.Right,
	}
	var rs []Rect
	var bottom float64
	for i := 0; i < rows; i++ {
		bottom = r.Top + h*float64(i+1)
		child.Bottom = bottom
		rs = append(rs, child)
		child.Top = bottom
	}
	return rs
}
