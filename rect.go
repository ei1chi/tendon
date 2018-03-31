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
	res := Rect{}
	switch anchor {
	case 1, 4, 7: // snap left
		res.Left = r.Left
		res.Right = r.Right + w
	case 2, 5, 8: // snap center
		center := r.Width() / 2
		res.Left = center - w/2
		res.Right = center + w/2
	case 3, 6, 9: // snap right
		res.Right = r.Right
		res.Left = r.Right - w
	}
	switch anchor {
	case 1, 2, 3: // snap top
		res.Top = r.Top
		res.Bottom = r.Bottom + h
	case 4, 5, 6: // snap center
		center := r.Height() / 2
		res.Top = center - h/2
		res.Bottom = center + h/2
	case 7, 8, 9: // snap bottom
		res.Bottom = r.Bottom
		res.Top = r.Bottom - h
	}
	return res
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
	left := r.Left
	var rs []Rect
	for _, w := range ws {
		res := base
		res.Left = left
		left += w
		res.Right = left
		rs = append(rs, res)
	}
	// 残った領域を追加
	if left < r.Right {
		res := base
		res.Left = left
		res.Right = r.Width()
		rs = append(rs, res)
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
