package Bayes

import (
	"fmt"
	"io"
)

type CacheTable struct {
	StartKey interface{}
	EndKey   interface{}
	Hit      int
	Lose     int
	Remote   bool
	Address  string
	Elements Elelist
}

func (c *CacheTable) Read() {

}
func (c *CacheTable) Write() bool  {}
func (c *CacheTable) Updata() bool {}
func (c *CacheTable) Explode() bool {
	fmt.Println("not implement yet")
}
