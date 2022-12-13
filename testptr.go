package main

import (
	"fmt"
)

type Square struct {
	length int
}
type Shape interface {
	Area() int
}

// receiver(t T)
//
//	func (sq Square) Area() int {
//		return sq.length * sq.length
//	}
func describe(s Shape) {
	fmt.Println("Area", s.Area())
}

// func main() {
// sq := Square{
// 	length: 5,
// }
// 	describe(sq) // value `sq` (T)
// }

// receiver(t *T)
func (sq *Square) Area() int {
	return sq.length * sq.length
}

func main() {
	sq := Square{
		length: 5,
	}
	describe(sq) // value sq (T)
}

// // receiver(t *T)
// func (sq *Square) Area() int {
//     return sq.length * sq.length
// }

// func main() {
//     describe(&sq)// value sq (*T)
// }

// // receiver(t *T)
// func (sq *Square) Area() int {
//     return sq.length * sq.length
// }

// func main() {
//     describe(&sq)// value sq (T)
// }

// func main() {
//     sq := Square{
//         length: 5,
//     }
//     describe(sq)// value `sq` (T)
// }
