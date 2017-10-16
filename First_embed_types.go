package main
import "fmt"

type person struct{
	name string
}

type android struct{
	person person                  //android is a person
	model string
}

func main() {

	a :=new(person)
	b :=new(android)
	//fmt.Println(a.talk())
	a.talk()                                 //accessing the method talk() associated with person
	b.person.talk()                          //accessing method talk() associated with person again but this time coz android is-a person
}

func (p *person) talk(){                     //associating a method talk() with struct and perform a particular task
	fmt.Println("Hi all",p.name)
}
