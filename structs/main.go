package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v\n", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Kennedy",
		contactInfo: contactInfo{
			email:   "jim123@hotmail.com",
			zipCode: 14256,
		},
	}
	// jim := person{"alex", "kogner", contactInfo{"123@alex.com", 14526}}
	jim.print()
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	//pointerToPerson.firstName = newFirstName
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v,%T", p, p)
}
