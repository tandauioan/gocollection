
package gollection

/*
* A very basic singly linked list structure that generalizes
* the Collection and Stack interfaces. This is also a node in the list.
*/
type SLList struct { 
	Collection
	Stack

	/* the element in the current node */
	value interface{}
	/* the node following this node */
	next *SLList
}

/* 
* A forward iterator over SLList
*/
type sllist_forward_Iterator struct {Iterator;
	prev, node *SLList
}

/*
* Returns the next element of the list as long as valid is true
*/
func (self *sllist_forward_Iterator) Next() (value interface{}, valid bool) {
	if(self.node.next==nil) {
		value=nil
		valid=false
	} else {
		if self.prev!=self.node {
			self.prev=self.node
		}
		self.node=self.node.next
		value=self.node.value
		valid=true
	}
	return
}

/*
* Removes the node that was iterated over last 
*/
func (self *sllist_forward_Iterator) Remove() {
	if self.prev!=self.node {
		self.prev.next=self.node.next
		self.node=self.prev
	}
}

/*
* Creates and returns a new singly linked list
*/
func NewSLList() (list *SLList) {
	list=new(SLList)
	list.value=nil
	list.next=nil
	return
}

/*
* Returns the last node of the list by iterating over all of the nodes
*/
func (self *SLList) _get_last_node() (lastNode *SLList) {
	lastNode=self
	for lastNode.next!=nil {
		lastNode=lastNode.next
	}
	return
}

/*
* Returns a forward iterator for the list 
*/
func (self *SLList) Iterator() Iterator {
	iter:=new(sllist_forward_Iterator)
	iter.prev=self
	iter.node=self
	return iter
}

/*
* Adds a new element to the top of the list. The operation is O(1).
*/
func (self *SLList) Prepend(object interface{}) bool {
	node:=new(SLList)
	node.value=object
	node.next=self.next
	self.next=node
	return true;
}

/*
* Adds a new element to the tail of the list. This method iterates
* over all the elements of the list to find the last node and 
* adds the object and the end. The insert operation is O(n)
*/
func (self *SLList) Add(object interface{}) bool {
	lastNode:=self._get_last_node()
	node:=new(SLList)
	node.value=object
	node.next=nil
	lastNode.next=node
	return true
}

/*
* Adds all the elements in the collection to the end of this list. 
* The end of the list is searched for just once, and all the elements
* are added in one iteration so this is a lot more efficient than 
* calling Add(object) for each element in the collection.
*/
func (self *SLList) AddAll(c Collection) bool {
	lastNode:=self._get_last_node()
	it:=self.Iterator()
	value, valid:=it.Next()
	if valid {
		node:=new(SLList)
		node.value=value
		node.next=nil
		top:=node
		bottom:=node
		value,valid=it.Next()
		for valid {
			node:=new(SLList)
			node.value=value
			node.next=nil
			bottom.next=node
			bottom=node
			value,valid=it.Next()
		}
		lastNode.next=top
	}
	return true
}


/* Removes all the elements from this list */
func (self *SLList) Clear() {
	self.next=nil
}

/*
* Returns true if there is an element in this list that equals 
* the given object using the provided equal function
*/
func (self *SLList) Contains(object interface{}, equal Equal) bool {
	it:=self.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			if equal(object, value) {
				return true
			}
		}
	}
	return false
}

/*
* Returns true if all the elements of the collection are contained in this list
*/
func (self *SLList) ContainsAll(c Collection, equal Equal) bool {
	it:=c.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			if !self.Contains(value, equal) {
				return false
			}
		}
	}
	return true
}

/* Returns true if the list is empty */
func (self *SLList) IsEmpty() bool {
	return self.next==nil
}

/*
* Removes all the elements that are equal to the given object
* using the provided equal function.
*/
func (self *SLList) Remove(object interface{}, equal Equal) {
	node:=self
	for node.next!=nil {
		if equal(object, node.next.value) {
			node.next=node.next.next
		} else {
			node=node.next
		}
	}
}

/*
* Removes all the elements in the given collection from this list
*/
func (self *SLList) RemoveAll(c Collection, equal Equal) {
	if self==c {
		self.Clear()
	} else {
		it:=c.Iterator()
		for {
			value,valid:=it.Next()
			if !valid {
				break
			} else {
				self.Remove(value, equal)
			}
		}
	}
}

/*
* Removes all the elements from this list that are not contained
* in the given collection
*/
func (self *SLList) RetainAll(c Collection, equal Equal) {
	if self==c {
		return
	}
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

/*
* Returns the number of elements in this list. The operation is O(n) as 
* the function needs to be iterated over to find out its size.
*/
func (self *SLList) Size() int {
	it:=self.next
	count:=0
	for it!=nil {
		count++
		it=it.next
	}
	return count
}

/* Stack peek function */
func (self *SLList) Peek() interface{} {
	return self.next
}

/* Stack pop function */
func (self *SLList) Pop() interface{} {
	ret:=self.next
	if ret!=nil {
		self.next=self.next.next
	}
	return ret
}

/* Stack push function */
func (self *SLList) Push(element interface{}) {
	self.Prepend(element)
}

/* Stack search function */
func (self *SLList) Search(element interface{}, equal Equal) int {
	index:=0
	result:=-1
	it:=self.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			if equal(value,element) {
				result=index
				break
			} else {
				index++
			}
		}
	}
	return result
}

