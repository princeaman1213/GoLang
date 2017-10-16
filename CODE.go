package main

import "fmt"

func main() {
//	a := 'a'     //rune ?
//	z := 'z'

//	fmt.Println(int(a),"   ",int(z))  //print ascii

//	var b int=68
//	fmt.Println(string(b),"\n")  //print character from ascii

//	q :=[5]int{68,69,70,71,72}
//	for i:=0;i<=4;i++{
//		fmt.Println(string(q[i]))
//	}


	var input [5]rune


	for i:=0;i<5;i++{
		fmt.Scanf("%c",&input[i])
	}

	//encoding
	fmt.Println("Encoding.....")

	fmt.Println("ascii of input before conversion :",input)
	for j:=0;j<5;j++{
		//input[j]=input[j]+(input[j]-96)
		input[j]=input[j]+1
		input[j]=input[j]-40
	}
	fmt.Println("ascii of input after conversion :",input)

	//encoded string
	fmt.Println("\nThe encoded string is :")
	for i:=0;i<5;i++{
		fmt.Printf("%c",input[i])
	}

	//decoding
	for k:=0;k<5;k++{
		//input[k]=input[k]-1
		input[k]=input[k]+39
	}

	fmt.Println("\nThe decoded string is :")
	for i:=0;i<5;i++{
		fmt.Printf("%c",input[i])
	}







}
