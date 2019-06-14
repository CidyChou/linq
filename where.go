package linq

// Where is 获取所有满足条件
func (q Query) Where(predicate func(interface{}) bool) Query {
	return Query{
		Iterate: func() Iterator {
			next := q.Iterate()
			return func() (item interface{}, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(item) {
						return
					}
				}
				return
			}
		},
	}
}
