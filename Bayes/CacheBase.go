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
	Set(value interface{})
}

//cache data table interface
type Table interface {
	Read()
	Write() bool
	Updata() bool
	Explode() bool
}
