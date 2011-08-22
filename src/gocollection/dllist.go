package gocollection


type DLList struct {
	Collection
	Stack
	Deque

	head, tail *element_DLList
	size int
}

type element_DLList struct {
	value interface{}
	prev, next *element_DLList
}


type dllist_forward_Iterator struct {
	Iterator

	list *DLList
	node *element_DLList
}

type dllist_backward_Iterator struct {
	Iterator

	list *DLList
	node *element_DLList
}

func NewDLList() *DLList {
	list:=new(DLList)
	list.size=0
	list.head=new(element_DLList)
	list.head.value=nil
	list.head.prev=nil
	list.tail=new(element_DLList)
	list.tail.value=nil
	list.tail.next=nil
	list.head.next=list.tail
	list.tail.prev=list.head
	return list
}

func getForwardIterator(list *DLList) Iterator {
	it:=new(dllist_forward_Iterator)
	it.list=list
	it.node=list.head
	return it
}

func getBackwardIterator(list *DLList) Iterator {
	it:=new(dllist_backward_Iterator)
	it.list=list
	it.node=list.tail
	return it
}

func (self *dllist_forward_Iterator) Next() (value interface{}, valid bool) {
	self.node=self.node.next
	if self.node==self.list.tail {
		value=nil
		valid=false
	} else {
		value=self.node.value
		valid=true
	}
	return
}

func (self *dllist_forward_Iterator) Remove() {
	self.node.prev.next=self.node.next
	self.node.next.prev=self.node.prev
	self.node=self.node.prev
	self.list.size--
}

func (self *dllist_backward_Iterator) Next() (value interface{}, valid bool) {
	self.node=self.node.prev
	if self.node==self.list.head {
		value=nil
		valid=false
	} else {
		value=self.node.value
		valid=true
	}
	return
}

func (self *dllist_backward_Iterator) Remove() {
	self.node.prev.next=self.node.next
	self.node.next.prev=self.node.prev
	self.node=self.node.next
	self.list.size--
}

func (self *DLList) Iterator() Iterator {
	return getForwardIterator(self)
}

func (self *DLList) ReverseIterator() Iterator {
	return getBackwardIterator(self)
}

func (self *DLList) Add(object interface{})bool {
	node:=new(element_DLList)
	node.next=self.tail
	node.prev=self.tail.prev
	self.tail.prev.next=node
	self.tail.prev=node
	self.size++
	return true
}

func (self *DLList) AddAll(c Collection) bool {
	it:=c.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			if !self.Add(value) {
				return false
			}
		}
	}
	return true
}

func (self *DLList) Contains(object interface{}, equal Equal) bool {
	it:=self.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			if equal(value, object) {
				return true
			}
		}
	}
	return false
}

func (self *DLList) ContainsAll(c Collection, equal Equal) bool {
	it:=c.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			if !self.Contains(value,equal) {
				return false
			}
		}
	}
	return true
}

func (self *DLList) Clear() {
	self.head.next=self.tail
	self.tail.prev=self.head
	self.size=0
}

func (self *DLList) IsEmpty() bool {
	return self.size==0
}

func (self *DLList) Remove(object interface{}, equal Equal) {
	it:=self.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break;
		} else {
			if equal(value,object) {
				it.Remove()
			} 
		}
	}
}

func (self *DLList) RemoveAll(c Collection, equal Equal) {
	it:=c.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			self.Remove(value,equal)
		}
	}
}

func (self *DLList) RetainAll(c Collection, equal Equal) {
	it:=self.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			if !c.Contains(value,equal) {
				it.Remove()
			}
		}
	}
}

func (self *DLList) Size() int {
	return self.size
}

func (self *DLList) Peek() interface{} {
	return self.head.next.value
}

func (self *DLList) Pop() interface{} {
	value:=self.head.next.value
	self.head.next.next.prev=self.head
	self.head.next=self.head.next.next
	self.size--
	return value
}

func (self *DLList) add_first(element interface{}) bool {
	node:=new(element_DLList)
	node.prev=self.head
	node.next=self.head.next
	self.head.next=node
	node.next.prev=node
	node.value=element
	self.size++
	return true
}

func (self *DLList) Push(element interface{}) bool {
	return self.add_first(element)
}

func (self *DLList) Search(element interface{}, equal Equal) int {
	it:=self.Iterator()
	pos:=0
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			if equal(element,value) {
				return pos
			}
			pos++
		}
	}
	return -1
}

func (self *DLList) AddTail(element interface{}) bool {
	node:=new(element_DLList)
	node.next=self.tail
	node.prev=self.tail.prev
	node.prev.next=node
	self.tail.prev=node
	node.value=element
	self.size++
	return true
}

func (self *DLList) PeekHead() interface{} {
	return self.Peek()
}

func (self *DLList) PopHead() interface{} {
	return self.Pop()
}

func (self *DLList) AddHead(element interface{}) bool {
	return self.add_first(element)
}

func (self *DLList) PeekTail() interface{} {
	return self.tail.prev.value
}

func (self *DLList) PopTail() interface{} {
	value:=self.tail.prev.value
	self.tail.prev.prev.next=self.tail
	self.tail.prev=self.tail.prev.prev
	self.size--
	return value
}













