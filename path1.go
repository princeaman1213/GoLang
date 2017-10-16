package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type person struct {
	vel float64
	dis float64
	index int
}

var check sync.WaitGroup

func main() {
	check.Add(3)
    var a,b,c person

		arrx:=[]float64{0,185,212,265,273,279,292,367,392,426,434,437,477,491,499}
		arry:=[]float64{0,448,99,146,188,56,243,69,321,324,115,132,364,225,121}    //coordinates of person 1 on route 1

		arrx1:=[]float64{0,123,156,190,223,235,278,302,306,387,404,458,480,499}
		arry1:=[]float64{0,35,63,355,255,144,35,366,465,234,134,234,465,324,211}    //coordinates of person 1 on route 2

		arrx2:=[]float64{0,142,156,178,234,278,333,387,390,406,453,487,499}
		arry2:=[]float64{0,24,57,167,388,299,120,268,444,100,98,167,22,110}
		fmt.Println(time.Now())//coordinates of person 1 on route 3

		go func(){
			a=calc(1,arrx,arry)
		}()
	    go func(){
		     b=calc(2,arrx1,arry1)
	    }()
	    go func(){
		    c=calc(3,arrx2,arry2)
	     }()


	check.Wait()

	if a.vel>=b.vel && a.vel>=c.vel{
		fmt.Printf("Best route is route %d with avg vel : %.2f\n",a.index,a.vel)
	}else if b.vel>=a.vel && b.vel>=c.vel{
		fmt.Printf("Best route is route %d with avg vel : %.2f\n",b.index,b.vel)
	}else{
		fmt.Printf("Best route is route %d with avg vel : %.2f\n",c.index,c.vel)
	}

}

func calc(n int,arrx,arry []float64) person{

	 var p2 person
	for p2.index=1;p2.index<len(arrx);p2.index++{
		p2.dis += math.Sqrt(math.Pow((arrx[p2.index-1] - arrx[p2.index]), 2) + math.Pow((arry[p2.index-1] - arry[p2.index]), 2))
	}

	p2.dis+=arry[p2.index-1]
	p2.vel=p2.dis/float64(p2.index*5)
	fmt.Printf("\ndistance of route %v is          : %.2f metres \naverage velocity on this route is : %.2f m/s\n\n",n,p2.dis,p2.vel)
p2.index=n
	check.Done()

	return p2

}




