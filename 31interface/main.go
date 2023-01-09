package main

import (
	"fmt"
	"math"
)

// type animal interface {
// 	Speak() string
// }

// type dog struct {
// }

// type cat struct {
// }
// type human struct {
// }

// func (d dog) Speak() string {
// 	return "Bark!"
// }
// func (c cat) Speak() string {
// 	return "Meow!"
// }
// func (h human) Speak() string {
// 	return "It is a good practice!"
// }

// func main() {
// 	being := []animal{dog{}, cat{}, human{}}
// 	for _, v := range being {
// 		fmt.Println(v.Speak())
// 	}
// }

type shapes interface {
	perimeter() float64
	area() float64
}

type circle struct{
	radius float64
}

type rectangle struct{
	length,breadth float64
}

type square struct{
	side float64
}

func (c circle) area() float64{
	fmt.Println("Area of circle:")
	return math.Pi*c.radius*c.radius
}

func (c circle) perimeter() float64{
	fmt.Println("Perimeter of circle:")
	return 2*math.Pi*c.radius
}

func (r rectangle) area() float64{
	fmt.Println("Area of rectangle:")
	return r.length*r.breadth
}
func (r rectangle) perimeter() float64{
	fmt.Println("Perimeter of rectangle:")
	return 2*r.length+2*r.breadth
}
func (s square) area() float64{
	fmt.Println("Area of square:")
	return s.side*s.side
}
func (s square) perimeter() float64{
	fmt.Println("Perimeter of square:")
	return 4*s.side
}
func measurement(s shapes){
//	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main(){
	geometry:=[]shapes{rectangle{length: 2,breadth: 4},circle{radius: 4},square{side: 10}}
	for _, v := range geometry {
		measurement(v)
	}
// 	r:=rectangle{length: 4,breadth: 4}
// 	c:=circle{radius: 4}
// 	s:=square{side: 4}
// measurement(r)
// measurement(c)
// measurement(s)
}