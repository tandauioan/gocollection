
package gocollection

import "testing"

func sampleEqual (o1 interface{}, o2 interface{}) bool {
	return o1==o2
}

func useEqualAsGocollection(a int, b int, equal Equal) bool {
	return sampleEqual(a,b)
}

func TestSimple(t *testing.T) {
	if !useEqualAsGocollection(1,1,sampleEqual) {
		t.Fail()
	}
}


