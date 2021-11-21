package main

import "fmt"
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}


func main()  {
	jim := person{
		firstName: "Jim", 
		lastName: "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//Turn value in to address with &value
	jimPointer := &jim // &jim give us the memory address of the value this variable is pointing at
	jimPointer.updateFirstName("Jimmy")
	jim.print()
}

//*pointer give me the value this memmory address is pointing at
//*person is a type dscription - it means we're working with a pointer to a person
//Turn address in to value with *address
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	//This is an operator - it means we want to manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}