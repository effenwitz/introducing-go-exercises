package main

import (
	"fmt"
	"math"
)

//a struct is like a domina object of java. it is type that contains named fields

//circle
type Circle struct {
	x float64
	y float64
	r float64
}

//rectangle
type Rectangle struct {
	l, b float64
}

func main() {
	fmt.Println("Structs and Interfaces")

	// var c Circle
	// c.x, c.y, c.r = 1, 2, 3
	// // d := new(Circle)
	// e := Circle{x: 0, y: 0, r: 1}
	// fmt.Println(e.x, e.y, e.r)

	// fmt.Println(circleArea(c))
	// fmt.Println(circleAreaMod(&c))
	// fmt.Println(c.r)
	// fmt.Println(c.area())

	// rect := Rectangle{l: 1, b: 2}
	// fmt.Println(rect.area())

	per := Person{Name: "Kirty"}
	per.printName()
	a := new(Android)
	a.Person.printName()
	calcAre()
}

//does not modify the original circle radius
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

//for modification write with pointer
func circleAreaMod(c *Circle) float64 {
	return math.Pi * float64(c.r*c.r)
}

//Methods
func (c Circle) area() float64 {
	fmt.Println("Cirlce area")
	return math.Pi * float64(c.r*c.r)
}

func (r Rectangle) area() float64 {
	fmt.Println("Rectangle area")
	return r.l * r.b
}

//Embedded Types
//struct field usually represents a has-a relationship.
//Person is a representation of person.
type Person struct {
	Name string
}

func (p Person) printName() {
	fmt.Println("name is ", p.Name)
}

//Profile is a representation of Person
type Profile struct {
	person  Person
	address string
}

//Android is a person
type Android struct {
	Person
	houseNum uint32
}

//Interfaces in Go

//interfaces are used to make similar methods explicit
//no fields but method set defined in interface
type Shape interface {
	area() float64
}

func totalArea(cir ...Circle) float64 {
	var total float64
	for _, v := range cir {
		total += v.area()
	}
	return total
}

func calcAre() {
	var c Circle
	c.x, c.y, c.r = 1, 2, 3
	var r Rectangle
	r.l, r.b = 1, 2
	//notice the way to call interface parameterized function
	fmt.Println(totalAr(&c, &r))
}

//to add all the areas of geometricla shapes, we can use inteface
func totalAr(shapes ...Shape) float64 {
	var area float64
	for _, v := range shapes {
		area += v.area()
	}
	return area
}

//interface can also be used as fields
//Multishape is made of several smaller shapes
type Multishape struct {
	shapes []Shape
}

//area method shoild not have pointer
//reciever
// func testMultiShape() {
// 	multishapes := Multishape{
// 		shapes: []Shape{
// 			Circle{0, 0, 5},
// 			Rectangle{1, 2},
// 		},
// 	}
// }

//turn multishape to shape
// func (m *Multishape) area() float64 {
// 	var totArea float64
// 	for _, v := range m.shapes {
// 		totalAr += v.area()
// 	}
// 	return totalAr
// }
