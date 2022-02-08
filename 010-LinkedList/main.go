//Author - Lalith Pallage
//A Linked List example

package main

//Node struct to accept any form of data to store in a node
type Node struct {
	next *Node
	data interface{}
}

type List struct {
	head   *Node
	length int
}

//Insert a node
//InsertNode
func (l *List) InsertNode(data interface{}) {
	//create a new node
	newNode := Node{
		next: l.head,
		data: data,
	}
	//Check if a node exists and assigned a new head
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}

func main() {
	linkedList := List{}
	linkedList.InsertNode("Hello")
	linkedList.InsertNode(10)
}
