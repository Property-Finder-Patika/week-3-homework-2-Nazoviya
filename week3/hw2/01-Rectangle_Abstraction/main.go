package main

import (
	"fmt"
	"os"
)

// create rectangle struct
type Rectangle struct {
	width  float64
	height float64
}

func main() {
	r := createR(2.7, 6.08)
	printer(r)
}

// create actual rectangle, return error if not expected values are provided.
func createR(r1 float64, r2 float64) (r Rectangle) {
	r = Rectangle{r1, r2}
	if r.width < 0 || r.height < 0 {
		err := fmt.Errorf("Error! Positive values are expected, found %v and %v", r1, r2)
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return r
}

// calculate circumference of the rectangle
func circumference(r Rectangle) float64 {
	result := r.width*2 + r.height*2
	return result

}

// calculate area of the rectangle
func area(r Rectangle) float64 {
	result := r.width * r.height
	return result
}

// print rectangle properties and expected values.
func printer(r Rectangle) {
	// print area and circumference
	fmt.Printf("Rectangle is: %.2f by %.2f\n", r.width, r.height)
	fmt.Printf("Area of the rectangle is: %.2f\n", area(r))
	fmt.Printf("Circumference of the rectangle is: %.2f\n", circumference(r))
}
