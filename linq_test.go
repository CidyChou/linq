package linq

import (
	"testing"
)

func Test_Linq_Success(t *testing.T) {
	res := Range(0, 100, 1).Select(func(v interface{}) interface{} {
		return -v.(int)
	}).Where(func(v interface{}) bool {
		return v.(int) < -20
	}).First()

	if res != -21 {
		t.Error("结果有误")
	}
}

type student struct {
	name string
	age  int
}

func Test_Linq_Select(t *testing.T) {
	var students []student
	students = append(students, student{
		name: "George",
		age:  12,
	})

	students = append(students, student{
		name: "John",
		age:  15,
	})

	students = append(students, student{
		name: "cidy chou",
		age:  19,
	})

	res := From(students).Select(func(c interface{}) interface{} {
		return c.(student).name
	}).First()

	if res != "George" {
		t.Error("结果有误", res)
	}
}
