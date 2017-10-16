package main
import(
	"fmt"
	"math"
)

type circle struct{
	x,y,r float64
}
func main() {
	c :=circle{0,0,4}
	fmt.Println(c.area())
	//fmt.Println(area(&c))
	//c.area()                               //accessing method area associated with circle

}

func (c *circle) area() float64{         //using method
	return math.Pi*c.r*c.r
	//fmt.Println(math.Pi*c.r*c.r)
}

/*
func area(c *circle) float64{           //using function
	return math.Pi*c.r*c.r
}
*/