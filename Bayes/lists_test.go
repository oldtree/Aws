package Bayes

import (
	"testing"
)

//an element list  inplements base on map[interface{}]interface{}
var l = make(KeyValueList)

func Test_Get(T *testing.T) {
	l.Set(1, 1)
	l.Set(2, 2)
	l.Set(3, 3)
	if l.Get(2).(int) != 2 {
		T.Fatal()
	}
}
func Test_Delete(T *testing.T) {
	l.Delete(2)
	if l.Get(2) != nil {
		T.Fatal()
	}
}
func Test_Set(T *testing.T) {
	l.Set(4, 4)
	if l.Get(4).(int) != 4 {
		T.Fatal()
	}
}
func Test_Len(T *testing.T) {
	if l.Len() < 2 {
		T.Error()
	}
}
