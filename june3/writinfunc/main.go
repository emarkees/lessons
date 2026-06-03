package main

import (
	"fmt"
)

func paintNeeded(width, height float64) float64 {
	area := width*height

	return area/10.0
}

func main(){
	fmt.Printf("%.2f\n",paintNeeded(4.2, 3.0))
}