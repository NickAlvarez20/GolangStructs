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

//Create main func to run executable
//Create an instance of Person struct
func main(){
p1 := Person{name: "Elon Tusk", age: 53}
fmt.Printf("%s is %d years old.\n",p1.name,p1.age)
p1.IsAdult()
}

//Generics

//Constructors

//Instances

//go run Heap.go

//made changes
