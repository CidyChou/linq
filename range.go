package linq

// Range is 生成区间数据
func Range(start, stop, step int) Query {
	return Query{
		Iterate: func() Iterator {
			index := 0
			current := start

			return func() (item interface{}, ok bool) {
				if step == 0 || start == stop {
					return current, true
				}

				item, ok = current, true

				if start < stop {
					if step > 0 {
						index += step
						current += step
					}
				}

				if start > stop {
					if step < 0 {
						index += step
						current += step
					}
				}

				return
			}
		},
	}
}
