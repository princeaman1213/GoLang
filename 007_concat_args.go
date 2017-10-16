package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main() {
	name :=os.Args[1]             //runs in terminal only as args accepts command line inputs
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	template :=fmt.Sprint(`
	    <!DOCTYPE html>
	    <html lang="eu">
	    <head>
	    <meta charset="UTF-8">
	    </head>
	    <body>
	    <h1>` +
		name  +`
        </h1>
        </body>
        </html>
	 `)

	f,err := os.Create("templateargs.html")
	if err !=nil{
		log.Fatal("error creating file :",err)
	}
	defer f.Close()
	io.Copy(f,strings.NewReader(template))


}
