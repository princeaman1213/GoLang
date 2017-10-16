package main
import(
	"fmt"
	"os"
)
func main() {
	file,err :=os.Open("name.txt")
	if err !=nil{
		return
	}
	defer file.Close()

	stat,err :=file.Stat()           // get file stats
	if err!=nil{
		return
	}
	fmt.Println(stat.Name())         //filename
	fmt.Println(stat.Mode())         // r -read , w -write
	//fmt.Println(stat.Size())       // no. of characters

	bs :=make([]byte,stat.Size())
	q,err :=file.Read(bs)
	if err!=nil{
		return
	}
	fmt.Println("q is :",q)
	str :=string(bs)       //converting bs which contains all ascii values to string i.e characters
	fmt.Println(str)

}
