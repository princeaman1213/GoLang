package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	fmt.Println("hey")
	var check sync.WaitGroup

	check.Add(3)
	go func(){
		var i int
		var d float64
		var v float64
		arrx:=[]float64{0,185,212,265,273,279,292,367,392,426,434,437,477,491,499}
		arry:=[]float64{0,78,99,146,188,56,243,69,321,324,115,132,364,225,121}    //coordinates of person 1 on route 1

		for i=1;i<len(arrx);i++{
			d += math.Sqrt(math.Pow((arrx[i-1] - arrx[i]), 2) + math.Pow((arry[i-1] - arry[i]), 2))
		}

		d+=arry[i-1]
		v=d/float64(i*5)
		fmt.Printf("\ndistance of route 1 is          : %.2f metres \naverage velocity on this route is : %.2f m/s",d,v)
		check.Done()
	}()

	go func(){
		var i1 int
		var d1 float64
		var v1 float64
		arrx:=[]float64{0,123,156,190,223,235,278,302,306,387,404,458,480,499}
		arry:=[]float64{0,35,63,355,255,144,35,366,465,234,134,234,465,324,211}    //coordinates of person 1 on route 1

		for i1=1;i1<len(arrx);i1++{
			d1 += math.Sqrt(math.Pow((arrx[i1-1] - arrx[i1]), 2) + math.Pow((arry[i1-1] - arry[i1]), 2))
		}

		d1+=arry[i1-1]
		v1=d1/float64(i1*5)
		fmt.Printf("\ndistance of route 2 is          : %.2f metres \naverage velocity on this route is : %.2f m/s",d1,v1)
		check.Done()
	}()

	go func(){
		var i2 int
		var d2 float64
		var v2 float64
		arrx:=[]float64{0,142,156,178,234,278,333,387,390,397,406,453,487,499}
		arry:=[]float64{0,24,57,167,388,299,120,268,444,222,100,98,167,273,110}    //coordinates of person 1 on route 1

		for i2=1;i2<len(arrx);i2++{
			d2 += math.Sqrt(math.Pow((arrx[i2-1] - arrx[i2]), 2) + math.Pow((arry[i2-1] - arry[i2]), 2))
		}

		d2+=arry[i2-1]
		v2=d2/float64(i2*5)
		fmt.Printf("\ndistance of route 3 is          : %.2f metres \naverage velocity on this route is : %.2f m/s",d2,v2)
		check.Done()
	}()

	check.Wait()


}




