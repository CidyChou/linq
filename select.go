package linq

// Select is 查询
func (q Query) Select(selector func(interface{}) interface{}) Query {
	return Query{
		Iterate: func() Iterator {
			next := q.Iterate()
			return func() (item interface{}, ok bool) {
				var it interface{}
				it, ok = next()
				if ok {
					item = selector(it)
				}
				return
			}
		},
	}
}
