package main

import (
	"fmt"
	"math"
)

type shape interface {
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
	r :=rect{4,5.26}
    fmt.Println(c.area())
	fmt.Println(r.area())
	fmt.Println(totalarea(&c,&r))
}

func (c circle) area() float64{         //using method
	return math.Pi*c.r*c.r
	//fmt.Println(math.Pi*c.r*c.r)
}

func (r rect) area() float64{         //using method
	return r.a*r.b
	//fmt.Println(math.Pi*c.r*c.r)
}

func totalarea(shapes ...shape) float64{       //vaguely understood
	var area float64
	for _, s :=range shapes{
		area+=s.area()
	}
	//area :=a.area()+b.area()
	return area
}



