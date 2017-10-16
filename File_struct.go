package main
import (
	"fmt"
	"math"
)

type circle struct{
	x,y,r float64
}
func main() {

	var c circle
	//c.x=0
//	c.y=0
	//c.r=4
	//fmt.Println("Area is:",3.14*c.r*c.r)
	//c :=circle{0,0,5}
	fmt.Println(area(c))
}

func area(c circle) float64{
	return math.Pi *c.r*c.r
}
