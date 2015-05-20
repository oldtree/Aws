package Bayes

import (
	"fmt"
	//	"io"
)

type CacheTable struct {
	//key range
	StartKey interface{}
	EndKey   interface{}
	Hit      int
	Lose     int
	Remote   bool
	Address  string
	Elements Elelist
}

// for read /write /updata operation with DB
func (c *CacheTable) Read() {

}

// for read /write /updata operation with DB
func (c *CacheTable) Write() bool {
	return false

}

// for read /write /updata operation with DB
func (c *CacheTable) Updata() bool {

	return false

}

//for product some new Virtual CachePoint
func (c *CacheTable) Explode() bool {
	fmt.Println("not implement yet")
	return false
}
