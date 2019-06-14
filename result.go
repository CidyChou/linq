package linq

import "reflect"

// First is 返回第一个元素
func (q Query) First() interface{} {
	item, _ := q.Iterate()()
	return item
}

// Last is 返回最后一个元素
func (q Query) Last() interface{} {
	var r interface{}
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		r = item
	}
	return r
}

// FirstWith is 根据条件返回第一个元素
func (q Query) FirstWith(predicate func(interface{}) bool) interface{} {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			return item
		}
	}
	return nil
}

// All is 是否存在
func (q Query) All(predicate func(interface{}) bool) bool {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func (q Query) ToSlice(v interface{}) {
	res := reflect.ValueOf(v)
	slice := reflect.Indirect(res)

	cap := slice.Cap()
	res.Elem().Set(slice.Slice(0, cap))

	next := q.Iterate()
	index := 0

	for item, ok := next(); ok; item, ok = next() {
		if index >= cap {
			slice, cap = grow(slice)
		}
		slice.Index(index).Set(reflect.ValueOf(item))
		index++
	}
	res.Elem().Set(slice.Slice(0, index))
}

func grow(s reflect.Value) (v reflect.Value, newCap int) {
	cap := s.Cap()
	if cap == 0 {
		cap = 1
	} else {
		cap *= 2
	}
	newSlice := reflect.MakeSlice(s.Type(), cap, cap)
	reflect.Copy(newSlice, s)
	return newSlice, cap
}
