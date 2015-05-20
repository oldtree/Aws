package Bayes

import (
	"fmt"
	"io"
)

type Elelist interface {
	Delete(interface{}) bool
	Len() int
	Get(key interface{}) interface{}
	Set(value interface{})
}
type Table interface {
	Read()
	Write() bool
	Updata() bool
	Explode() bool
}
