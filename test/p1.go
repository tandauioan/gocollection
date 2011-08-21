package main

import "golists"
import "fmt"

func equalInt (o1 interface{}, o2 interface{}) bool{
	return o1==o2
}

func main() {

	list:=golists.NewDLList()
	for i:=0;i<5;i++ {
		list.AddTail(i)
	}
	fmt.Printf("Size:%v\n",list.Size())
	it:=list.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			fmt.Printf("E: %v\n",value)
		}
	}

	fmt.Printf("%v\n",list)

}


