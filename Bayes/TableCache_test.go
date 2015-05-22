package Bayes

import (
	"testing"
)

var ctl = NewCacheTableList()
var ta = ctl.GetTable(1)

func Test_TableGet(t *testing.T) {
	ta.(Elelist).Set("123", 123)
	result := ta.(Elelist).Get("123")
	if result != nil {
		if result.(int) != 123 {
			t.Error("get element error")
		}
	}
}
func Test_TableSet(t *testing.T) {

	ta.(Elelist).Set("456", 465)
	result := ta.(Elelist).Get("456")
	if result != nil {
		if result.(int) != 465 {
			t.Error("Set element error")
		}
	}

}
func Test_TableUpdata(t *testing.T) {

	ta.(Elelist).Updata("123", 456)
	ta.(Elelist).Updata("456", 123)
	result := ta.(Elelist).Get("123")
	if result != nil {
		if result.(int) != 456 {
			t.Error("Updata element error")
		}
	}
	result = ta.(Elelist).Get("456")
	if result != nil {
		if result.(int) != 123 {
			t.Error("Updata element error")
		}
	}
}
func Test_TableDelete(t *testing.T) {

	ta.(Elelist).Delete("123")
	ta.(Elelist).Delete("456")
	result := ta.(Elelist).Get("123")
	if result != nil {
		t.Error("Delete element error")
	}
	result = ta.(Elelist).Get("456")
	if result != nil {
		t.Error("Delete element error")
	}
}
