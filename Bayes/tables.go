package Bayes

import (
	"fmt"
	"io"
)

const (
	READ_OPCODE = iota
	WRITE_OPCODE
	UPDATA_OPCODE
)

type CacheTable struct {
	//key range
	StartKey interface{}
	EndKey   interface{}
	//for location hit rate
	Hit  int
	Lose int
	//for if the cache machine is not local if this machine is main machine
	Remote bool

	//a writer reader implements io.ReadWriter interface
	ReadWrite io.ReadWriter
	//real list for store k-v,is a interface
	Elements Elelist
	//task queue ,not  a best deal
	TaskQueue chan *TableOperationCell
}

func NewCacheTable() *CacheTable {
	c := new(CacheTable)
	c.Hit = 0
	c.Lose = 0
	c.TaskQueue = make(chan *TableOperationCell, 512)
	c.ReadWrite = *NewIOAdaper()
	var kv KeyValueList
	kv = make(map[interface{}]interface{}, 1024)
	c.Elements = kv
	return c
}

func (c *CacheTable) Get(key interface{}) interface{} {
	return c.Elements.Get(key)
}
func (c *CacheTable) Set(key interface{}, value interface{}) {
	c.Elements.Set(key, value)
}

func (c *CacheTable) Updata(key interface{}, value interface{}) {
	c.Elements.Updata(key, value)
}
func (c *CacheTable) Delete(key interface{}) bool {
	return c.Elements.Delete(key)
}

func (c *CacheTable) Len() int {
	return c.Elements.Len()
}

type TableOperationCell struct {
	Key  interface{}
	Code int
}

func (c *CacheTable) GenerateCell(key interface{}, code int) {
	cell := &TableOperationCell{
		Key:  key,
		Code: code,
	}
	c.TaskQueue <- cell
}

func (c *CacheTable) TaskLoop() {
	for {
		select {
		case cell := <-c.TaskQueue:
			switch cell.Code {
			case READ_OPCODE:
				go c.Read(cell.Key)
			case WRITE_OPCODE:
				c.Write(cell.Key)
			case UPDATA_OPCODE:
				//				go c.Updata(cell.Key)
			}
		}
	}
}

// for read /write /updata operation with DB
func (c *CacheTable) Read(key interface{}) {

}

// for read /write /updata operation with DB
func (c *CacheTable) Write(key interface{}) bool {
	return false
}

//for product some new Virtual CachePoint
func (c *CacheTable) Explode() bool {
	fmt.Println("not implement yet")
	return false
}
