package main

import "fmt"

type medicine interface {
	mg() float64
}

type paracetamol struct{
	name string
	dose_a_day float64
	mg_1_tab float64
}

type digene struct{
	name string
	dose_a_day float64
	mg_1_tab float64
}

type amphotericin struct{
	name string
	dose_a_day float64
	mg_1_tab float64
}

func main() {
	p :=paracetamol{"Paracetamol",4,200}
	d :=digene{"Digine",3,400}
	a :=amphotericin{"Amphotericin-B",5,187}

	fmt.Println("Total mg of Paracetamol a day :",p.mg(),"mg")
	fmt.Println("Total mg of Digene a day :",d.mg(),"mg")
	fmt.Println("Total mg of Amphotericin-B a day :",a.mg(),"mg")
	fmt.Println("Let The safe limit of mg is 1000 mg a day \nSo safe medicines are :")
	//apt_mg()


	func (medi ...medicine) {
	for _, s :=range medi{
	if s.mg()<1000{
	fmt.Println(s)
	}
	}
	}(&p,&d,&a)
}

func (p paracetamol) mg() float64{
	return p.dose_a_day*p.mg_1_tab
}

func (d digene) mg() float64{
	return d.dose_a_day*d.mg_1_tab
}

func (a amphotericin) mg() float64{
	return a.dose_a_day*a.mg_1_tab
}




