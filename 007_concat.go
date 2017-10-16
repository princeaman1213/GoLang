package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main() {
	name :="Aman Patel"

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

	f,err := os.Create("template.html")
	if err !=nil{
		log.Fatal("error creating file :",err)
	}
	defer f.Close()
	io.Copy(f,strings.NewReader(template))


}
