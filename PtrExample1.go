package main

// This will Not Work as we are passing Rectangle as of type shape to a method which expects a pointer to Rectangle
type square struct {
	width  int
	length int
}

type shape interface {
	area() int
}

func (s *square) area() int {
	return s.width * s.length
}

func info(s shape) int {
	return s.area()
}

func main() {
	s := square{5, 5}
	info(s)
}
