//Golang Learning Exercises
//Import Packages: Main
package main

//Import fmt for formatting I/O operations
import "fmt"

//Structs

//Define a person struct with name and age
type Person struct{
	name string
	age int
}

//Extend person to include method IsAdult that returns whether the persons age is 18 or older
func (p Person) IsAdult() bool{
	//create instance
	isAdult := p.age >= 18
	if isAdult {
		fmt.Printf("Yes %s is older than 18\n",p.name)
	} else {
		fmt.Printf("No, %s is not older than 18\n",p.name)
	}
	return isAdult
}

//Nested Structs-Create a Family struct with a slice of Person. Implement a method to find the oldest family member
type Family struct{
	Members []Person
}

func (f Family) OldestMember() Person {
	var oldest Person
	for _, member := range f.Members{
		if member.age > oldest.age{
			oldest = member
		}
	}
	return oldest
}

//Create main func to run executable
//Create an instance of Person struct
func main(){
	//Instantiate a few people
p1 := Person{name: "Elon Tusk", age: 53}
p2 := Person{name: "Grok Musk", age: 2500}
p3 := Person{name: "Dusk Musk", age: 22}
//Print information about one person
fmt.Printf("%s is %d years old.\n",p1.name,p1.age)
fmt.Printf("%s is an adult: %v\n",p1.name, p1.IsAdult())
//Create an instance of Family with members
familyGood := Family{
	Members: []Person{p1,p2,p3},
}
//Find and print the oldest family member
oldest := familyGood.OldestMember()
fmt.Printf("The oldest person in the family is %s, aged %d\n",oldest.name,oldest.age)
}

//Generics

//Constructors

//Instances

//go run Heap.go

//made changes
