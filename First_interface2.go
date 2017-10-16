package main
import(
	"fmt"
	"math"
)

type shape interface{
	area() float64
}

type circle struct{
    x,y,r float64
}

type rect struct{
	a,b float64
}

func main() {
	c :=circle{0,0,4}
	r :=rect{5,5.26}
	//fmt.Println(c.area())
	//fmt.Println(r.area())
	fmt.Println(areaofshape(c))
	fmt.Println(areaofshape(r))

}

func (c circle) area() float64{
	return math.Pi*c.r*c.r
}

func (r rect) area() float64{
	return r.a*r.b
}

func areaofshape(s shape) float64{
	return s.area()
}

                    //OR
/*
func areaofshape(s interface{area() float64}) float64{
	return s.area()
}
*/