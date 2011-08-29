package gollection

import "testing"
import "gollection"

const (
	testElementCount=4321
)

func intEqual(o1 interface{}, o2 interface{}) bool {
	return o1==o2
}

func Test_empty_list_has_zero_size(t *testing.T) {
	list:=gollection.NewSLList()
	if list.Size()!=0 {
		t.Fatal("Expected a 0 size list");
	}
}

func Test_empty_list_iterator_starts_invalid(t *testing.T) {
	list:=gollection.NewSLList()
	it:=list.Iterator()
	_,valid:=it.Next()
	if valid {
		t.Fatal("The iterator should be invalid")
	}
}

func Test_prepend(t *testing.T) {
	list:=gollection.NewSLList()
	for i:=0;i<testElementCount;i++ {
		list.Prepend(i)
	}
	if list.Size()!=testElementCount {
		t.Fatal("The list size does not match the test size")
	}
	it:=list.Iterator()
	for i:=testElementCount-1;i>=0;i-- {
		value, valid:=it.Next()
		if !valid {
			t.Fatal("The iterator should be valid")
		}
		if !intEqual(value,i) {
			t.Fatal("The values should be equal")
		}
	}
	_, valid:=it.Next()
	if valid {
		t.Fatal("The iterator should be invalid")
	}
}

func Test_add(t *testing.T) {
	list:=gollection.NewSLList()
	for i:=0;i<testElementCount;i++ {
		list.Add(i)
	}
	if list.Size()!=testElementCount {
		t.Fatal("The size of the list does not match the test size")
	}
	it:=list.Iterator()
	for i:=0;i<testElementCount;i++ {
		value, valid:=it.Next()
		if !valid {
			t.Fatal("The iterator should be valid")
		}
		if !intEqual(value,i) {
			t.Fatal("The values should be equal")
		}
	}
	_,valid:=it.Next()
	if valid {
		t.Fatal("The iterator should be invalid")
	}
}

func Test_addAll(t *testing.T) {
	list1:=gollection.NewSLList()
	list2:=gollection.NewSLList()
	for i:=testElementCount-1;i>=0;i-- {
		list1.Prepend(i)
		list2.Prepend(i)
	}
	list1.AddAll(list2)
	it:=list1.Iterator()
	for j:=0;j<2;j++ {
		for i:=0;i<testElementCount;i++ {
			value,valid:=it.Next()
			if !valid {
				t.Fatal("The iterator should be valid")
			}
			if !intEqual(i,value) {
				t.Fatal("The values should be equal")
			}
		}
	}
	_,valid:=it.Next()
	if valid {
		t.Fatal("The iterator should be invalid")
	}
}

func Test_addAll_self(t *testing.T) {
	list1:=gollection.NewSLList()
	for i:=testElementCount-1;i>=0;i-- {
		list1.Prepend(i)
	}
	list1.AddAll(list1)
	it:=list1.Iterator()
	for j:=0;j<2;j++ {
		for i:=0;i<testElementCount;i++ {
			value,valid:=it.Next()
			if !valid {
				t.Fatal("The iterator should be valid")
			}
			if !intEqual(i,value) {
				t.Fatal("The values should be equal")
			}
		}
	}
	_,valid:=it.Next()
	if valid {
		t.Fatal("The iterator should be invalid")
	}
}

func Test_Clear(t *testing.T) {
	list:=gollection.NewSLList()
	/* clear empty list */
	list.Clear()
	if list.Size()!=0 {t.Fatal("The size of the list should be 0")}
	/* clear not empty list */
	for i:=0;i>testElementCount;i++ {
		list.Prepend(i)
	}
	list.Clear()
	if list.Size()!=0 {t.Fatal("The size of the list should be 0")}
}

func Test_Contains(t *testing.T) {
	list:=gollection.NewSLList()
	for i:=0;i<testElementCount;i++ {
		list.Prepend(i)
	}
	for i:=0;i<testElementCount;i++ {
		if !list.Contains(i,intEqual) {
			t.Fatal("Element %v should be contained in the list",i)
		}
	}
	if list.Contains(testElementCount,intEqual) {
		t.Fatal("Element %v should not be contained in the list",testElementCount)
	}
}

func Test_ContainsAll(t *testing.T) {
	// test list contains itself
	list:=gollection.NewSLList()
	for i:=0;i<testElementCount;i++ {
		list.Prepend(i)
	}
	if !list.ContainsAll(list, intEqual) {
		t.Fatal("The list must contain itself")
	}
	// test against different list
	list1:=gollection.NewSLList()
	list1.Add(testElementCount);
	list1.Add(1);
	if list.ContainsAll(list1,intEqual) {
		t.Fatal("The list should not be contained")
	}
	list1.Clear()
	for i:=0;i<testElementCount-1;i++ {
		list1.Prepend(i)
	}
	if !list.ContainsAll(list1,intEqual) {
		t.Fatal("The list must contain a sub-list")
	}
}

func Test_IsEmpty(t *testing.T) {
	list:=gollection.NewSLList()
	if !list.IsEmpty() {t.Fatal("The list must be empty")}
	for i:=0;i<testElementCount;i++ {
		list.Prepend(i)
	}
	if list.IsEmpty() {t.Fatal("The list should not be empty")}
	list1:=gollection.NewSLList()
	list1.Add(list)
	list1.Clear()
	if !list1.IsEmpty() {t.Fatal("The list1 must be empty")}
	it:=list.Iterator()
	for {
		_,valid:=it.Next()
		if !valid {break}
		it.Remove()
	}
	if !list.IsEmpty() {t.Fatal("The list must be empty")}
}





