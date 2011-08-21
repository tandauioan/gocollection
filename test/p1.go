package main

import "gocollection"
import "golists"
import "fmt"

func equalInt (o1 interface{}, o2 interface{}) bool{
	return o1==o2
}

func printIteratorElements(it gocollection.Iterator) {
	i:=0
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			fmt.Printf("E%v: %v\n",i,value)
			i++
		}
	}
}

func main() {

	list:=golists.NewDLList()
	for i:=0;i<5;i++ {
		list.AddTail(i)
	}
	fmt.Printf("Size:%v\n",list.Size())
	fmt.Printf("---\n")
	printIteratorElements(list.Iterator())
	fmt.Printf("----\n")
	printIteratorElements(list.ReverseIterator())
	fmt.Printf("%v\n",list)

}


