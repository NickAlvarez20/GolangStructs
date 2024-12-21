//Import package
package main

//Import fmt for formatting I/O operations
import "fmt"

//Create Christmas Tree Struct
type ChristmasTree struct{
	height int
	name string
}

//Create Collection for Trees | Utilize slice for data structure to make it dynamically change
type ChristmasTreeCollection struct{
	Trees []ChristmasTree
}

//Add function | AddTree | Adds Trees to the ChristmasTreeCollection
func (c *ChristmasTreeCollection) AddTree(tree ChristmasTree){
	c.Trees := append(c.Trees, ChristmasTreeCollection)
}

//Test Add Tree Function

//go run Structs.go