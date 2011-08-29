package gollection

/*
* A doubly linked list structure. It maintains the size of the list
* as opposed to the SLList structure so it does not need to recalculate
* it every time.
*/
type DLList struct {
	Collection
	Stack
	Deque

	/* the first (hear) and last (tail) nodes in the list
	* (preallocated) */
	head, tail *element_DLList
	/* the size of the list */
	size int
}

/* a node in the list structure */
type element_DLList struct {
	value interface{}
	prev, next *element_DLList
}

/* DLList forward iterator */
type dllist_forward_Iterator struct {
	Iterator

	list *DLList
	node *element_DLList
}

/* DLList backward iterator */
type dllist_backward_Iterator struct {
	Iterator

	list *DLList
	node *element_DLList
}

/* Creates and returns an empty DLList */
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

/* Returns a forward iterator over the list */
func getForwardIterator(list *DLList) Iterator {
	it:=new(dllist_forward_Iterator)
	it.list=list
	it.node=list.head
	return it
}

/* returns a backward iterator over the list */
func getBackwardIterator(list *DLList) Iterator {
	it:=new(dllist_backward_Iterator)
	it.list=list
	it.node=list.tail
	return it
}

/* Forward iterator Next() function */
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

/* Forward iterator Remove() function */
func (self *dllist_forward_Iterator) Remove() {
	self.node.prev.next=self.node.next
	self.node.next.prev=self.node.prev
	self.node=self.node.prev
	self.list.size--
}

/* Backward iterator Next() function */
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

/* Backward iterator Remove() function */
func (self *dllist_backward_Iterator) Remove() {
	self.node.prev.next=self.node.next
	self.node.next.prev=self.node.prev
	self.node=self.node.next
	self.list.size--
}

/* Returns a forward iterator over the list */
func (self *DLList) Iterator() Iterator {
	return getForwardIterator(self)
}

/* Returns a backward (reverse) iterator over the list */
func (self *DLList) ReverseIterator() Iterator {
	return getBackwardIterator(self)
}

/* Adds a new element to the tail of the list */
func (self *DLList) Add(object interface{})bool {
	node:=new(element_DLList)
	node.value=object
	node.next=self.tail
	node.prev=self.tail.prev
	self.tail.prev.next=node
	self.tail.prev=node
	self.size++
	return true
}

/* Adds all the elements of the collection at the tail of the list */
func (self *DLList) AddAll(c Collection) bool {
	it:=c.Iterator()
	value, valid:=it.Next()
	if valid {
		node:=new(element_DLList)
		node.value=value
		node.prev=nil
		top:=node
		bottom:=node
		self.size++
		value,valid=it.Next()
		for {
			node:=new(element_DLList)
			node.value=value
			node.prev=bottom
			bottom.next=node
			bottom=node
			self.size++
			value,valid=it.Next()
		}
		self.tail.prev.next=top
		top.prev=self.tail.prev
		bottom.next=self.tail
		self.tail.prev=bottom
	}
	return true
}

/* Returns true if the given object is contained in this list */
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

/* Returns true if all the elements of the collection are contained in this list */
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

/* Removes all the elements of this list */
func (self *DLList) Clear() {
	self.head.next=self.tail
	self.tail.prev=self.head
	self.size=0
}

/* Returns true if this list is empty */
func (self *DLList) IsEmpty() bool {
	return self.size==0
}

/* Removes the given object from this list */
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

/* Removes all the elements of the given collection from this list */
func (self *DLList) RemoveAll(c Collection, equal Equal) {
		if self==c {self.Clear()} else {
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
}

/* Removes all the elements from this list that are not in the given collection */
func (self *DLList) RetainAll(c Collection, equal Equal) {
	if self==c { return }
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

/* Returns the number of elements in this list */
func (self *DLList) Size() int {
	return self.size
}

/* The stack interface Peek function */
func (self *DLList) Peek() interface{} {
	return self.head.next.value
}

/* The stack interface Pop function */
func (self *DLList) Pop() interface{} {
	value:=self.head.next.value
	self.head.next.next.prev=self.head
	self.head.next=self.head.next.next
	self.size--
	return value
}

/* adds an element to the head of the list */
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

/* The stack interface Push function */
func (self *DLList) Push(element interface{}) bool {
	return self.add_first(element)
}

/* The stack interface Search function */
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

/* The queue interface AddTail function */
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

/* The queue interface PeekHead function */
func (self *DLList) PeekHead() interface{} {
	return self.Peek()
}

/* The queue interface PopHead function */
func (self *DLList) PopHead() interface{} {
	return self.Pop()
}

/* The deque interface AddHead function */
func (self *DLList) AddHead(element interface{}) bool {
	return self.add_first(element)
}

/* The deque interface PeekTail function */
func (self *DLList) PeekTail() interface{} {
	return self.tail.prev.value
}

/* The deque interface PopTail function */
func (self *DLList) PopTail() interface{} {
	value:=self.tail.prev.value
	self.tail.prev.prev.next=self.tail
	self.tail.prev=self.tail.prev.prev
	self.size--
	return value
}


