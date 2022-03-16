package main

import (
	"fmt"
	"math"
)

type Shape interface {
	perimeter() float64
	area() float64
}
type Point struct {
	x, y float64
}

func distance(point1, point2 Point) float64 {
	deltaX := point2.x - point1.x
	deltaY := point2.y - point1.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

type Rectangle struct {
	leftLowVertex, rigthUpVertex Point
}

func (r *Rectangle) perimeter() float64 {
	leftLowVertex := Point{r.leftLowVertex.x, r.leftLowVertex.y}
	leftUpVertex := Point{r.leftLowVertex.x, r.rigthUpVertex.y}
	rightUpVertex := Point{r.rigthUpVertex.x, r.rigthUpVertex.y}
	return 2*distance(leftLowVertex, leftUpVertex) + 2*distance(leftUpVertex, rightUpVertex)
}
func (r *Rectangle) area() float64 {
	leftLowVertex := Point{r.leftLowVertex.x, r.leftLowVertex.y}
	leftUpVertex := Point{r.leftLowVertex.x, r.rigthUpVertex.y}
	rightUpVertex := Point{r.rigthUpVertex.x, r.rigthUpVertex.y}
	return distance(leftLowVertex, leftUpVertex) * distance(leftUpVertex, rightUpVertex)
}

type Circle struct {
	center Point
	radius float64
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, v := range shapes {
		perimeter += v.perimeter()
	}
	return perimeter
}
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, v := range shapes {
		area += v.area()
	}
	return area
}
func main() {
	r := Rectangle{Point{0, 0}, Point{5, 10}}
	c := Circle{Point{0, 0}, 3}
	fmt.Println("rectangle perimeter = ", r.perimeter())
	fmt.Println("circle perimeter = ", c.perimeter())
	fmt.Println("total perimeter = ", totalPerimeter(&r, &c))
	fmt.Println("rectangle area = ", r.area())
	fmt.Println("circle area = ", c.area())
	fmt.Println("total area = ", totalArea(&r, &c))
}
