package Bayes

import (
//	"fmt"
//	"io"
)

//interface for real element list interface
type Elelist interface {
	Delete(interface{}) bool
	Len() int
	Get(key interface{}) interface{}
	Set(key interface{}, value interface{})
	Updata(key interface{}, value interface{})
}

//cache data table interface
type Table interface {
	Read(key interface{})
	Write(key interface{}) bool
	Explode() bool
}
