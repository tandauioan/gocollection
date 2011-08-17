
package gocollection

type Equal func(o1 interface{}, o2 interface{}) bool

type Iterator interface {
	Next() (value interface{}, valid bool)
	Remove()
}

type Collection interface {
	Add(object interface{}) bool
	AddAll(c Collection) bool
	Clear()
	Contains(object interface{}, equal Equal) bool
	ContainsAll(c Collection, equal Equal) bool
	IsEmpty() bool
	Remove(object interface{}, equal Equal)
	RemoveAll(c Collection, equal Equal)
	RetainAll(c Collection, equal Equal)
	Size() int
	Iterator() Iterator
}

type Stack interface {
	Collection

	Peek() interface{}
	Pop() interface{}
	Push(element interface{}) bool
	Search(element interface{}, equal Equal) int
}

type Queue interface {
	Collection

	AddTail(element interface{}) bool
	PeekHead() interface{}
	PopHead() interface{}
}

type Deque interface {
	Queue

	AddHead(element interface{}) bool
	PeekTail() interface{}
	PopTail() interface{}
	ReverseIterator() Iterator
}

