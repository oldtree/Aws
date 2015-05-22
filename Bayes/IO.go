package Bayes

import (
	"fmt"
	"io"
)

type IOAdaper struct {
	io.ReadWriter
	//DB address,if not db ,then other address
	Address string
}

func NewIOAdaper(Address string) *IOAdaper {
	rw := new(IOAdaper)
	rw.Address = Address
	return rw
}

/*
------key-------data
[0,1,2,3,4,5,6,7][8-*]
key byte maybe too long
*/

func (rw IOAdaper) Read(p []byte) (n int, err error) {
	fmt.Println("how stupid this function")
	return 0, nil
}
func (rw IOAdaper) Write(p []byte) (n int, err error) {
	fmt.Println("how stupid this function")
	return 0, nil
}
