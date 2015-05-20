package Bayes

import (
	"fmt"
	"io"
)

type KeyValueList map[interface{}]interface{}

func (l KeyValueList) Get(key interface{}) interface{} {
	return l[key]
}
func (l KeyValueList) Set(key interface{}, value interface{}) {
	l[key] = value
}

func (l KeyValueList) Delete(key interface{}) bool {
	delete(l[key])
	return true
}
func (l KeyValueList) Len() int {
	return len(l)
}
