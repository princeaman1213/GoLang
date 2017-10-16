package main

import (
	"net/http"
	//"io/ioutil"
	"fmt"
	"bufio"
)

func main() {

	res,_ :=http.Get("https://currentaffairs.gktoday.in/")

	//bs,_ :=ioutil.ReadAll(res.Body)
	//fmt.Println(string(bs))

	sc :=bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)
	w :=make(map[string]string)
	for sc.Scan(){
		w[sc.Text()]="rrrrrr"

	}

	for k,_ :=range w{
		fmt.Println(k)
	}


}
