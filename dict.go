package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
   // var i int
	f, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words_alpha.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Body.Close()
	a, err1 := ioutil.ReadAll(f.Body)
	if err != nil {
		fmt.Println(err1)
	}

	var slice =make([]string,1)

	//slice[0]="shjshks"
	//slice=append(slice, "ddfsf")
	//slice[1]="dhsd"
	//fmt.Println(slice,len(slice))

	aa:=string(a)
	//aa=strings.TrimSpace(aa)
	fmt.Println(len(aa))

	//fmt.Println(aa[0],aa[1])

	for i:=0;i<len(aa)-1;i++ {
		if aa[i]==10{
			//fmt.Printf("\n br :  %v",aa[i])
			slice=append(slice,string(aa[i]))
		}
		if aa[i]==13{
			continue
		}else{
			//fmt.Printf("%v",string(aa[i]))
			slice=append(slice,string(aa[i]))
		}

		/*if aa[i]<97 || aa[i]>122{
			if aa[i] != 10{
				fmt.Println("vcxvgvfdg",a[i],string(aa[i]))
			}

		}*/

	}
//fmt.Println(len(slice))
for j:=0;j<20;j++{
	fmt.Println(j," --",slice[j])
}



}




/*
for i:=0;i<len(aa);i++ {
		if aa[i]==13{
			continue
		}else if aa[i]==10{
			fmt.Println("")
			continue
		}else{
			fmt.Printf("%v",string(aa[i]))
		}

	}
 */