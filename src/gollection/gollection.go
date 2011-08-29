/*
* 
* Copyright Ioan - Ciprian Tandau
* All rights reserved.
* This code is governed by GPL3 which can be found in the
* root of the project.
*
*/


package gollection

/*
* This is an equality function prototype; a predicate that
* returns true if its arguments are equal and false
* otherwise. The concrete function implementation
* defines what equal means.
*/
type Equal func(o1 interface{}, o2 interface{}) bool

/*
* This is a hash function prototype. It must consistently return
* the same value for the same object
*/
type Hash func(o1 interface{}) int


/*
* This is an iterator interface. Unless otherwise specified, 
* there is no defined behavior if the iteration goes beyond 
* the last element of the collection
*/
type Iterator interface {
	/*
	* Returns the next value of the iterated collection
	* and a validity flag. If the validity flag is false
	* then the iterator reached the end of the collection
	* and the value should not be read. The value should
	* only be read if the validity flag is true
	*/
	Next() (value interface{}, valid bool)
	/*
	* Removes the last element returned by the iterator
	* This method is allowed to do nothing but implementations
	* should document the behavior of the method
	*/
	Remove()
}

/*
* Collection interface. Declares operations that add  elements
* to the collection, remove elements from the collection and
* iterates over the elements of the collection.
*/
type Collection interface {
	/* Add an object to the collection */
	Add(object interface{}) bool
	/* 
	* Add all the objects from the given collection
	* to this collection
	*/
	AddAll(c Collection) bool
	/* Clear all elements from this collection */
	Clear()
	/*
	* true if the object equals any of the objects
	* in this collection and false otherwise. The 
	* object equality is tested using the equal 
	* function.
	*/
	Contains(object interface{}, equal Equal) bool
	/*
	* Like Contains but returns true only if all
	* the elements of the given collection are
	* contained in this collection
	*/
	ContainsAll(c Collection, equal Equal) bool
	/*
	* true if this collection has no elements and
	* false otherwise.
	*/
	IsEmpty() bool
	/*
	* removes all the elements that equal object
	* from this collection
	*/
	Remove(object interface{}, equal Equal)
	/*
	* Like Remove but for each elements in the 
	* given collection
	*/
	RemoveAll(c Collection, equal Equal)
	/*
	* Removes all the elements from this collection
	* that are not equal to at least one element
	* in the given collection
	*/
	RetainAll(c Collection, equal Equal)
	/*
	* Returns the number of elements in this collection
	*/
	Size() int
	/*
	* Returns an iterator over all the elements
	* in this collection.
	*/
	Iterator() Iterator
}

/*
* Declares the Stack interface. All the Collection interface
* operations inherited and, in addition, there are stack specific
* operations: peek, pop, push. 
*/
type Stack interface {
	Collection

	/*
	* Return the element at the top of the stack without
	* removing it.
	*/
	Peek() interface{}
	/*
	* Remove and return the element at the top of the stack 
	*/
	Pop() interface{}
	/* Push a new element to the top of the stack */
	Push(element interface{}) bool
	/*
	* Search for an element that equals the given element
	* and return it's position in the stack. Return -1 if 
	* the search finds no match.
	*/
	Search(element interface{}, equal Equal) int
}

/*
* Queue interface. The queue has a tail where elements are pushed
* and a head where elements are removed in a FIFO fashion.
*/
type Queue interface {
	Collection

	/* Adds a new element to the queue */
	AddTail(element interface{}) bool
	/* 
	* Returns the oldest element of the queue (the head)
	* without removing it.
	*/
	PeekHead() interface{}
	/* Removes the head of the queue and returns it */
	PopHead() interface{}
}

/* 
* A double-ended queue interface. It generalizes the queue
* interface and adds operations symmetrical operations to the 
* other ends of the queue. The Iterator() method returns 
* a forward iterator that iterates over the elements from 
* the head to the tail. A reverse iterator is also provided
* that iterates over the elements from the tail toward the head.
*/
type Deque interface {
	Queue

	/* Adds a new element to the head of the queue */
	AddHead(element interface{}) bool
	/* Returns the tail element without removing it */
	PeekTail() interface{}
	/* Removes and returns the tail of the queue */
	PopTail() interface{}
	/* 
	* Creates and returns a reverse iterator that will
	* iterate over the elements from the tail toward the head
	*/
	ReverseIterator() Iterator
}

