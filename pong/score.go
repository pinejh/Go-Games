package main

import sf "github.com/manyminds/gosfml"

type score struct {
	rs []*sf.RectangleShape
}

func NewScore(x, y float32) *score {
	r1, _ := sf.NewRectangleShape()
	r2, _ := sf.NewRectangleShape()
	r3, _ := sf.NewRectangleShape()
	r4, _ := sf.NewRectangleShape()
	r5, _ := sf.NewRectangleShape()
	r6, _ := sf.NewRectangleShape()
	r7, _ := sf.NewRectangleShape()
	r8, _ := sf.NewRectangleShape()
	r9, _ := sf.NewRectangleShape()
	r10, _ := sf.NewRectangleShape()
	r11, _ := sf.NewRectangleShape()
	r12, _ := sf.NewRectangleShape()
	r13, _ := sf.NewRectangleShape()
	r14, _ := sf.NewRectangleShape()
	r15, _ := sf.NewRectangleShape()
	rs := []*sf.RectangleShape{r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, r13, r14, r15}
	for i := 0; i < 5; i++ {
		for z := 0; z < 3; z++ {
			rs[i*3+z].SetSize(sf.Vector2f{6, 6})
			rs[i*3+z].SetFillColor(sf.Color{255, 255, 255, 255})
			rs[i*3+z].SetOrigin(sf.Vector2f{3, 3})
			rs[i*3+z].SetPosition(sf.Vector2f{x + float32(z-1)*6, y + float32(i-2)*6})
		}
	}
	s := &score{rs}
	s.SetNumber(0)
	return s
}

func (s *score) Update() {
	for i := 0; i < 15; i++ {
		Window.Draw(s.rs[i], sf.DefaultRenderStates())
	}
}

func (s *score) SetNumber(num int) {
	switch num {
	case 0:
		s.toggleRects([15]int{1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1})
		break
	case 1:
		s.toggleRects([15]int{0, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1})
		break
	case 2:
		s.toggleRects([15]int{0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1})
		break
	case 3:
		s.toggleRects([15]int{1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1})
		break
	case 4:
		s.toggleRects([15]int{1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 0, 0, 1})
		break
	case 5:
		s.toggleRects([15]int{1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0})
		break
	case 6:
		s.toggleRects([15]int{1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1})
		break
	case 7:
		s.toggleRects([15]int{1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1})
		break
	}
}

func (s *score) toggleRects(nums [15]int) {
	for i := 0; i < 15; i++ {
		if nums[i] == 1 {
			s.rs[i].SetFillColor(sf.Color{255, 255, 255, 255})
		} else if nums[i] == 2 {
			s.rs[i].SetFillColor(sf.Color{255, 0, 0, 255})
		} else {
			s.rs[i].SetFillColor(sf.Color{0, 0, 0, 0})
		}
	}
}
