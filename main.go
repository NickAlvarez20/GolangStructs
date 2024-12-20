//implement a few methods that will implement a doubly linked list implementation
//has to have the right time complexities
//circularly linked list

package main

type Node[T any] struct {
	//Node struct
	next *Node[T]
	prev *Node[T]
	value T
}

type LinkedList[T any] struct{
	//Hold the head node and length of the list
	//Note: These are different because they are generics so be careful with printing and testing 
	head *Node[T]
	length uint
}

func (11 *LinkedList[T]) index(idx uint) (T, bool){
	//return value and "ok", indicating if index is valid
	//returns the index of whatever value is within the linked list 
}

func (11 *LinkedList[T]) append(value T) {
	//add an element to the end of the linked list

}

func (11 *LinkedList[T]) pop() T {
	//removing and returning the last element of the linked list
	//edge case = nil or return value for when linked list in empty
}

func (11 *LinkedList[T]) printList() {
	//prints all the values in the doubly linked list
	//idea is how to iterate through the linked list until you reach the end of it

}

func main(){

}