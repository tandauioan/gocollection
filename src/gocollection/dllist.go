package gocollection

import collection "gocollection"

type DLList struct {
	collection.Collection

	head, tail *element_DLList
	count int
}

type element_DLList struct {
	value interface{}
	prev, next *element_DLList
}


func NewDLList() (list *DLList) {
	list=new(DLList)
	list.head=new(element_DLList)
	list.tail=new(element_DLList)
	list.head.prev=nil
	list.head.next=list.tail
	list.tail.prev=list.head
	list.tail.next=nil
	list.head.value=nil
	list.tail.value=nil
	list.count=0
	return
}

func (self *DLList) Add(object interface{}) bool {
	node:=new(element_DLList)
	node.value=object
	node.next=self.tail
	node.prev=self.tail.prev
	self.tail.prev=node.prev
	return true
}

func (self *DLList) AddAll(c collection.Collection) bool {
	it:=c.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			self.Add(value)
		}
	}
	return true
}

func (self *DLList) Clear() {
	self.head.next=self.tail
	self.tail.prev=self.head
}






