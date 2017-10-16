package main

import "fmt"

type person struct{
	fname string
	lname string
	favfood []string
}

type transportation interface{
	transportationDevice() string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
    vehicle
	luxury bool
}

type gator int
type flamingo bool

type swamp interface {
	greeting()
}

func main() {
	sl :=[]int{3,4,2,1}
	for i,val :=range sl{
		fmt.Println(val,i)
	}

	fmt.Println("\n \n")

	m :=map[string]int{
		"Dhoni": 1,
		"Kohli": 2,
		"Rahane": 3,
		"Sharma": 4,
		"Pandya": 5,
	}
	for j,val :=range m{
		fmt.Println(val,j)
	}

	fmt.Println("\n \n")

	p1 :=person{"Aman","Patel",[]string{"Pizza","Burger"}}
	fmt.Println(p1.favfood)
	for j,val :=range p1.favfood{
		fmt.Println(val,j)
	}
	a:=fmt.Sprintln(p1.walk(),"is walking")
	fmt.Println(a,"\n \n")


	t :=truck{vehicle{2,"Red"},true}
	s :=sedan{vehicle{4,"White"},true}
	fmt.Println(t,s)
	fmt.Println(t.color,s.luxury)

	fmt.Println("\n \n")

	fmt.Println(t.transportationDevice(),"\n",s.transportationDevice())

	fmt.Println("\n \n")

	report(t,s)

	fmt.Println("\n \n")

	var new gator=3
	var fl flamingo
	var new1 int=4
	fmt.Println(new,new1)
	fmt.Printf("%T %T",new,new1)

	fmt.Println("\n \n")
	//new1=new         //cant be done
	new1=int(new)
	fmt.Println(new1)  //typecasting

	fmt.Println("\n \n")

	new.greeting()

	fmt.Println("\n \n")

	bayou(new,fl)

	fmt.Println("\n \n")

	s1:="I am sorry , I cant do that"
	fmt.Println(s1)
	fmt.Println([]byte(s1))
	fmt.Println(string([]byte(s1)))
	fmt.Println(string([]byte(s1)[:10]))
	qq:=[]byte(s1)
	for _,v:=range  qq{
		fmt.Println(string(v))
	}





}

func (p person) walk() string{
	return p.fname
}

func(t truck) transportationDevice() string{
	a :="A truck carries cargo"
	return a
}

func(s sedan) transportationDevice() string{
    a :="A sedan is for richies"
	return a
}

func(g gator) greeting(){
	fmt.Println("I am a gator")
}

func(f flamingo) greeting(){
	fmt.Println("I am a flamingo")
}

func bayou(cre ...swamp){
	for _,v:=range cre{
		v.greeting()
	}
}

func report(tport ...transportation) {
	for _,v:=range tport{
		fmt.Println(v.transportationDevice())
	}

}