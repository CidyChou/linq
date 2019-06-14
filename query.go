package linq

import "reflect"

// Iterator is 迭代器
type Iterator func() (item interface{}, ok bool)

// Query is 聚合类
type Query struct {
	Iterate func() Iterator
}

// Iterable is 接口
type Iterable interface {
	Iterate() Iterator
}

// KeyValue is Key-Value
type KeyValue struct {
	Key   interface{}
	Value interface{}
}

// From is 包装对象
func From(source interface{}) Query {
	src := reflect.ValueOf(source)

	switch src.Kind() {
	case reflect.Slice, reflect.Array:
		len := src.Len()

		return Query{
			Iterate: func() Iterator {
				index := 0

				return func() (item interface{}, ok bool) {
					ok = index < len
					if ok {
						item = src.Index(index).Interface()
						index++
					}

					return
				}
			},
		}
	case reflect.Map:
		len := src.Len()

		return Query{
			Iterate: func() Iterator {
				index := 0
				keys := src.MapKeys()

				return func() (item interface{}, ok bool) {
					ok = index < len
					if ok {
						key := keys[index]
						item = KeyValue{
							Key:   key.Interface(),
							Value: src.MapIndex(key).Interface(),
						}

						index++
					}

					return
				}
			},
		}
	default:
		return FromIterable(source.(Iterable))
	}
}

func FromIterable(source Iterable) Query {
	return Query{
		Iterate: source.Iterate,
	}
}
