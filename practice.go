//Import package
package main

//Import fmt for formatting I/O operations
import "fmt"

//Define a node struct
type Node struct{
	data int
	next *Node
}

//Define the LinkedList structure
type LinkedList struct{
	head *Node
}

//Insert node at the beginning
func (ll *LinkedList) InsertAtBeginning(data int){
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}

//Displaying the list
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

//Adding a node at the end
func (ll *LinkedList) Append(data int){
	newNode := &Node{data: data, next:nil}

	if ll.head == nil{
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func main() {
    // Create a new LinkedList
    ll := LinkedList{}

    // Test InsertAtBeginning
    ll.InsertAtBeginning(10)
    ll.InsertAtBeginning(20)
    ll.InsertAtBeginning(30)

    fmt.Println("After inserting 30, 20, 10 at the beginning:")
    ll.Display()

    // Test Append
    ll.Append(40)
    ll.Append(50)

    fmt.Println("After appending 40 and 50:")
    ll.Display()

    // Test more InsertAtBeginning
    ll.InsertAtBeginning(60)

    fmt.Println("After inserting 60 at the beginning:")
    ll.Display()
}

// go run practice.go