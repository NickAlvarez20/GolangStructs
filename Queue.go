//Design a program that implements a queue for learning purposes
//Queue is a linear data structure that follows FIFO, First in/First out
//Elements are added at one end(enqueue) and removed from the other end(dequeue)
/*Implementatino choices
Ideally it is a good idea to use slice as they are dynamically changing array data structures that can grow or shrink at runtime
*/



//Import packages: main
package main

//Import fmt for formatting I/O operations
import "fmt"

//Define a queue struct that contains a slice of interface{}
type Queue struct{
	items []interface{}
}

//Block 2: Constructor
func NewQueue() *Queue {
	return &Queue{
		items: make([]interface{}, 0),
	}
}

//Block 3: Enqueue Method
func (q *Queue) Enqueue(item interface{}){
	q.items = append(q.items, item)
}

//Block 4: Dequeue Method
func (q *Queue) Dequeue() (interface{}, bool){
	if q.IsEmpty(){
		return nil, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

//Block 5: IsEmpty Method
func (q *Queue) IsEmpty() bool{
	//Check if the length of the items slice is zero
	return len(q.items) == 0
}

//Block 6: Helper Methods-Size Method
func (q *Queue) Size() int{
	return len(q.items)
}

//Testing out the newly created queue methods and queue implementation
func main(){
	//Create a new queue
	queue := NewQueue()
	
	//Check if the queue is empty initially
	fmt.Printf("Is the queue empty? %v\n", queue.IsEmpty())

	//Enqueue some itmes
	queue.Enqueue("First item")
	queue.Enqueue("Second item")
	queue.Enqueue("Third item")

	//Check size and emptiness again
	fmt.Printf("Queue size: %d\n", queue.Size())
	fmt.Printf("Is the queue empty? %v\n", queue.IsEmpty())

	//Dequeue items
	item, success := queue.Dequeue()
	if success{
		fmt.Printf("Dequeued item: %v\n", item)
	} else{
		fmt.Println("Queue was empty when trying to dequeue.")
	}

	//Check the size after dequeue
	fmt.Printf("Queue size after dequeue: %d\n", queue.Size())

	//Dequeue until the queue is empty
	for !queue.IsEmpty(){
		item, success := queue.Dequeue()
		if success{
			fmt.Printf("Dequeued item: %v\n", item)
		}
	}

	//Final check for emptiness
	fmt.Printf("Is the queue empty now? %v\n", queue.IsEmpty())

}

//go run Queue.go