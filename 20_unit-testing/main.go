package main

func RectangleArea(length, width int) string {
	area := length * width
	if area%2 == 0 {
		return "even rectangle"
	} else {
		return "odd rectangle"
	}
}

func RectanglePerimeter(length, width int) bool {
	perimeter := 2 * (length + width)
	if perimeter%5 == 0 {
		return true
	} else {
		return false
	}
}

func SquareArea(s int) string {
	area := s * s
	if area%2 == 0 {
		return "even square"
	} else {
		return "odd square"
	}
}

func SquarePerimeter(s int) bool {
	perimeter := 4 * s
	if perimeter >= 40 {
		return true
	} else {
		return false
	}
}
