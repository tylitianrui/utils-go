package set

import "sync"

type Setter interface {
	Add(item interface{}) int
	Pop(item interface{}) int
	Len() int
	Intersection(set Setter) Setter
	Union(set Setter) Setter
	Difference(set Setter) Setter
	Subset(set Setter) bool
}
type HashSet struct {
	item map[interface{}]int8
	m    sync.Mutex
}

func (self *HashSet) Add(item interface{}) int {
	self.m.Lock()
	defer self.m.Unlock()
	self.item[item] = 1
	return len(self.item)

}

func (self *HashSet) Pop(item interface{}) int {
	self.m.Lock()
	defer self.m.Unlock()
	delete(self.item, item)
	return len(self.item)

}

func (self *HashSet) Len() int {
	self.m.Lock()
	defer self.m.Unlock()
	return len(self.item)

}

func (*HashSet) Intersection(set Setter) Setter {
	panic("implement me")
}

func (*HashSet) Union(set Setter) Setter {
	panic("implement me")
}

func (*HashSet) Difference(set Setter) Setter {
	panic("implement me")
}

func (*HashSet) Subset(set Setter) bool {
	panic("implement me")
}

func NewHashSet() *HashSet {
	return &HashSet{
		item: make(map[interface{}]int8),
	}
}
