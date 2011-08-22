
package gocollection

type SLList struct { 
	Collection
	Stack


	value interface{}
	next *SLList
}

type sllist_forward_Iterator struct {Iterator;
	prev, node *SLList
}

func (self *sllist_forward_Iterator) HasNext() bool {
	return self.node.next!=nil
}

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

func (self *sllist_forward_Iterator) Remove() {
	if self.prev!=self.node {
		self.prev.next=self.node.next
		self.node=self.prev
	}
}

func NewSLList() (list *SLList) {
	list=new(SLList)
	list.value=nil
	list.next=nil
	return
}


func (self *SLList) _get_last_node() (lastNode *SLList) {
	lastNode=self
	for lastNode.next!=nil {
		lastNode=lastNode.next
	}
	return
}

func (self *SLList) Iterator() Iterator {
	iter:=new(sllist_forward_Iterator)
	iter.prev=self
	iter.node=self
	return iter
}

func (self *SLList) Prepend(object interface{}) bool {
	node:=new(SLList)
	node.value=object
	node.next=self.next
	self.next=node
	return true;
}

func (self *SLList) Add(object interface{}) bool {
	lastNode:=self._get_last_node()
	node:=new(SLList)
	node.value=object
	node.next=nil
	lastNode.next=node
	return true
}

func (self *SLList) AddAll(c Collection) bool {
	lastNode:=self._get_last_node()
	it:=c.Iterator()
	for {
		value,valid:=it.Next()
		if !valid {
			break
		} else {
			node:=new(SLList)
			node.value=value
			lastNode.next=node
			lastNode=node
		}
	}
	lastNode.next=nil
	return true
}

func (self *SLList) Clear() {
	self.next=nil
}

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

func (self *SLList) IsEmpty() bool {
	return self.next==nil
}

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

func (self *SLList) RemoveAll(c Collection, equal Equal) {
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

func (self *SLList) RetainAll(c Collection, equal Equal) {
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

func (self *SLList) Size() int {
	it:=self.next
	count:=0
	for it!=nil {
		count++
		it=it.next
	}
	return count
}

func (self *SLList) Peek() interface{} {
	return self.next
}

func (self *SLList) Pop() interface{} {
	ret:=self.next
	if ret!=nil {
		self.next=self.next.next
	}
	return ret
}

func (self *SLList) Push(element interface{}) {
	self.Prepend(element)
}

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

