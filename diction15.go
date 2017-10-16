package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Body.Close()
	a, err1 := ioutil.ReadAll(f.Body)
	if err != nil {
		fmt.Println(err1)
	}

	slice := []string{}
	slice = append(slice, string(a))  //fmt.Println(len(slice))
	slice = strings.Split(slice[0], "\n")

	var ip, ch, ip1 string
    var index int
	fmt.Println("Enter some(1 or more) letters of your word :")
	fmt.Scanf("%v", &ip)

	for {
		for i, v := range slice {

			if strings.HasPrefix(v, ip) {
				fmt.Println(i, " : ", v)
			}
		}
		fmt.Println("Want to enter further letters ?(y/n)")
		fmt.Scanf("%s", &ch)
		if ch == "y" {
			fmt.Printf(ip)
			fmt.Scanf("%v", &ip1)
			ip = ip + ip1
			//fmt.Println(ip)
		}else{
			break
		}
	}
	fmt.Println("Enter the index of selectes word :")
	fmt.Scanf("%v",&index)
	fmt.Println(slice[index])
}