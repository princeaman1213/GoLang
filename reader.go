package main

import (
	"net/http"
	//"io/ioutil"
	"fmt"
	"bufio"
)

func main() {
	res,_ :=http.Get("https://www.golang-book.com/books/intro/13#section3")

	//bs,_:=ioutil.ReadAll(res.Body)
	//fmt.Println(string(bs))

           //OR

	sc :=bufio.NewScanner(res.Body)
	for sc.Scan(){
		fmt.Println(sc.Text())
	}

}
