package main

import (
	"image"

	"image/color"
	"os"
	"image/png"

)
var img = image.NewRGBA(image.Rect(0,0,500,500))     //area of figure
var col color.Color

func main() {
	col = color.RGBA{0, 0, 255, 255} // Red
	//VLine(10, 20, 80)
	//HLine(10, 20, 80)
	col = color.RGBA{255, 0, 0, 255} // Green
	//Rect(0,0,100,100)

	col = color.RGBA{255, 0, 0, 255}
	circle(0,0,300)

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f,img)  //To write the image in file
}

// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}


func circle(x,y,r int){
	x=x+r
	y=0

	for x>=0{
		px := (x+(x-1))/2
		py := y+1
		if (px*px + py*py)<(r*r){
			img.Set(x,y+1, col)

		}else {
			img.Set(x-1,y+1, col)
			x--
		}
		y++
		if x-y<1{
			break
		}

	}

	y=x
	x=0

	for y>=0{
		px := x+1
		py := (y+(y+1))/2
		if (px*px + py*py)<(r*r){
			img.Set(x+1,y, col)

		}else {
			img.Set(x+1,y+1, col)
			y--
		}
		x++
		if y==x{
			break
		}

	}


}
