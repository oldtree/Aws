package Bayes

import ()

type Hasher interface {
	Hash(doorkey interface{}) int
	SetServiceTaeget(ct *CacheTableList)
}

type DefaultHash struct {
	ct *CacheTableList
}

func (d *DefaultHash) Hash(doorkey interface{}) int {
	return doorkey.(int) % d.ct.Len()
}

func (d *DefaultHash) SetServiceTaeget(ct *CacheTableList) {
	d.ct = ct
}

type CacheTableList struct {
	Tables []Table
	Hash   Hasher
}

func NewCacheTableList() *CacheTableList {
	ct := new(CacheTableList)
	ct.Tables = make([]Table, 64)
	for index, _ := range ct.Tables {
		ct.Tables[index] = NewCacheTable()
	}
	ct.Hash = new(DefaultHash)
	ct.Hash.SetServiceTaeget(ct)
	return ct
}

func (ct *CacheTableList) SetHasher(h Hasher) {
	ct.Hash = h
}

func (ct *CacheTableList) Len() int {
	return len(ct.Tables)
}

func (ct *CacheTableList) GetTable(doorkey int64) Table {
	return ct.Tables[doorkey]
}
